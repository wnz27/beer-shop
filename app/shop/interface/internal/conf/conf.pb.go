// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/conf/conf.proto

package conf // import "shop/interface/internal/conf"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import durationpb "google.golang.org/protobuf/types/known/durationpb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Bootstrap struct {
	Trace                *Trace   `protobuf:"bytes,1,opt,name=trace,proto3" json:"trace,omitempty"`
	Server               *Server  `protobuf:"bytes,2,opt,name=server,proto3" json:"server,omitempty"`
	Data                 *Data    `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Auth                 *Auth    `protobuf:"bytes,4,opt,name=auth,proto3" json:"auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bootstrap) Reset()         { *m = Bootstrap{} }
func (m *Bootstrap) String() string { return proto.CompactTextString(m) }
func (*Bootstrap) ProtoMessage()    {}
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return fileDescriptor_conf_5ff28fd32eaf8ea8, []int{0}
}
func (m *Bootstrap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bootstrap.Unmarshal(m, b)
}
func (m *Bootstrap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bootstrap.Marshal(b, m, deterministic)
}
func (dst *Bootstrap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bootstrap.Merge(dst, src)
}
func (m *Bootstrap) XXX_Size() int {
	return xxx_messageInfo_Bootstrap.Size(m)
}
func (m *Bootstrap) XXX_DiscardUnknown() {
	xxx_messageInfo_Bootstrap.DiscardUnknown(m)
}

var xxx_messageInfo_Bootstrap proto.InternalMessageInfo

func (m *Bootstrap) GetTrace() *Trace {
	if m != nil {
		return m.Trace
	}
	return nil
}

func (m *Bootstrap) GetServer() *Server {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *Bootstrap) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Bootstrap) GetAuth() *Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

type Trace struct {
	Endpoint             string   `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Trace) Reset()         { *m = Trace{} }
func (m *Trace) String() string { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()    {}
func (*Trace) Descriptor() ([]byte, []int) {
	return fileDescriptor_conf_5ff28fd32eaf8ea8, []int{1}
}
func (m *Trace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trace.Unmarshal(m, b)
}
func (m *Trace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trace.Marshal(b, m, deterministic)
}
func (dst *Trace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trace.Merge(dst, src)
}
func (m *Trace) XXX_Size() int {
	return xxx_messageInfo_Trace.Size(m)
}
func (m *Trace) XXX_DiscardUnknown() {
	xxx_messageInfo_Trace.DiscardUnknown(m)
}

var xxx_messageInfo_Trace proto.InternalMessageInfo

func (m *Trace) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

type Server struct {
	Http                 *Server_HTTP `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Grpc                 *Server_GRPC `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Server) Reset()         { *m = Server{} }
func (m *Server) String() string { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()    {}
func (*Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_conf_5ff28fd32eaf8ea8, []int{2}
}
func (m *Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Server.Unmarshal(m, b)
}
func (m *Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Server.Marshal(b, m, deterministic)
}
func (dst *Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server.Merge(dst, src)
}
func (m *Server) XXX_Size() int {
	return xxx_messageInfo_Server.Size(m)
}
func (m *Server) XXX_DiscardUnknown() {
	xxx_messageInfo_Server.DiscardUnknown(m)
}

var xxx_messageInfo_Server proto.InternalMessageInfo

func (m *Server) GetHttp() *Server_HTTP {
	if m != nil {
		return m.Http
	}
	return nil
}

func (m *Server) GetGrpc() *Server_GRPC {
	if m != nil {
		return m.Grpc
	}
	return nil
}

type Server_HTTP struct {
	Network              string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr                 string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout              *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Server_HTTP) Reset()         { *m = Server_HTTP{} }
func (m *Server_HTTP) String() string { return proto.CompactTextString(m) }
func (*Server_HTTP) ProtoMessage()    {}
func (*Server_HTTP) Descriptor() ([]byte, []int) {
	return fileDescriptor_conf_5ff28fd32eaf8ea8, []int{2, 0}
}
func (m *Server_HTTP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Server_HTTP.Unmarshal(m, b)
}
func (m *Server_HTTP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Server_HTTP.Marshal(b, m, deterministic)
}
func (dst *Server_HTTP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server_HTTP.Merge(dst, src)
}
func (m *Server_HTTP) XXX_Size() int {
	return xxx_messageInfo_Server_HTTP.Size(m)
}
func (m *Server_HTTP) XXX_DiscardUnknown() {
	xxx_messageInfo_Server_HTTP.DiscardUnknown(m)
}

var xxx_messageInfo_Server_HTTP proto.InternalMessageInfo

func (m *Server_HTTP) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *Server_HTTP) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Server_HTTP) GetTimeout() *durationpb.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

type Server_GRPC struct {
	Network              string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr                 string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout              *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Server_GRPC) Reset()         { *m = Server_GRPC{} }
func (m *Server_GRPC) String() string { return proto.CompactTextString(m) }
func (*Server_GRPC) ProtoMessage()    {}
func (*Server_GRPC) Descriptor() ([]byte, []int) {
	return fileDescriptor_conf_5ff28fd32eaf8ea8, []int{2, 1}
}
func (m *Server_GRPC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Server_GRPC.Unmarshal(m, b)
}
func (m *Server_GRPC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Server_GRPC.Marshal(b, m, deterministic)
}
func (dst *Server_GRPC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server_GRPC.Merge(dst, src)
}
func (m *Server_GRPC) XXX_Size() int {
	return xxx_messageInfo_Server_GRPC.Size(m)
}
func (m *Server_GRPC) XXX_DiscardUnknown() {
	xxx_messageInfo_Server_GRPC.DiscardUnknown(m)
}

var xxx_messageInfo_Server_GRPC proto.InternalMessageInfo

func (m *Server_GRPC) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *Server_GRPC) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Server_GRPC) GetTimeout() *durationpb.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

type Data struct {
	Database             *Data_Database `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_conf_5ff28fd32eaf8ea8, []int{3}
}
func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (dst *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(dst, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetDatabase() *Data_Database {
	if m != nil {
		return m.Database
	}
	return nil
}

type Data_Database struct {
	Driver               string   `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Source               string   `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data_Database) Reset()         { *m = Data_Database{} }
func (m *Data_Database) String() string { return proto.CompactTextString(m) }
func (*Data_Database) ProtoMessage()    {}
func (*Data_Database) Descriptor() ([]byte, []int) {
	return fileDescriptor_conf_5ff28fd32eaf8ea8, []int{3, 0}
}
func (m *Data_Database) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data_Database.Unmarshal(m, b)
}
func (m *Data_Database) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data_Database.Marshal(b, m, deterministic)
}
func (dst *Data_Database) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data_Database.Merge(dst, src)
}
func (m *Data_Database) XXX_Size() int {
	return xxx_messageInfo_Data_Database.Size(m)
}
func (m *Data_Database) XXX_DiscardUnknown() {
	xxx_messageInfo_Data_Database.DiscardUnknown(m)
}

var xxx_messageInfo_Data_Database proto.InternalMessageInfo

func (m *Data_Database) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *Data_Database) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

type Data_Redis struct {
	Network              string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr                 string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	ReadTimeout          *durationpb.Duration `protobuf:"bytes,3,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`
	WriteTimeout         *durationpb.Duration `protobuf:"bytes,4,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Data_Redis) Reset()         { *m = Data_Redis{} }
func (m *Data_Redis) String() string { return proto.CompactTextString(m) }
func (*Data_Redis) ProtoMessage()    {}
func (*Data_Redis) Descriptor() ([]byte, []int) {
	return fileDescriptor_conf_5ff28fd32eaf8ea8, []int{3, 1}
}
func (m *Data_Redis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data_Redis.Unmarshal(m, b)
}
func (m *Data_Redis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data_Redis.Marshal(b, m, deterministic)
}
func (dst *Data_Redis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data_Redis.Merge(dst, src)
}
func (m *Data_Redis) XXX_Size() int {
	return xxx_messageInfo_Data_Redis.Size(m)
}
func (m *Data_Redis) XXX_DiscardUnknown() {
	xxx_messageInfo_Data_Redis.DiscardUnknown(m)
}

var xxx_messageInfo_Data_Redis proto.InternalMessageInfo

func (m *Data_Redis) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *Data_Redis) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Data_Redis) GetReadTimeout() *durationpb.Duration {
	if m != nil {
		return m.ReadTimeout
	}
	return nil
}

func (m *Data_Redis) GetWriteTimeout() *durationpb.Duration {
	if m != nil {
		return m.WriteTimeout
	}
	return nil
}

type Registry struct {
	Consul               *Registry_Consul `protobuf:"bytes,1,opt,name=consul,proto3" json:"consul,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Registry) Reset()         { *m = Registry{} }
func (m *Registry) String() string { return proto.CompactTextString(m) }
func (*Registry) ProtoMessage()    {}
func (*Registry) Descriptor() ([]byte, []int) {
	return fileDescriptor_conf_5ff28fd32eaf8ea8, []int{4}
}
func (m *Registry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Registry.Unmarshal(m, b)
}
func (m *Registry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Registry.Marshal(b, m, deterministic)
}
func (dst *Registry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Registry.Merge(dst, src)
}
func (m *Registry) XXX_Size() int {
	return xxx_messageInfo_Registry.Size(m)
}
func (m *Registry) XXX_DiscardUnknown() {
	xxx_messageInfo_Registry.DiscardUnknown(m)
}

var xxx_messageInfo_Registry proto.InternalMessageInfo

func (m *Registry) GetConsul() *Registry_Consul {
	if m != nil {
		return m.Consul
	}
	return nil
}

type Registry_Consul struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Scheme               string   `protobuf:"bytes,2,opt,name=scheme,proto3" json:"scheme,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Registry_Consul) Reset()         { *m = Registry_Consul{} }
func (m *Registry_Consul) String() string { return proto.CompactTextString(m) }
func (*Registry_Consul) ProtoMessage()    {}
func (*Registry_Consul) Descriptor() ([]byte, []int) {
	return fileDescriptor_conf_5ff28fd32eaf8ea8, []int{4, 0}
}
func (m *Registry_Consul) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Registry_Consul.Unmarshal(m, b)
}
func (m *Registry_Consul) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Registry_Consul.Marshal(b, m, deterministic)
}
func (dst *Registry_Consul) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Registry_Consul.Merge(dst, src)
}
func (m *Registry_Consul) XXX_Size() int {
	return xxx_messageInfo_Registry_Consul.Size(m)
}
func (m *Registry_Consul) XXX_DiscardUnknown() {
	xxx_messageInfo_Registry_Consul.DiscardUnknown(m)
}

var xxx_messageInfo_Registry_Consul proto.InternalMessageInfo

func (m *Registry_Consul) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Registry_Consul) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

type Auth struct {
	ServiceKey           string   `protobuf:"bytes,1,opt,name=service_key,json=serviceKey,proto3" json:"service_key,omitempty"`
	ApiKey               string   `protobuf:"bytes,2,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth) Reset()         { *m = Auth{} }
func (m *Auth) String() string { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()    {}
func (*Auth) Descriptor() ([]byte, []int) {
	return fileDescriptor_conf_5ff28fd32eaf8ea8, []int{5}
}
func (m *Auth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth.Unmarshal(m, b)
}
func (m *Auth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth.Marshal(b, m, deterministic)
}
func (dst *Auth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth.Merge(dst, src)
}
func (m *Auth) XXX_Size() int {
	return xxx_messageInfo_Auth.Size(m)
}
func (m *Auth) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth.DiscardUnknown(m)
}

var xxx_messageInfo_Auth proto.InternalMessageInfo

func (m *Auth) GetServiceKey() string {
	if m != nil {
		return m.ServiceKey
	}
	return ""
}

func (m *Auth) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func init() {
	proto.RegisterType((*Bootstrap)(nil), "kratos.api.Bootstrap")
	proto.RegisterType((*Trace)(nil), "kratos.api.Trace")
	proto.RegisterType((*Server)(nil), "kratos.api.Server")
	proto.RegisterType((*Server_HTTP)(nil), "kratos.api.Server.HTTP")
	proto.RegisterType((*Server_GRPC)(nil), "kratos.api.Server.GRPC")
	proto.RegisterType((*Data)(nil), "kratos.api.Data")
	proto.RegisterType((*Data_Database)(nil), "kratos.api.Data.Database")
	proto.RegisterType((*Data_Redis)(nil), "kratos.api.Data.Redis")
	proto.RegisterType((*Registry)(nil), "kratos.api.Registry")
	proto.RegisterType((*Registry_Consul)(nil), "kratos.api.Registry.Consul")
	proto.RegisterType((*Auth)(nil), "kratos.api.Auth")
}

func init() { proto.RegisterFile("internal/conf/conf.proto", fileDescriptor_conf_5ff28fd32eaf8ea8) }

var fileDescriptor_conf_5ff28fd32eaf8ea8 = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xdf, 0x6b, 0xd4, 0x40,
	0x10, 0xc7, 0xb9, 0x9a, 0xcb, 0xdd, 0xcd, 0x55, 0xd0, 0x7d, 0xb0, 0x69, 0x04, 0x7f, 0xa4, 0x82,
	0xa2, 0x90, 0x80, 0x87, 0x2f, 0x55, 0x44, 0xdb, 0x82, 0x42, 0x5f, 0xca, 0x7a, 0x4f, 0xbe, 0x94,
	0xbd, 0x64, 0xee, 0xb2, 0xdc, 0x35, 0x1b, 0x76, 0x37, 0x96, 0xc3, 0x7f, 0x49, 0xf0, 0xc1, 0x17,
	0xff, 0x3c, 0xd9, 0xc9, 0xe6, 0x7a, 0x5a, 0x41, 0x7d, 0xf1, 0x65, 0xc9, 0xcc, 0x7c, 0x26, 0x3b,
	0xf3, 0x9d, 0xdd, 0x85, 0x48, 0x56, 0x16, 0x75, 0x25, 0x56, 0x59, 0xae, 0xaa, 0x39, 0x2d, 0x69,
	0xad, 0x95, 0x55, 0x0c, 0x96, 0x5a, 0x58, 0x65, 0x52, 0x51, 0xcb, 0xf8, 0xde, 0x42, 0xa9, 0xc5,
	0x0a, 0x33, 0x8a, 0xcc, 0x9a, 0x79, 0x56, 0x34, 0x5a, 0x58, 0xa9, 0xaa, 0x96, 0x4d, 0xbe, 0xf4,
	0x60, 0x74, 0xa4, 0x94, 0x35, 0x56, 0x8b, 0x9a, 0x3d, 0x86, 0xbe, 0xd5, 0x22, 0xc7, 0xa8, 0xf7,
	0xa0, 0xf7, 0x64, 0xfc, 0xfc, 0x76, 0x7a, 0xf5, 0xa7, 0x74, 0xea, 0x02, 0xbc, 0x8d, 0xb3, 0xa7,
	0x10, 0x1a, 0xd4, 0x9f, 0x50, 0x47, 0x3b, 0x44, 0xb2, 0x6d, 0xf2, 0x03, 0x45, 0xb8, 0x27, 0xd8,
	0x23, 0x08, 0x0a, 0x61, 0x45, 0x74, 0x83, 0xc8, 0x5b, 0xdb, 0xe4, 0x89, 0xb0, 0x82, 0x53, 0xd4,
	0x51, 0xa2, 0xb1, 0x65, 0x14, 0x5c, 0xa7, 0xde, 0x36, 0xb6, 0xe4, 0x14, 0x4d, 0x0e, 0xa0, 0x4f,
	0x75, 0xb0, 0x18, 0x86, 0x58, 0x15, 0xb5, 0x92, 0x95, 0xa5, 0x62, 0x47, 0x7c, 0x63, 0x27, 0xdf,
	0x77, 0x20, 0x6c, 0x6b, 0x60, 0xcf, 0x20, 0x28, 0xad, 0xad, 0x7d, 0x3f, 0x7b, 0xd7, 0xab, 0x4c,
	0xdf, 0x4f, 0xa7, 0x67, 0x9c, 0x20, 0x07, 0x2f, 0x74, 0x9d, 0xfb, 0x96, 0x7e, 0x07, 0xbf, 0xe3,
	0x67, 0xc7, 0x9c, 0xa0, 0x58, 0x42, 0xe0, 0x52, 0x59, 0x04, 0x83, 0x0a, 0xed, 0xa5, 0xd2, 0x4b,
	0x5f, 0x47, 0x67, 0x32, 0x06, 0x81, 0x28, 0x8a, 0x56, 0xa1, 0x11, 0xa7, 0x6f, 0x36, 0x81, 0x81,
	0x95, 0x17, 0xa8, 0x1a, 0xeb, 0xe5, 0xd8, 0x4f, 0xdb, 0x01, 0xa5, 0xdd, 0x80, 0xd2, 0x13, 0x3f,
	0x20, 0xde, 0x91, 0x6e, 0x2b, 0xb7, 0xf1, 0x7f, 0xd8, 0x2a, 0xf9, 0xba, 0x03, 0x81, 0x1b, 0x0a,
	0x7b, 0x01, 0x43, 0x37, 0x96, 0x99, 0x30, 0xdd, 0x61, 0xd8, 0xff, 0x75, 0x70, 0xb4, 0x38, 0x80,
	0x6f, 0xd0, 0xf8, 0x10, 0x86, 0x9d, 0x97, 0xdd, 0x81, 0xb0, 0xd0, 0xd2, 0x9d, 0x91, 0xb6, 0x5a,
	0x6f, 0x39, 0xbf, 0x51, 0x8d, 0xce, 0xd1, 0x97, 0xeb, 0xad, 0xf8, 0x5b, 0x0f, 0xfa, 0x1c, 0x0b,
	0x69, 0xfe, 0xb1, 0xd1, 0x57, 0xb0, 0xab, 0x51, 0x14, 0xe7, 0x7f, 0xdd, 0xed, 0xd8, 0xe1, 0xd3,
	0x96, 0x66, 0xaf, 0xe1, 0xe6, 0xa5, 0x96, 0x16, 0x37, 0xe9, 0xc1, 0x9f, 0xd2, 0x77, 0x89, 0xf7,
	0xf9, 0xc9, 0x67, 0x18, 0x72, 0x5c, 0x48, 0x63, 0xf5, 0x9a, 0x4d, 0x20, 0xcc, 0x55, 0x65, 0x9a,
	0x95, 0x97, 0xec, 0xee, 0xb6, 0x64, 0x1d, 0x95, 0x1e, 0x13, 0xc2, 0x3d, 0x1a, 0x1f, 0x42, 0xd8,
	0x7a, 0x5c, 0xdb, 0xae, 0x21, 0x34, 0xa6, 0x6b, 0xdb, 0x9b, 0x24, 0x59, 0x5e, 0xe2, 0xc5, 0x95,
	0x64, 0x64, 0x25, 0x6f, 0x20, 0x70, 0x97, 0x83, 0xdd, 0x87, 0xb1, 0xbb, 0x6c, 0x32, 0xc7, 0xf3,
	0x25, 0xae, 0x7d, 0x36, 0x78, 0xd7, 0x29, 0xae, 0xd9, 0x1e, 0x0c, 0x44, 0x2d, 0x29, 0xe8, 0xff,
	0x20, 0x6a, 0x79, 0x8a, 0xeb, 0xa3, 0x83, 0x8f, 0x0f, 0x4d, 0xa9, 0xea, 0x8c, 0x1e, 0x93, 0xb9,
	0xc8, 0x31, 0xfb, 0xe9, 0x59, 0x79, 0xe9, 0x96, 0x59, 0x48, 0x22, 0x4c, 0x7e, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x44, 0x1f, 0x65, 0x13, 0x73, 0x04, 0x00, 0x00,
}
