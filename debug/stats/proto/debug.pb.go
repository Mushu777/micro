// Code generated by protoc-gen-go. DO NOT EDIT.
// source: micro/micro/debug/stats/proto/debug.proto

package go_micro_debug_stats

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

// Service describes a service running in the micro network.
type Service struct {
	// Service name, e.g. go.micro.service.greeter
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Node                 *Node    `protobuf:"bytes,3,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c744ea44250254b, []int{0}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Service) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// Node describes a single instance of a service.
type Node struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c744ea44250254b, []int{1}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// Snapshot is a snapshot of Debug.Stats from a particular service when called.
type Snapshot struct {
	Service *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// Unix timestamp, e.g. 1575561487
	Started int64 `protobuf:"varint,2,opt,name=started,proto3" json:"started,omitempty"`
	// Uptime in seconds
	Uptime uint64 `protobuf:"varint,3,opt,name=uptime,proto3" json:"uptime,omitempty"`
	// Heap allocated in bytes (TODO: change to resident set size)
	Memory uint64 `protobuf:"varint,4,opt,name=memory,proto3" json:"memory,omitempty"`
	// Number of Goroutines
	Threads uint64 `protobuf:"varint,5,opt,name=threads,proto3" json:"threads,omitempty"`
	// GC Pause total in ns
	Gc                   uint64   `protobuf:"varint,6,opt,name=gc,proto3" json:"gc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c744ea44250254b, []int{2}
}

func (m *Snapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Snapshot.Unmarshal(m, b)
}
func (m *Snapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Snapshot.Marshal(b, m, deterministic)
}
func (m *Snapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Snapshot.Merge(m, src)
}
func (m *Snapshot) XXX_Size() int {
	return xxx_messageInfo_Snapshot.Size(m)
}
func (m *Snapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_Snapshot.DiscardUnknown(m)
}

var xxx_messageInfo_Snapshot proto.InternalMessageInfo

func (m *Snapshot) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *Snapshot) GetStarted() int64 {
	if m != nil {
		return m.Started
	}
	return 0
}

func (m *Snapshot) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *Snapshot) GetMemory() uint64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *Snapshot) GetThreads() uint64 {
	if m != nil {
		return m.Threads
	}
	return 0
}

func (m *Snapshot) GetGc() uint64 {
	if m != nil {
		return m.Gc
	}
	return 0
}

type ReadRequest struct {
	// If set, only return services matching the filter
	Service              *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c744ea44250254b, []int{3}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type ReadResponse struct {
	Stats                []*Snapshot `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c744ea44250254b, []int{4}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetStats() []*Snapshot {
	if m != nil {
		return m.Stats
	}
	return nil
}

type WriteRequest struct {
	// If set, only return services matching the filter
	Service *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// snapshot to write
	Stats                *Snapshot `protobuf:"bytes,2,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c744ea44250254b, []int{5}
}

func (m *WriteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRequest.Unmarshal(m, b)
}
func (m *WriteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRequest.Marshal(b, m, deterministic)
}
func (m *WriteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRequest.Merge(m, src)
}
func (m *WriteRequest) XXX_Size() int {
	return xxx_messageInfo_WriteRequest.Size(m)
}
func (m *WriteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRequest proto.InternalMessageInfo

func (m *WriteRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *WriteRequest) GetStats() *Snapshot {
	if m != nil {
		return m.Stats
	}
	return nil
}

type WriteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteResponse) Reset()         { *m = WriteResponse{} }
func (m *WriteResponse) String() string { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()    {}
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c744ea44250254b, []int{6}
}

func (m *WriteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteResponse.Unmarshal(m, b)
}
func (m *WriteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteResponse.Marshal(b, m, deterministic)
}
func (m *WriteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteResponse.Merge(m, src)
}
func (m *WriteResponse) XXX_Size() int {
	return xxx_messageInfo_WriteResponse.Size(m)
}
func (m *WriteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WriteResponse proto.InternalMessageInfo

type StreamRequest struct {
	// If set, only return services matching the filter
	Service *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// If set, only return services matching the namespace
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c744ea44250254b, []int{7}
}

func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (m *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(m, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *StreamRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type StreamResponse struct {
	Stats                []*Snapshot `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *StreamResponse) Reset()         { *m = StreamResponse{} }
func (m *StreamResponse) String() string { return proto.CompactTextString(m) }
func (*StreamResponse) ProtoMessage()    {}
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c744ea44250254b, []int{8}
}

func (m *StreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse.Unmarshal(m, b)
}
func (m *StreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse.Marshal(b, m, deterministic)
}
func (m *StreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse.Merge(m, src)
}
func (m *StreamResponse) XXX_Size() int {
	return xxx_messageInfo_StreamResponse.Size(m)
}
func (m *StreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse proto.InternalMessageInfo

func (m *StreamResponse) GetStats() []*Snapshot {
	if m != nil {
		return m.Stats
	}
	return nil
}

func init() {
	proto.RegisterType((*Service)(nil), "go.micro.debug.stats.Service")
	proto.RegisterType((*Node)(nil), "go.micro.debug.stats.Node")
	proto.RegisterType((*Snapshot)(nil), "go.micro.debug.stats.Snapshot")
	proto.RegisterType((*ReadRequest)(nil), "go.micro.debug.stats.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "go.micro.debug.stats.ReadResponse")
	proto.RegisterType((*WriteRequest)(nil), "go.micro.debug.stats.WriteRequest")
	proto.RegisterType((*WriteResponse)(nil), "go.micro.debug.stats.WriteResponse")
	proto.RegisterType((*StreamRequest)(nil), "go.micro.debug.stats.StreamRequest")
	proto.RegisterType((*StreamResponse)(nil), "go.micro.debug.stats.StreamResponse")
}

func init() {
	proto.RegisterFile("micro/micro/debug/stats/proto/debug.proto", fileDescriptor_1c744ea44250254b)
}

var fileDescriptor_1c744ea44250254b = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x8a, 0xd4, 0x30,
	0x18, 0xa5, 0x9d, 0x4e, 0xc7, 0xfd, 0xba, 0xbb, 0x42, 0x10, 0x09, 0x45, 0x65, 0x8c, 0x5e, 0x8c,
	0x37, 0x9d, 0xa5, 0x0a, 0xbe, 0x80, 0xec, 0x9d, 0x22, 0x29, 0xe2, 0x75, 0xb6, 0xf9, 0xec, 0xf6,
	0xa2, 0x4d, 0x4d, 0x32, 0x0b, 0x5e, 0xf8, 0x5c, 0x3e, 0x9e, 0x92, 0x9f, 0xb2, 0x2b, 0xcc, 0x0c,
	0xe2, 0xdc, 0x0c, 0x39, 0x67, 0x4e, 0xce, 0xf9, 0x7e, 0x42, 0xe1, 0xcd, 0xd0, 0xb7, 0x5a, 0x6d,
	0xc3, 0xaf, 0xc4, 0x9b, 0x5d, 0xb7, 0x35, 0x56, 0x58, 0xb3, 0x9d, 0xb4, 0xb2, 0x91, 0xa9, 0xfc,
	0x99, 0x3c, 0xe9, 0x54, 0xe5, 0x75, 0x55, 0x60, 0xbd, 0x8e, 0x75, 0xb0, 0x6a, 0x50, 0xdf, 0xf5,
	0x2d, 0x12, 0x02, 0xd9, 0x28, 0x06, 0xa4, 0xc9, 0x3a, 0xd9, 0x9c, 0x71, 0x7f, 0x26, 0x14, 0x56,
	0x77, 0xa8, 0x4d, 0xaf, 0x46, 0x9a, 0x7a, 0x7a, 0x86, 0xa4, 0x82, 0x6c, 0x54, 0x12, 0xe9, 0x62,
	0x9d, 0x6c, 0x8a, 0xba, 0xac, 0xf6, 0xb9, 0x57, 0x9f, 0x94, 0x44, 0xee, 0x75, 0xec, 0x0a, 0x32,
	0x87, 0xc8, 0x25, 0xa4, 0xbd, 0x8c, 0x19, 0x69, 0x2f, 0x5d, 0x82, 0x90, 0x52, 0xa3, 0x31, 0x73,
	0x42, 0x84, 0xec, 0x57, 0x02, 0x8f, 0x9a, 0x51, 0x4c, 0xe6, 0x56, 0x59, 0xf2, 0x1e, 0x56, 0x26,
	0xd4, 0xe9, 0xef, 0x16, 0xf5, 0xf3, 0xfd, 0x89, 0xb1, 0x19, 0x3e, 0xab, 0x9d, 0xbf, 0xb1, 0x42,
	0x5b, 0x94, 0xde, 0x7f, 0xc1, 0x67, 0x48, 0x9e, 0x42, 0xbe, 0x9b, 0x6c, 0x3f, 0x84, 0x1e, 0x32,
	0x1e, 0x91, 0xe3, 0x07, 0x1c, 0x94, 0xfe, 0x41, 0xb3, 0xc0, 0x07, 0xe4, 0x9c, 0xec, 0xad, 0x46,
	0x21, 0x0d, 0x5d, 0xfa, 0x3f, 0x66, 0xe8, 0x7a, 0xea, 0x5a, 0x9a, 0x7b, 0x32, 0xed, 0x5a, 0x76,
	0x0d, 0x05, 0x47, 0x21, 0x39, 0x7e, 0xdf, 0xa1, 0xf9, 0xff, 0xda, 0xd9, 0x07, 0x38, 0x0f, 0x3e,
	0x66, 0x52, 0xa3, 0x41, 0xf2, 0x0e, 0x96, 0x5e, 0x49, 0x93, 0xf5, 0x62, 0x53, 0xd4, 0x2f, 0x0e,
	0xd8, 0xc4, 0x99, 0xf1, 0x20, 0x66, 0x3f, 0xe1, 0xfc, 0xab, 0xee, 0x2d, 0x9e, 0x5a, 0xce, 0x7d,
	0x7c, 0xea, 0xaf, 0xfd, 0x63, 0xfc, 0x63, 0xb8, 0x88, 0xf1, 0xa1, 0x0b, 0xf6, 0x0d, 0x2e, 0x1a,
	0xab, 0x51, 0x0c, 0x27, 0x17, 0xf4, 0x0c, 0xce, 0xdc, 0x2b, 0x35, 0x93, 0x68, 0x31, 0xbe, 0x9e,
	0x7b, 0x82, 0x5d, 0xc3, 0xe5, 0x9c, 0x73, 0xca, 0xfc, 0xea, 0xdf, 0x09, 0x2c, 0x1b, 0x77, 0x22,
	0x1f, 0x21, 0x73, 0xfb, 0x20, 0x2f, 0xf7, 0x5f, 0x7c, 0xb0, 0xf3, 0x92, 0x1d, 0x93, 0xc4, 0x72,
	0x3e, 0xc3, 0xd2, 0x4f, 0x86, 0x1c, 0x10, 0x3f, 0xdc, 0x5a, 0xf9, 0xea, 0xa8, 0x26, 0x3a, 0x7e,
	0x81, 0x3c, 0xb4, 0x4c, 0x0e, 0xc8, 0xff, 0x1a, 0x7c, 0xf9, 0xfa, 0xb8, 0x28, 0x98, 0x5e, 0x25,
	0x37, 0xb9, 0xff, 0x82, 0xbc, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0x34, 0x92, 0x28, 0x08, 0x6e,
	0x04, 0x00, 0x00,
}