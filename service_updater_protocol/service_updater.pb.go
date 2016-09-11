// Code generated by protoc-gen-go.
// source: service_updater.proto
// DO NOT EDIT!

/*
Package service_updater is a generated protocol buffer package.

It is generated from these files:
	service_updater.proto

It has these top-level messages:
	UpdateRequest
	UpdateResponse
*/
package service_updater

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

//
// The status code of the service update operation.
type UpdateResponse_Status int32

const (
	UpdateResponse_SUCCESS UpdateResponse_Status = 0
	UpdateResponse_FAILURE UpdateResponse_Status = 1
)

var UpdateResponse_Status_name = map[int32]string{
	0: "SUCCESS",
	1: "FAILURE",
}
var UpdateResponse_Status_value = map[string]int32{
	"SUCCESS": 0,
	"FAILURE": 1,
}

func (x UpdateResponse_Status) String() string {
	return proto.EnumName(UpdateResponse_Status_name, int32(x))
}
func (UpdateResponse_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type UpdateRequest struct {
	//
	// The name of the service within the swarm. This is not necessarily the name of the image. Example: polymer-demo
	ServiceName string `protobuf:"bytes,1,opt,name=serviceName" json:"serviceName,omitempty"`
	//
	// The name of the image with which to update the service. This should be in the typical Docker image form. Example: cookinrelaxin/polymer-demo:18
	ImageName string `protobuf:"bytes,2,opt,name=imageName" json:"imageName,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type UpdateResponse struct {
	Status UpdateResponse_Status `protobuf:"varint,1,opt,name=status,enum=service_updater.UpdateResponse_Status" json:"status,omitempty"`
	//
	// The exact message returned by Docker Engine upon attempting to update the service.
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*UpdateRequest)(nil), "service_updater.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "service_updater.UpdateResponse")
	proto.RegisterEnum("service_updater.UpdateResponse_Status", UpdateResponse_Status_name, UpdateResponse_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for ServiceUpdater service

type ServiceUpdaterClient interface {
	//
	// Accepts an update request for a Docker swarm-mode service.
	UpdateService(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
}

type serviceUpdaterClient struct {
	cc *grpc.ClientConn
}

func NewServiceUpdaterClient(cc *grpc.ClientConn) ServiceUpdaterClient {
	return &serviceUpdaterClient{cc}
}

func (c *serviceUpdaterClient) UpdateService(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := grpc.Invoke(ctx, "/service_updater.ServiceUpdater/UpdateService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ServiceUpdater service

type ServiceUpdaterServer interface {
	//
	// Accepts an update request for a Docker swarm-mode service.
	UpdateService(context.Context, *UpdateRequest) (*UpdateResponse, error)
}

func RegisterServiceUpdaterServer(s *grpc.Server, srv ServiceUpdaterServer) {
	s.RegisterService(&_ServiceUpdater_serviceDesc, srv)
}

func _ServiceUpdater_UpdateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceUpdaterServer).UpdateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_updater.ServiceUpdater/UpdateService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceUpdaterServer).UpdateService(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceUpdater_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service_updater.ServiceUpdater",
	HandlerType: (*ServiceUpdaterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateService",
			Handler:    _ServiceUpdater_UpdateService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("service_updater.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x8d, 0x2f, 0x2d, 0x48, 0x49, 0x2c, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x47, 0x13, 0x56, 0xf2, 0xe7, 0xe2, 0x0d, 0x05, 0x33, 0x83, 0x52, 0x0b, 0x4b,
	0x53, 0x8b, 0x4b, 0x84, 0x14, 0xb8, 0xb8, 0xa1, 0x6a, 0xfc, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0x90, 0x85, 0x84, 0x64, 0xb8, 0x38, 0x33, 0x73, 0x13, 0xd3, 0x21, 0xf2,
	0x4c, 0x60, 0x79, 0x84, 0x80, 0x52, 0x1f, 0x23, 0x17, 0x1f, 0xcc, 0xc4, 0xe2, 0x82, 0xfc, 0xbc,
	0xe2, 0x54, 0x21, 0x3b, 0x2e, 0xb6, 0xe2, 0x92, 0xc4, 0x92, 0xd2, 0x62, 0xb0, 0x69, 0x7c, 0x46,
	0x6a, 0x7a, 0xe8, 0x8e, 0x43, 0xd5, 0xa0, 0x17, 0x0c, 0x56, 0x1d, 0x04, 0xd5, 0x25, 0x24, 0xc1,
	0xc5, 0x9e, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x0e, 0xb3, 0x0e, 0xc6, 0x55, 0x52, 0xe2, 0x62, 0x83,
	0xa8, 0x15, 0xe2, 0xe6, 0x62, 0x0f, 0x0e, 0x75, 0x76, 0x76, 0x0d, 0x0e, 0x16, 0x60, 0x00, 0x71,
	0xdc, 0x1c, 0x3d, 0x7d, 0x42, 0x83, 0x5c, 0x05, 0x18, 0x8d, 0x52, 0xb8, 0xf8, 0x82, 0x21, 0xd6,
	0x41, 0x6c, 0x29, 0x12, 0x0a, 0x82, 0xf9, 0x19, 0x2a, 0x2e, 0x24, 0x87, 0xd3, 0x41, 0xe0, 0x30,
	0x91, 0x92, 0x27, 0xe0, 0x60, 0x25, 0x86, 0x24, 0x36, 0x70, 0xf8, 0x1a, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x48, 0x29, 0xae, 0x21, 0x78, 0x01, 0x00, 0x00,
}
