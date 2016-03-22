// Code generated by protoc-gen-gogo.
// source: envelope.proto
// DO NOT EDIT!

/*
	Package events is a generated protocol buffer package.

	It is generated from these files:
		envelope.proto
		error.proto
		http.proto
		log.proto
		metric.proto
		uuid.proto

	It has these top-level messages:
		Envelope
		Error
		HttpStart
		HttpStop
		HttpStartStop
		LogMessage
		ValueMetric
		CounterEvent
		ContainerMetric
		UUID
*/
package events

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

// / Type of the wrapped event.
type Envelope_EventType int32

const (
	// Removed Heartbeat at position 1
	Envelope_HttpStart       Envelope_EventType = 2
	Envelope_HttpStop        Envelope_EventType = 3
	Envelope_HttpStartStop   Envelope_EventType = 4
	Envelope_LogMessage      Envelope_EventType = 5
	Envelope_ValueMetric     Envelope_EventType = 6
	Envelope_CounterEvent    Envelope_EventType = 7
	Envelope_Error           Envelope_EventType = 8
	Envelope_ContainerMetric Envelope_EventType = 9
)

var Envelope_EventType_name = map[int32]string{
	2: "HttpStart",
	3: "HttpStop",
	4: "HttpStartStop",
	5: "LogMessage",
	6: "ValueMetric",
	7: "CounterEvent",
	8: "Error",
	9: "ContainerMetric",
}
var Envelope_EventType_value = map[string]int32{
	"HttpStart":       2,
	"HttpStop":        3,
	"HttpStartStop":   4,
	"LogMessage":      5,
	"ValueMetric":     6,
	"CounterEvent":    7,
	"Error":           8,
	"ContainerMetric": 9,
}

func (x Envelope_EventType) Enum() *Envelope_EventType {
	p := new(Envelope_EventType)
	*p = x
	return p
}
func (x Envelope_EventType) String() string {
	return proto.EnumName(Envelope_EventType_name, int32(x))
}
func (x *Envelope_EventType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Envelope_EventType_value, data, "Envelope_EventType")
	if err != nil {
		return err
	}
	*x = Envelope_EventType(value)
	return nil
}
func (Envelope_EventType) EnumDescriptor() ([]byte, []int) { return fileDescriptorEnvelope, []int{0, 0} }

// / Envelope wraps an Event and adds metadata.
type Envelope struct {
	Origin     *string             `protobuf:"bytes,1,req,name=origin" json:"origin,omitempty"`
	EventType  *Envelope_EventType `protobuf:"varint,2,req,name=eventType,enum=events.Envelope_EventType" json:"eventType,omitempty"`
	Timestamp  *int64              `protobuf:"varint,6,opt,name=timestamp" json:"timestamp,omitempty"`
	Deployment *string             `protobuf:"bytes,13,opt,name=deployment" json:"deployment,omitempty"`
	Job        *string             `protobuf:"bytes,14,opt,name=job" json:"job,omitempty"`
	Index      *string             `protobuf:"bytes,15,opt,name=index" json:"index,omitempty"`
	Ip         *string             `protobuf:"bytes,16,opt,name=ip" json:"ip,omitempty"`
	// Removed Heartbeat at position 3
	HttpStart        *HttpStart       `protobuf:"bytes,4,opt,name=httpStart" json:"httpStart,omitempty"`
	HttpStop         *HttpStop        `protobuf:"bytes,5,opt,name=httpStop" json:"httpStop,omitempty"`
	HttpStartStop    *HttpStartStop   `protobuf:"bytes,7,opt,name=httpStartStop" json:"httpStartStop,omitempty"`
	LogMessage       *LogMessage      `protobuf:"bytes,8,opt,name=logMessage" json:"logMessage,omitempty"`
	ValueMetric      *ValueMetric     `protobuf:"bytes,9,opt,name=valueMetric" json:"valueMetric,omitempty"`
	CounterEvent     *CounterEvent    `protobuf:"bytes,10,opt,name=counterEvent" json:"counterEvent,omitempty"`
	Error            *Error           `protobuf:"bytes,11,opt,name=error" json:"error,omitempty"`
	ContainerMetric  *ContainerMetric `protobuf:"bytes,12,opt,name=containerMetric" json:"containerMetric,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptorEnvelope, []int{0} }

func (m *Envelope) GetOrigin() string {
	if m != nil && m.Origin != nil {
		return *m.Origin
	}
	return ""
}

func (m *Envelope) GetEventType() Envelope_EventType {
	if m != nil && m.EventType != nil {
		return *m.EventType
	}
	return Envelope_HttpStart
}

func (m *Envelope) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *Envelope) GetDeployment() string {
	if m != nil && m.Deployment != nil {
		return *m.Deployment
	}
	return ""
}

func (m *Envelope) GetJob() string {
	if m != nil && m.Job != nil {
		return *m.Job
	}
	return ""
}

func (m *Envelope) GetIndex() string {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return ""
}

func (m *Envelope) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *Envelope) GetHttpStart() *HttpStart {
	if m != nil {
		return m.HttpStart
	}
	return nil
}

func (m *Envelope) GetHttpStop() *HttpStop {
	if m != nil {
		return m.HttpStop
	}
	return nil
}

func (m *Envelope) GetHttpStartStop() *HttpStartStop {
	if m != nil {
		return m.HttpStartStop
	}
	return nil
}

func (m *Envelope) GetLogMessage() *LogMessage {
	if m != nil {
		return m.LogMessage
	}
	return nil
}

func (m *Envelope) GetValueMetric() *ValueMetric {
	if m != nil {
		return m.ValueMetric
	}
	return nil
}

func (m *Envelope) GetCounterEvent() *CounterEvent {
	if m != nil {
		return m.CounterEvent
	}
	return nil
}

func (m *Envelope) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Envelope) GetContainerMetric() *ContainerMetric {
	if m != nil {
		return m.ContainerMetric
	}
	return nil
}

func init() {
	proto.RegisterType((*Envelope)(nil), "events.Envelope")
	proto.RegisterEnum("events.Envelope_EventType", Envelope_EventType_name, Envelope_EventType_value)
}
func (m *Envelope) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Envelope) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Origin == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("origin")
	} else {
		data[i] = 0xa
		i++
		i = encodeVarintEnvelope(data, i, uint64(len(*m.Origin)))
		i += copy(data[i:], *m.Origin)
	}
	if m.EventType == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("eventType")
	} else {
		data[i] = 0x10
		i++
		i = encodeVarintEnvelope(data, i, uint64(*m.EventType))
	}
	if m.HttpStart != nil {
		data[i] = 0x22
		i++
		i = encodeVarintEnvelope(data, i, uint64(m.HttpStart.Size()))
		n1, err := m.HttpStart.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.HttpStop != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintEnvelope(data, i, uint64(m.HttpStop.Size()))
		n2, err := m.HttpStop.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Timestamp != nil {
		data[i] = 0x30
		i++
		i = encodeVarintEnvelope(data, i, uint64(*m.Timestamp))
	}
	if m.HttpStartStop != nil {
		data[i] = 0x3a
		i++
		i = encodeVarintEnvelope(data, i, uint64(m.HttpStartStop.Size()))
		n3, err := m.HttpStartStop.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.LogMessage != nil {
		data[i] = 0x42
		i++
		i = encodeVarintEnvelope(data, i, uint64(m.LogMessage.Size()))
		n4, err := m.LogMessage.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.ValueMetric != nil {
		data[i] = 0x4a
		i++
		i = encodeVarintEnvelope(data, i, uint64(m.ValueMetric.Size()))
		n5, err := m.ValueMetric.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.CounterEvent != nil {
		data[i] = 0x52
		i++
		i = encodeVarintEnvelope(data, i, uint64(m.CounterEvent.Size()))
		n6, err := m.CounterEvent.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.Error != nil {
		data[i] = 0x5a
		i++
		i = encodeVarintEnvelope(data, i, uint64(m.Error.Size()))
		n7, err := m.Error.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.ContainerMetric != nil {
		data[i] = 0x62
		i++
		i = encodeVarintEnvelope(data, i, uint64(m.ContainerMetric.Size()))
		n8, err := m.ContainerMetric.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	if m.Deployment != nil {
		data[i] = 0x6a
		i++
		i = encodeVarintEnvelope(data, i, uint64(len(*m.Deployment)))
		i += copy(data[i:], *m.Deployment)
	}
	if m.Job != nil {
		data[i] = 0x72
		i++
		i = encodeVarintEnvelope(data, i, uint64(len(*m.Job)))
		i += copy(data[i:], *m.Job)
	}
	if m.Index != nil {
		data[i] = 0x7a
		i++
		i = encodeVarintEnvelope(data, i, uint64(len(*m.Index)))
		i += copy(data[i:], *m.Index)
	}
	if m.Ip != nil {
		data[i] = 0x82
		i++
		data[i] = 0x1
		i++
		i = encodeVarintEnvelope(data, i, uint64(len(*m.Ip)))
		i += copy(data[i:], *m.Ip)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Envelope(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Envelope(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintEnvelope(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Envelope) Size() (n int) {
	var l int
	_ = l
	if m.Origin != nil {
		l = len(*m.Origin)
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.EventType != nil {
		n += 1 + sovEnvelope(uint64(*m.EventType))
	}
	if m.HttpStart != nil {
		l = m.HttpStart.Size()
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.HttpStop != nil {
		l = m.HttpStop.Size()
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.Timestamp != nil {
		n += 1 + sovEnvelope(uint64(*m.Timestamp))
	}
	if m.HttpStartStop != nil {
		l = m.HttpStartStop.Size()
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.LogMessage != nil {
		l = m.LogMessage.Size()
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.ValueMetric != nil {
		l = m.ValueMetric.Size()
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.CounterEvent != nil {
		l = m.CounterEvent.Size()
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.Error != nil {
		l = m.Error.Size()
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.ContainerMetric != nil {
		l = m.ContainerMetric.Size()
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.Deployment != nil {
		l = len(*m.Deployment)
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.Job != nil {
		l = len(*m.Job)
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.Index != nil {
		l = len(*m.Index)
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.Ip != nil {
		l = len(*m.Ip)
		n += 2 + l + sovEnvelope(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovEnvelope(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEnvelope(x uint64) (n int) {
	return sovEnvelope(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Envelope) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnvelope
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Envelope: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Envelope: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Origin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Origin = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventType", wireType)
			}
			var v Envelope_EventType
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (Envelope_EventType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EventType = &v
			hasFields[0] |= uint64(0x00000002)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpStart", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HttpStart == nil {
				m.HttpStart = &HttpStart{}
			}
			if err := m.HttpStart.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpStop", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HttpStop == nil {
				m.HttpStop = &HttpStop{}
			}
			if err := m.HttpStop.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Timestamp = &v
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpStartStop", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HttpStartStop == nil {
				m.HttpStartStop = &HttpStartStop{}
			}
			if err := m.HttpStartStop.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogMessage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LogMessage == nil {
				m.LogMessage = &LogMessage{}
			}
			if err := m.LogMessage.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValueMetric", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ValueMetric == nil {
				m.ValueMetric = &ValueMetric{}
			}
			if err := m.ValueMetric.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CounterEvent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CounterEvent == nil {
				m.CounterEvent = &CounterEvent{}
			}
			if err := m.CounterEvent.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Error == nil {
				m.Error = &Error{}
			}
			if err := m.Error.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContainerMetric", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ContainerMetric == nil {
				m.ContainerMetric = &ContainerMetric{}
			}
			if err := m.ContainerMetric.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deployment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Deployment = &s
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Job", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Job = &s
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Index = &s
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ip", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Ip = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEnvelope(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEnvelope
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("origin")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("eventType")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEnvelope(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEnvelope
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthEnvelope
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEnvelope
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipEnvelope(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthEnvelope = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEnvelope   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorEnvelope = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x95, 0x34, 0x49, 0xb3, 0x63, 0x3b, 0xd9, 0x4e, 0x41, 0xac, 0x22, 0x84, 0x4a, 0x84,
	0x50, 0x84, 0xa8, 0x8b, 0x7a, 0xe7, 0x52, 0x54, 0xc4, 0x81, 0x9e, 0x40, 0xdc, 0x1d, 0x67, 0xeb,
	0x1a, 0xc5, 0x9e, 0xd5, 0x7a, 0x1d, 0x91, 0x97, 0xe0, 0xb9, 0x38, 0xf2, 0x08, 0x88, 0x17, 0xe0,
	0x15, 0xd8, 0x4c, 0xec, 0xfc, 0x81, 0xc3, 0x4a, 0x33, 0x3b, 0xbf, 0x99, 0xd9, 0xef, 0xb3, 0x61,
	0xa4, 0xcb, 0x95, 0x5e, 0x92, 0xd1, 0xb1, 0xb1, 0xe4, 0x08, 0x07, 0x7a, 0xa5, 0x4b, 0x57, 0x4d,
	0x2e, 0xb3, 0xdc, 0x3d, 0xd4, 0xf3, 0x38, 0xa5, 0xe2, 0x2a, 0xa3, 0x8c, 0xae, 0xb8, 0x3c, 0xaf,
	0xef, 0x39, 0xe3, 0x84, 0xa3, 0x6d, 0xdb, 0x04, 0x1e, 0x9c, 0x33, 0x4d, 0x2c, 0x96, 0x94, 0x35,
	0x61, 0x58, 0x68, 0x67, 0xf3, 0xb4, 0xc9, 0x02, 0x6d, 0x2d, 0xd9, 0x6d, 0x32, 0xfd, 0xd3, 0x83,
	0xe1, 0x6d, 0xb3, 0x1b, 0x47, 0x30, 0x20, 0x9b, 0x67, 0x79, 0xa9, 0x3a, 0x17, 0xdd, 0x99, 0xc0,
	0x4b, 0x10, 0xfc, 0x8e, 0xcf, 0x6b, 0xa3, 0x55, 0xd7, 0x5f, 0x8d, 0xae, 0x27, 0xf1, 0xf6, 0x65,
	0x71, 0xdb, 0x14, 0xdf, 0xb6, 0x04, 0xbe, 0x00, 0xb1, 0xd9, 0xff, 0xc9, 0x25, 0xd6, 0xa9, 0xde,
	0x45, 0x67, 0x16, 0x5c, 0x9f, 0xb5, 0xf8, 0x87, 0xb6, 0x80, 0x53, 0x18, 0x6e, 0x29, 0x32, 0xaa,
	0xcf, 0x90, 0x3c, 0x86, 0xc8, 0xe0, 0x19, 0x08, 0x97, 0x17, 0xba, 0x72, 0x49, 0x61, 0xd4, 0xc0,
	0x43, 0x27, 0xf8, 0x1a, 0xa2, 0xdd, 0x70, 0xee, 0x3d, 0xe5, 0xde, 0xc7, 0xff, 0x2d, 0xe0, 0x01,
	0x2f, 0x01, 0xbc, 0xfc, 0x3b, 0x5d, 0x55, 0x49, 0xa6, 0xd5, 0x90, 0x51, 0x6c, 0xd1, 0x8f, 0xbb,
	0x0a, 0xce, 0x20, 0x58, 0x25, 0xcb, 0x5a, 0xdf, 0xb1, 0x41, 0x4a, 0x30, 0x78, 0xde, 0x82, 0x5f,
	0xf6, 0x25, 0x7c, 0x05, 0x61, 0x4a, 0x75, 0xe9, 0xb4, 0x65, 0xc1, 0x0a, 0x18, 0x7d, 0xd4, 0xa2,
	0xef, 0x0e, 0x6a, 0xf8, 0x14, 0xfa, 0xec, 0xb1, 0x0a, 0x18, 0x8a, 0x76, 0x9e, 0x6d, 0x2e, 0xf1,
	0x0d, 0x8c, 0x53, 0x2a, 0x5d, 0x92, 0x97, 0xda, 0x36, 0x7b, 0x43, 0xe6, 0x9e, 0xec, 0x87, 0x1d,
	0x95, 0x11, 0x01, 0x16, 0xda, 0x2c, 0x69, 0x5d, 0x6c, 0x36, 0x47, 0x1e, 0x16, 0x18, 0xc0, 0xc9,
	0x57, 0x9a, 0xab, 0x11, 0x27, 0x11, 0xf4, 0xf3, 0x72, 0xa1, 0xbf, 0xa9, 0x31, 0xa7, 0x00, 0xdd,
	0xdc, 0x28, 0xb9, 0x89, 0xa7, 0xdf, 0x3b, 0x20, 0xf6, 0x9f, 0x28, 0x02, 0xb1, 0x33, 0x4a, 0x76,
	0x31, 0x84, 0x61, 0xeb, 0xb9, 0x3c, 0xf1, 0xae, 0x47, 0x47, 0x2e, 0xca, 0x9e, 0xff, 0x23, 0x60,
	0xef, 0x96, 0xec, 0xe3, 0x18, 0x82, 0x03, 0x53, 0xe4, 0x00, 0x25, 0x84, 0x87, 0xd2, 0xe5, 0x29,
	0x0a, 0xe8, 0xb3, 0x4e, 0x39, 0xc4, 0x73, 0x18, 0xff, 0x23, 0x45, 0x8a, 0x9b, 0xb7, 0x3f, 0x7e,
	0x3f, 0xeb, 0xfc, 0xf4, 0xe7, 0x97, 0x3f, 0xf0, 0x9c, 0x6c, 0x16, 0xa7, 0x4b, 0xaa, 0x17, 0xf7,
	0x7e, 0xcc, 0xc2, 0xae, 0xe3, 0x85, 0x25, 0x53, 0x91, 0x57, 0xd3, 0xb8, 0x71, 0x13, 0xf2, 0xe4,
	0xf7, 0x49, 0xea, 0xc8, 0xae, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x5a, 0x34, 0x2a, 0x2a,
	0x03, 0x00, 0x00,
}