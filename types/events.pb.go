// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// An event emitted from a smart contract
type Event struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	SequenceNumber       uint64   `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	EventData            []byte   `protobuf:"bytes,3,opt,name=event_data,json=eventData,proto3" json:"event_data,omitempty"`
	TypeTag              []byte   `protobuf:"bytes,4,opt,name=type_tag,json=typeTag,proto3" json:"type_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Event) GetSequenceNumber() uint64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *Event) GetEventData() []byte {
	if m != nil {
		return m.EventData
	}
	return nil
}

func (m *Event) GetTypeTag() []byte {
	if m != nil {
		return m.TypeTag
	}
	return nil
}

// An event along with the proof for the event
type EventWithProof struct {
	TransactionVersion   uint64      `protobuf:"varint,1,opt,name=transaction_version,json=transactionVersion,proto3" json:"transaction_version,omitempty"`
	EventIndex           uint64      `protobuf:"varint,2,opt,name=event_index,json=eventIndex,proto3" json:"event_index,omitempty"`
	Event                *Event      `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
	Proof                *EventProof `protobuf:"bytes,4,opt,name=proof,proto3" json:"proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *EventWithProof) Reset()         { *m = EventWithProof{} }
func (m *EventWithProof) String() string { return proto.CompactTextString(m) }
func (*EventWithProof) ProtoMessage()    {}
func (*EventWithProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{1}
}

func (m *EventWithProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventWithProof.Unmarshal(m, b)
}
func (m *EventWithProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventWithProof.Marshal(b, m, deterministic)
}
func (m *EventWithProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventWithProof.Merge(m, src)
}
func (m *EventWithProof) XXX_Size() int {
	return xxx_messageInfo_EventWithProof.Size(m)
}
func (m *EventWithProof) XXX_DiscardUnknown() {
	xxx_messageInfo_EventWithProof.DiscardUnknown(m)
}

var xxx_messageInfo_EventWithProof proto.InternalMessageInfo

func (m *EventWithProof) GetTransactionVersion() uint64 {
	if m != nil {
		return m.TransactionVersion
	}
	return 0
}

func (m *EventWithProof) GetEventIndex() uint64 {
	if m != nil {
		return m.EventIndex
	}
	return 0
}

func (m *EventWithProof) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *EventWithProof) GetProof() *EventProof {
	if m != nil {
		return m.Proof
	}
	return nil
}

// A list of events.
type EventsList struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventsList) Reset()         { *m = EventsList{} }
func (m *EventsList) String() string { return proto.CompactTextString(m) }
func (*EventsList) ProtoMessage()    {}
func (*EventsList) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{2}
}

func (m *EventsList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventsList.Unmarshal(m, b)
}
func (m *EventsList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventsList.Marshal(b, m, deterministic)
}
func (m *EventsList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventsList.Merge(m, src)
}
func (m *EventsList) XXX_Size() int {
	return xxx_messageInfo_EventsList.Size(m)
}
func (m *EventsList) XXX_DiscardUnknown() {
	xxx_messageInfo_EventsList.DiscardUnknown(m)
}

var xxx_messageInfo_EventsList proto.InternalMessageInfo

func (m *EventsList) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

// A list of EventList's, each representing all events for a transaction.
type EventsForVersions struct {
	EventsForVersion     []*EventsList `protobuf:"bytes,1,rep,name=events_for_version,json=eventsForVersion,proto3" json:"events_for_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *EventsForVersions) Reset()         { *m = EventsForVersions{} }
func (m *EventsForVersions) String() string { return proto.CompactTextString(m) }
func (*EventsForVersions) ProtoMessage()    {}
func (*EventsForVersions) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{3}
}

func (m *EventsForVersions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventsForVersions.Unmarshal(m, b)
}
func (m *EventsForVersions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventsForVersions.Marshal(b, m, deterministic)
}
func (m *EventsForVersions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventsForVersions.Merge(m, src)
}
func (m *EventsForVersions) XXX_Size() int {
	return xxx_messageInfo_EventsForVersions.Size(m)
}
func (m *EventsForVersions) XXX_DiscardUnknown() {
	xxx_messageInfo_EventsForVersions.DiscardUnknown(m)
}

var xxx_messageInfo_EventsForVersions proto.InternalMessageInfo

func (m *EventsForVersions) GetEventsForVersion() []*EventsList {
	if m != nil {
		return m.EventsForVersion
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "types.Event")
	proto.RegisterType((*EventWithProof)(nil), "types.EventWithProof")
	proto.RegisterType((*EventsList)(nil), "types.EventsList")
	proto.RegisterType((*EventsForVersions)(nil), "types.EventsForVersions")
}

func init() {
	proto.RegisterFile("events.proto", fileDescriptor_8f22242cb04491f9)
}

var fileDescriptor_8f22242cb04491f9 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xd1, 0x4a, 0x33, 0x31,
	0x10, 0x85, 0xd9, 0x7f, 0xbb, 0xfd, 0x75, 0xb6, 0xd4, 0x76, 0xbc, 0x59, 0x0b, 0x62, 0x59, 0x84,
	0xf6, 0xaa, 0x42, 0x7d, 0x00, 0x6f, 0x54, 0x10, 0x44, 0x64, 0x29, 0x7a, 0x19, 0xd2, 0xed, 0xb4,
	0x5d, 0xc4, 0x64, 0x4d, 0xd2, 0x62, 0xc1, 0x77, 0xf2, 0x15, 0x25, 0x93, 0x0a, 0x4b, 0xbd, 0xcb,
	0x9c, 0x93, 0x99, 0xf3, 0x65, 0x02, 0x1d, 0xda, 0x92, 0x72, 0x76, 0x52, 0x1b, 0xed, 0x34, 0x26,
	0x6e, 0x57, 0x93, 0x1d, 0xf4, 0x65, 0x59, 0x92, 0xb5, 0xa2, 0x96, 0x6e, 0x1d, 0x9c, 0x41, 0x5a,
	0x1b, 0xad, 0x97, 0xa1, 0xc8, 0xbf, 0x20, 0xb9, 0xf3, 0x6d, 0xd8, 0x83, 0xf8, 0x8d, 0x76, 0x59,
	0x34, 0x8c, 0xc6, 0x9d, 0xc2, 0x1f, 0x71, 0x04, 0x27, 0x96, 0x3e, 0x36, 0xa4, 0x4a, 0x12, 0x6a,
	0xf3, 0x3e, 0x27, 0x93, 0xfd, 0x1b, 0x46, 0xe3, 0x56, 0xd1, 0xfd, 0x95, 0x9f, 0x58, 0xc5, 0x73,
	0x00, 0x8e, 0x16, 0x0b, 0xe9, 0x64, 0x16, 0xf3, 0x84, 0x63, 0x56, 0x6e, 0xa5, 0x93, 0x78, 0x06,
	0x47, 0x9e, 0x45, 0x38, 0xb9, 0xca, 0x5a, 0x6c, 0xfe, 0xf7, 0xf5, 0x4c, 0xae, 0xf2, 0xef, 0x08,
	0xba, 0x1c, 0xff, 0x5a, 0xb9, 0xf5, 0xb3, 0xc7, 0xc2, 0x2b, 0x38, 0x75, 0x46, 0x2a, 0x2b, 0x4b,
	0x57, 0x69, 0x25, 0xb6, 0x64, 0x6c, 0xa5, 0x15, 0x73, 0xb5, 0x0a, 0x6c, 0x58, 0x2f, 0xc1, 0xc1,
	0x0b, 0x48, 0x43, 0x7a, 0xa5, 0x16, 0xf4, 0xb9, 0x47, 0x0c, 0x40, 0x0f, 0x5e, 0xc1, 0x1c, 0x12,
	0xae, 0x98, 0x2c, 0x9d, 0x76, 0x26, 0xbc, 0x99, 0x09, 0xe7, 0x16, 0xc1, 0xc2, 0x11, 0x24, 0xbc,
	0x15, 0x06, 0x4c, 0xa7, 0xfd, 0xe6, 0x1d, 0xe6, 0x2a, 0x82, 0x9f, 0x4f, 0x01, 0x58, 0xb4, 0x8f,
	0x95, 0x75, 0x78, 0x09, 0xed, 0xb0, 0xf4, 0x2c, 0x1a, 0xc6, 0x7f, 0x66, 0xef, 0xbd, 0x7c, 0x06,
	0xfd, 0xd0, 0x73, 0xaf, 0xcd, 0x9e, 0xda, 0xe2, 0x0d, 0x60, 0xb0, 0xc5, 0x52, 0x9b, 0xc6, 0x33,
	0xe3, 0xc3, 0x78, 0x4e, 0x2a, 0x7a, 0x74, 0x30, 0x61, 0xde, 0xe6, 0x0f, 0xbc, 0xfe, 0x09, 0x00,
	0x00, 0xff, 0xff, 0xf0, 0xea, 0x79, 0xcb, 0xf7, 0x01, 0x00, 0x00,
}
