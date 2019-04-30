// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transform.proto

package transform

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type DecodeRequest_Type int32

const (
	DecodeRequest_FROM_CACHE DecodeRequest_Type = 0
	DecodeRequest_FROM_BYTES DecodeRequest_Type = 1
	DecodeRequest_FROM_URL   DecodeRequest_Type = 2
)

var DecodeRequest_Type_name = map[int32]string{
	0: "FROM_CACHE",
	1: "FROM_BYTES",
	2: "FROM_URL",
}

var DecodeRequest_Type_value = map[string]int32{
	"FROM_CACHE": 0,
	"FROM_BYTES": 1,
	"FROM_URL":   2,
}

func (x DecodeRequest_Type) String() string {
	return proto.EnumName(DecodeRequest_Type_name, int32(x))
}

func (DecodeRequest_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cb4a498eeb2ba07d, []int{1, 0}
}

type Response struct {
	ContentType          string   `protobuf:"bytes,1,opt,name=contentType,proto3" json:"contentType,omitempty"`
	ContentLength        int64    `protobuf:"varint,2,opt,name=contentLength,proto3" json:"contentLength,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb4a498eeb2ba07d, []int{0}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Response) GetContentLength() int64 {
	if m != nil {
		return m.ContentLength
	}
	return 0
}

func (m *Response) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type DecodeRequest struct {
	Data                 []byte             `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Filename             string             `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Url                  string             `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	From                 string             `protobuf:"bytes,4,opt,name=from,proto3" json:"from,omitempty"`
	To                   string             `protobuf:"bytes,5,opt,name=to,proto3" json:"to,omitempty"`
	Type                 DecodeRequest_Type `protobuf:"varint,6,opt,name=type,proto3,enum=transform.DecodeRequest_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DecodeRequest) Reset()         { *m = DecodeRequest{} }
func (m *DecodeRequest) String() string { return proto.CompactTextString(m) }
func (*DecodeRequest) ProtoMessage()    {}
func (*DecodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb4a498eeb2ba07d, []int{1}
}

func (m *DecodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecodeRequest.Unmarshal(m, b)
}
func (m *DecodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecodeRequest.Marshal(b, m, deterministic)
}
func (m *DecodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecodeRequest.Merge(m, src)
}
func (m *DecodeRequest) XXX_Size() int {
	return xxx_messageInfo_DecodeRequest.Size(m)
}
func (m *DecodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DecodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DecodeRequest proto.InternalMessageInfo

func (m *DecodeRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *DecodeRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *DecodeRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *DecodeRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *DecodeRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *DecodeRequest) GetType() DecodeRequest_Type {
	if m != nil {
		return m.Type
	}
	return DecodeRequest_FROM_CACHE
}

type ThumbnailRequest struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Width                int32    `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height               int32    `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThumbnailRequest) Reset()         { *m = ThumbnailRequest{} }
func (m *ThumbnailRequest) String() string { return proto.CompactTextString(m) }
func (*ThumbnailRequest) ProtoMessage()    {}
func (*ThumbnailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb4a498eeb2ba07d, []int{2}
}

func (m *ThumbnailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThumbnailRequest.Unmarshal(m, b)
}
func (m *ThumbnailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThumbnailRequest.Marshal(b, m, deterministic)
}
func (m *ThumbnailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThumbnailRequest.Merge(m, src)
}
func (m *ThumbnailRequest) XXX_Size() int {
	return xxx_messageInfo_ThumbnailRequest.Size(m)
}
func (m *ThumbnailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ThumbnailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ThumbnailRequest proto.InternalMessageInfo

func (m *ThumbnailRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ThumbnailRequest) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *ThumbnailRequest) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterEnum("transform.DecodeRequest_Type", DecodeRequest_Type_name, DecodeRequest_Type_value)
	proto.RegisterType((*Response)(nil), "transform.Response")
	proto.RegisterType((*DecodeRequest)(nil), "transform.DecodeRequest")
	proto.RegisterType((*ThumbnailRequest)(nil), "transform.ThumbnailRequest")
}

func init() { proto.RegisterFile("transform.proto", fileDescriptor_cb4a498eeb2ba07d) }

var fileDescriptor_cb4a498eeb2ba07d = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x4f, 0xc2, 0x40,
	0x10, 0x85, 0xdd, 0xd2, 0x36, 0x74, 0x04, 0x6c, 0x56, 0x63, 0x1a, 0x8c, 0x49, 0xd3, 0x78, 0xe8,
	0x89, 0x44, 0x34, 0xf1, 0xe4, 0x41, 0x11, 0xc3, 0x01, 0x63, 0x5c, 0xea, 0xc1, 0x93, 0x29, 0x30,
	0xa5, 0x4d, 0xe8, 0x2e, 0xb6, 0x8b, 0xc6, 0x3f, 0xeb, 0x6f, 0x31, 0x5d, 0xa0, 0x14, 0xa3, 0xde,
	0xe6, 0xbd, 0x9d, 0x7d, 0x99, 0x6f, 0x76, 0xe1, 0x40, 0x66, 0x21, 0xcf, 0x23, 0x91, 0xa5, 0x9d,
	0x45, 0x26, 0xa4, 0xa0, 0x56, 0x69, 0x78, 0x11, 0xd4, 0x19, 0xe6, 0x0b, 0xc1, 0x73, 0xa4, 0x2e,
	0xec, 0x4f, 0x04, 0x97, 0xc8, 0x65, 0xf0, 0xb9, 0x40, 0x87, 0xb8, 0xc4, 0xb7, 0x58, 0xd5, 0xa2,
	0x67, 0xd0, 0x5c, 0xcb, 0x21, 0xf2, 0x99, 0x8c, 0x1d, 0xcd, 0x25, 0x7e, 0x8d, 0xed, 0x9a, 0x94,
	0x82, 0x3e, 0x0d, 0x65, 0xe8, 0xd4, 0x5c, 0xe2, 0x37, 0x98, 0xaa, 0xbd, 0x2f, 0x02, 0xcd, 0x3b,
	0x9c, 0x88, 0x29, 0x32, 0x7c, 0x5b, 0x62, 0x2e, 0xcb, 0x2e, 0xb2, 0xed, 0xa2, 0x6d, 0xa8, 0x47,
	0xc9, 0x1c, 0x79, 0x98, 0xa2, 0x8a, 0xb6, 0x58, 0xa9, 0xa9, 0x0d, 0xb5, 0x65, 0x36, 0x57, 0xa1,
	0x16, 0x2b, 0xca, 0x22, 0x21, 0xca, 0x44, 0xea, 0xe8, 0xca, 0x52, 0x35, 0x6d, 0x81, 0x26, 0x85,
	0x63, 0x28, 0x47, 0x93, 0x82, 0x9e, 0x83, 0x2e, 0x0b, 0x18, 0xd3, 0x25, 0x7e, 0xab, 0x7b, 0xda,
	0xd9, 0xae, 0x62, 0x67, 0x9a, 0x4e, 0x81, 0xc7, 0x54, 0xab, 0x77, 0x09, 0xba, 0x82, 0x6d, 0x01,
	0xdc, 0xb3, 0xc7, 0x87, 0xd7, 0xde, 0x4d, 0x6f, 0xd0, 0xb7, 0xf7, 0x4a, 0x7d, 0xfb, 0x12, 0xf4,
	0x47, 0x36, 0xa1, 0x0d, 0xa8, 0x2b, 0xfd, 0xcc, 0x86, 0xb6, 0xe6, 0x05, 0x60, 0x07, 0xf1, 0x32,
	0x1d, 0xf3, 0x30, 0x99, 0xff, 0x87, 0x78, 0x04, 0xc6, 0x47, 0x32, 0x5d, 0xaf, 0xce, 0x60, 0x2b,
	0x41, 0x8f, 0xc1, 0x8c, 0x31, 0x99, 0xc5, 0x52, 0xf1, 0x19, 0x6c, 0xad, 0xba, 0x83, 0xcd, 0xd6,
	0x46, 0x98, 0xbd, 0x27, 0x13, 0xa4, 0x57, 0x60, 0xae, 0x0c, 0xea, 0xfc, 0xc5, 0xd2, 0x3e, 0xac,
	0x9c, 0x6c, 0x1e, 0xb7, 0xfb, 0x54, 0x99, 0x6f, 0x13, 0x76, 0x0d, 0x56, 0xe9, 0xd1, 0x93, 0xca,
	0xad, 0x9f, 0x24, 0xbf, 0x46, 0x8e, 0x4d, 0xf5, 0x9b, 0x2e, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xce, 0x7a, 0xf5, 0x94, 0x60, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DecodeServiceClient is the client API for DecodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DecodeServiceClient interface {
	// Create new image in the storage
	Decode(ctx context.Context, in *DecodeRequest, opts ...grpc.CallOption) (*Response, error)
}

type decodeServiceClient struct {
	cc *grpc.ClientConn
}

func NewDecodeServiceClient(cc *grpc.ClientConn) DecodeServiceClient {
	return &decodeServiceClient{cc}
}

func (c *decodeServiceClient) Decode(ctx context.Context, in *DecodeRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/transform.DecodeService/Decode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DecodeServiceServer is the server API for DecodeService service.
type DecodeServiceServer interface {
	// Create new image in the storage
	Decode(context.Context, *DecodeRequest) (*Response, error)
}

// UnimplementedDecodeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDecodeServiceServer struct {
}

func (*UnimplementedDecodeServiceServer) Decode(ctx context.Context, req *DecodeRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decode not implemented")
}

func RegisterDecodeServiceServer(s *grpc.Server, srv DecodeServiceServer) {
	s.RegisterService(&_DecodeService_serviceDesc, srv)
}

func _DecodeService_Decode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecodeServiceServer).Decode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transform.DecodeService/Decode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecodeServiceServer).Decode(ctx, req.(*DecodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DecodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transform.DecodeService",
	HandlerType: (*DecodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Decode",
			Handler:    _DecodeService_Decode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transform.proto",
}

// ThumbnailServiceClient is the client API for ThumbnailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ThumbnailServiceClient interface {
	// Create new image in the storage
	Thumbnail(ctx context.Context, in *ThumbnailRequest, opts ...grpc.CallOption) (*Response, error)
}

type thumbnailServiceClient struct {
	cc *grpc.ClientConn
}

func NewThumbnailServiceClient(cc *grpc.ClientConn) ThumbnailServiceClient {
	return &thumbnailServiceClient{cc}
}

func (c *thumbnailServiceClient) Thumbnail(ctx context.Context, in *ThumbnailRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/transform.ThumbnailService/Thumbnail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThumbnailServiceServer is the server API for ThumbnailService service.
type ThumbnailServiceServer interface {
	// Create new image in the storage
	Thumbnail(context.Context, *ThumbnailRequest) (*Response, error)
}

// UnimplementedThumbnailServiceServer can be embedded to have forward compatible implementations.
type UnimplementedThumbnailServiceServer struct {
}

func (*UnimplementedThumbnailServiceServer) Thumbnail(ctx context.Context, req *ThumbnailRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Thumbnail not implemented")
}

func RegisterThumbnailServiceServer(s *grpc.Server, srv ThumbnailServiceServer) {
	s.RegisterService(&_ThumbnailService_serviceDesc, srv)
}

func _ThumbnailService_Thumbnail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThumbnailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThumbnailServiceServer).Thumbnail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transform.ThumbnailService/Thumbnail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThumbnailServiceServer).Thumbnail(ctx, req.(*ThumbnailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ThumbnailService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transform.ThumbnailService",
	HandlerType: (*ThumbnailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Thumbnail",
			Handler:    _ThumbnailService_Thumbnail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transform.proto",
}