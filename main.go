package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/CrowdSurge/banner"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/boltdb/bolt"
	"github.com/cloudfoundry-community/firehose-to-syslog/caching"
	"github.com/cloudfoundry-community/firehose-to-syslog/firehose"
	"github.com/cloudfoundry-community/firehose-to-syslog/logging"
	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/pivotalservices/app-usage-nozzle/service"
	"github.com/pivotalservices/app-usage-nozzle/usageevents"
)

var (
	debug              = kingpin.Flag("debug", "Enable debug mode. This disables forwarding to syslog").Default("false").OverrideDefaultFromEnvar("DEBUG").Bool()
	apiEndpoint        = kingpin.Flag("api-endpoint", "Api endpoint address. For bosh-lite installation of CF: https://api.10.244.0.34.xip.io").OverrideDefaultFromEnvar("API_ENDPOINT").Required().String()
	dopplerEndpoint    = kingpin.Flag("doppler-endpoint", "Overwrite default doppler endpoint return by /v2/info").OverrideDefaultFromEnvar("DOPPLER_ENDPOINT").String()
	syslogServer       = kingpin.Flag("syslog-server", "Syslog server.").OverrideDefaultFromEnvar("SYSLOG_ENDPOINT").String()
	subscriptionID     = kingpin.Flag("subscription-id", "Id for the subscription.").Default("firehose").OverrideDefaultFromEnvar("FIREHOSE_SUBSCRIPTION_ID").String()
	user               = kingpin.Flag("user", "Admin user.").Default("admin").OverrideDefaultFromEnvar("FIREHOSE_USER").String()
	password           = kingpin.Flag("password", "Admin password.").Default("admin").OverrideDefaultFromEnvar("FIREHOSE_PASSWORD").String()
	skipSSLValidation  = kingpin.Flag("skip-ssl-validation", "Please don't").Default("false").OverrideDefaultFromEnvar("SKIP_SSL_VALIDATION").Bool()
	logEventTotals     = kingpin.Flag("log-event-totals", "Logs the counters for all selected events since nozzle was last started.").Default("false").OverrideDefaultFromEnvar("LOG_EVENT_TOTALS").Bool()
	logEventTotalsTime = kingpin.Flag("log-event-totals-time", "How frequently the event totals are calculated (in sec).").Default("30s").OverrideDefaultFromEnvar("LOG_EVENT_TOTALS_TIME").Duration()
	boltDatabasePath   = kingpin.Flag("boltdb-path", "Bolt Database path ").Default("my.db").OverrideDefaultFromEnvar("BOLTDB_PATH").String()
	tickerTime         = kingpin.Flag("cc-pull-time", "CloudController Polling time in sec").Default("60s").OverrideDefaultFromEnvar("CF_PULL_TIME").Duration()
	extraFields        = kingpin.Flag("extra-fields", "Extra fields you want to annotate your events with, example: '--extra-fields=env:dev,something:other ").Default("").OverrideDefaultFromEnvar("EXTRA_FIELDS").String()
)

const (
	version = "0.0.1"
)

func main() {

	banner.Print("usage nozzle")
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	// Start web server
	go func() {
		server := service.NewServer()
		server.Run(":" + port)
	}()

	kingpin.Version(version)
	kingpin.Parse()
	logging.LogStd(fmt.Sprintf("Starting app-usage-nozzle %s ", version), true)

	logging.SetupLogging(*syslogServer, *debug)

	c := cfclient.Config{
		ApiAddress:        *apiEndpoint,
		Username:          *user,
		Password:          *password,
		SkipSslValidation: *skipSSLValidation,
	}
	cfClient := cfclient.NewClient(&c)

	if len(*dopplerEndpoint) > 0 {
		cfClient.Endpoint.DopplerEndpoint = *dopplerEndpoint
	}
	logging.LogStd(fmt.Sprintf("Using %s as doppler endpoint", cfClient.Endpoint.DopplerEndpoint), true)

	//Use bolt for in-memory  - file caching
	db, err := bolt.Open(*boltDatabasePath, 0600, &bolt.Options{Timeout: 3 * time.Second})
	if err != nil {
		log.Fatal("Error opening bolt db: ", err)
		os.Exit(1)

	}
	defer db.Close()

	caching.SetCfClient(cfClient)
	caching.SetAppDb(db)
	caching.CreateBucket()

	//Let's Update the database the first time
	logging.LogStd("Start filling app/space/org cache.", true)
	apps := caching.GetAllApp()
	for idx := range apps {
		org := apps[idx].OrgName
		space := apps[idx].SpaceName
		app := apps[idx].Name
		key := usageevents.GetMapKeyFromAppData(org, space, app)
		usageevents.AppStats[key] = usageevents.ApplicationStat{AppName: app, SpaceName: space, OrgName: org}
	}

	logging.LogStd(fmt.Sprintf("Done filling cache! Found [%d] Apps", len(apps)), true)

	// Ticker Pooling the CC every X sec
	ccPolling := time.NewTicker(*tickerTime)

	go func() {
		for range ccPolling.C {
			logging.LogStd("Re-loading application cache.", true)
			apps = caching.GetAllApp()
		}
	}()

	firehose := firehose.CreateFirehoseChan(cfClient.Endpoint.DopplerEndpoint, cfClient.GetToken(), *subscriptionID, *skipSSLValidation)
	if firehose != nil {
		logging.LogStd("Firehose Subscription Succesfull! Routing events...", true)
		usageevents.ProcessEvents(firehose)
	} else {
		logging.LogError("Failed connecting to Firehose...Please check settings and try again!", "")
	}
}