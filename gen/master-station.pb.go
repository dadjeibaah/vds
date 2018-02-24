// Code generated by protoc-gen-go. DO NOT EDIT.
// source: master-station.proto

/*
Package master_station is a generated protocol buffer package.

It is generated from these files:
	master-station.proto

It has these top-level messages:
	LineRequest
	LineInfo
	TrainPing
	Empty
*/
package master_station

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

type LineRequest struct {
	Line      string `protobuf:"bytes,1,opt,name=line" json:"line,omitempty"`
	TrainType string `protobuf:"bytes,2,opt,name=trainType" json:"trainType,omitempty"`
}

func (m *LineRequest) Reset()                    { *m = LineRequest{} }
func (m *LineRequest) String() string            { return proto.CompactTextString(m) }
func (*LineRequest) ProtoMessage()               {}
func (*LineRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LineRequest) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

func (m *LineRequest) GetTrainType() string {
	if m != nil {
		return m.TrainType
	}
	return ""
}

type LineInfo struct {
	TrainType string `protobuf:"bytes,1,opt,name=trainType" json:"trainType,omitempty"`
}

func (m *LineInfo) Reset()                    { *m = LineInfo{} }
func (m *LineInfo) String() string            { return proto.CompactTextString(m) }
func (*LineInfo) ProtoMessage()               {}
func (*LineInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LineInfo) GetTrainType() string {
	if m != nil {
		return m.TrainType
	}
	return ""
}

type TrainPing struct {
	TrainId           string  `protobuf:"bytes,1,opt,name=trainId" json:"trainId,omitempty"`
	Route             string  `protobuf:"bytes,2,opt,name=route" json:"route,omitempty"`
	Status            string  `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	Speed             float64 `protobuf:"fixed64,4,opt,name=speed" json:"speed,omitempty"`
	NextStop          float64 `protobuf:"fixed64,5,opt,name=nextStop" json:"nextStop,omitempty"`
	MinutesToNextStop int32   `protobuf:"varint,6,opt,name=minutesToNextStop" json:"minutesToNextStop,omitempty"`
	Line              string  `protobuf:"bytes,7,opt,name=line" json:"line,omitempty"`
	Direction         string  `protobuf:"bytes,8,opt,name=direction" json:"direction,omitempty"`
	NextStation       string  `protobuf:"bytes,9,opt,name=nextStation" json:"nextStation,omitempty"`
	Timestamp         int64   `protobuf:"varint,10,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *TrainPing) Reset()                    { *m = TrainPing{} }
func (m *TrainPing) String() string            { return proto.CompactTextString(m) }
func (*TrainPing) ProtoMessage()               {}
func (*TrainPing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TrainPing) GetTrainId() string {
	if m != nil {
		return m.TrainId
	}
	return ""
}

func (m *TrainPing) GetRoute() string {
	if m != nil {
		return m.Route
	}
	return ""
}

func (m *TrainPing) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TrainPing) GetSpeed() float64 {
	if m != nil {
		return m.Speed
	}
	return 0
}

func (m *TrainPing) GetNextStop() float64 {
	if m != nil {
		return m.NextStop
	}
	return 0
}

func (m *TrainPing) GetMinutesToNextStop() int32 {
	if m != nil {
		return m.MinutesToNextStop
	}
	return 0
}

func (m *TrainPing) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

func (m *TrainPing) GetDirection() string {
	if m != nil {
		return m.Direction
	}
	return ""
}

func (m *TrainPing) GetNextStation() string {
	if m != nil {
		return m.NextStation
	}
	return ""
}

func (m *TrainPing) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*LineRequest)(nil), "LineRequest")
	proto.RegisterType((*LineInfo)(nil), "LineInfo")
	proto.RegisterType((*TrainPing)(nil), "TrainPing")
	proto.RegisterType((*Empty)(nil), "Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MasterService service

type MasterServiceClient interface {
	ReceiveTrainPing(ctx context.Context, opts ...grpc.CallOption) (MasterService_ReceiveTrainPingClient, error)
}

type masterServiceClient struct {
	cc *grpc.ClientConn
}

func NewMasterServiceClient(cc *grpc.ClientConn) MasterServiceClient {
	return &masterServiceClient{cc}
}

func (c *masterServiceClient) ReceiveTrainPing(ctx context.Context, opts ...grpc.CallOption) (MasterService_ReceiveTrainPingClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_MasterService_serviceDesc.Streams[0], c.cc, "/MasterService/ReceiveTrainPing", opts...)
	if err != nil {
		return nil, err
	}
	x := &masterServiceReceiveTrainPingClient{stream}
	return x, nil
}

type MasterService_ReceiveTrainPingClient interface {
	Send(*TrainPing) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type masterServiceReceiveTrainPingClient struct {
	grpc.ClientStream
}

func (x *masterServiceReceiveTrainPingClient) Send(m *TrainPing) error {
	return x.ClientStream.SendMsg(m)
}

func (x *masterServiceReceiveTrainPingClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for MasterService service

type MasterServiceServer interface {
	ReceiveTrainPing(MasterService_ReceiveTrainPingServer) error
}

func RegisterMasterServiceServer(s *grpc.Server, srv MasterServiceServer) {
	s.RegisterService(&_MasterService_serviceDesc, srv)
}

func _MasterService_ReceiveTrainPing_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MasterServiceServer).ReceiveTrainPing(&masterServiceReceiveTrainPingServer{stream})
}

type MasterService_ReceiveTrainPingServer interface {
	SendAndClose(*Empty) error
	Recv() (*TrainPing, error)
	grpc.ServerStream
}

type masterServiceReceiveTrainPingServer struct {
	grpc.ServerStream
}

func (x *masterServiceReceiveTrainPingServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *masterServiceReceiveTrainPingServer) Recv() (*TrainPing, error) {
	m := new(TrainPing)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MasterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "MasterService",
	HandlerType: (*MasterServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReceiveTrainPing",
			Handler:       _MasterService_ReceiveTrainPing_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "master-station.proto",
}

// Client API for CarService service

type CarServiceClient interface {
	SendTrainPing(ctx context.Context, opts ...grpc.CallOption) (CarService_SendTrainPingClient, error)
}

type carServiceClient struct {
	cc *grpc.ClientConn
}

func NewCarServiceClient(cc *grpc.ClientConn) CarServiceClient {
	return &carServiceClient{cc}
}

func (c *carServiceClient) SendTrainPing(ctx context.Context, opts ...grpc.CallOption) (CarService_SendTrainPingClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_CarService_serviceDesc.Streams[0], c.cc, "/CarService/SendTrainPing", opts...)
	if err != nil {
		return nil, err
	}
	x := &carServiceSendTrainPingClient{stream}
	return x, nil
}

type CarService_SendTrainPingClient interface {
	Send(*TrainPing) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type carServiceSendTrainPingClient struct {
	grpc.ClientStream
}

func (x *carServiceSendTrainPingClient) Send(m *TrainPing) error {
	return x.ClientStream.SendMsg(m)
}

func (x *carServiceSendTrainPingClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for CarService service

type CarServiceServer interface {
	SendTrainPing(CarService_SendTrainPingServer) error
}

func RegisterCarServiceServer(s *grpc.Server, srv CarServiceServer) {
	s.RegisterService(&_CarService_serviceDesc, srv)
}

func _CarService_SendTrainPing_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CarServiceServer).SendTrainPing(&carServiceSendTrainPingServer{stream})
}

type CarService_SendTrainPingServer interface {
	SendAndClose(*Empty) error
	Recv() (*TrainPing, error)
	grpc.ServerStream
}

type carServiceSendTrainPingServer struct {
	grpc.ServerStream
}

func (x *carServiceSendTrainPingServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *carServiceSendTrainPingServer) Recv() (*TrainPing, error) {
	m := new(TrainPing)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CarService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CarService",
	HandlerType: (*CarServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendTrainPing",
			Handler:       _CarService_SendTrainPing_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "master-station.proto",
}

func init() { proto.RegisterFile("master-station.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcf, 0x4b, 0xfb, 0x30,
	0x14, 0x27, 0xdb, 0xda, 0xad, 0x6f, 0x0c, 0xbe, 0xdf, 0x30, 0x24, 0x0c, 0x0f, 0x65, 0x20, 0xe4,
	0xa0, 0x3b, 0x6c, 0x27, 0x4f, 0x1e, 0xc4, 0xc3, 0x40, 0x45, 0xba, 0xfd, 0x03, 0x75, 0x7d, 0x4a,
	0xc0, 0x26, 0x31, 0x79, 0x1d, 0xee, 0x0f, 0xf2, 0xff, 0x94, 0xa6, 0x5b, 0xe7, 0xf4, 0xe0, 0xed,
	0x7d, 0x7e, 0xf5, 0x35, 0x9f, 0x04, 0xc6, 0x65, 0xee, 0x09, 0xdd, 0x95, 0xa7, 0x9c, 0x94, 0xd1,
	0x33, 0xeb, 0x0c, 0x99, 0xe9, 0x0d, 0x0c, 0xef, 0x95, 0xc6, 0x0c, 0xdf, 0x2b, 0xf4, 0xc4, 0x39,
	0xf4, 0xde, 0x94, 0x46, 0xc1, 0x52, 0x26, 0x93, 0x2c, 0xcc, 0xfc, 0x1c, 0x12, 0x72, 0xb9, 0xd2,
	0xeb, 0x9d, 0x45, 0xd1, 0x09, 0xc2, 0x91, 0x98, 0x4a, 0x18, 0xd4, 0x1f, 0x58, 0xea, 0x17, 0x73,
	0xea, 0x64, 0x3f, 0x9d, 0x9f, 0x1d, 0x48, 0xd6, 0x35, 0x7a, 0x52, 0xfa, 0x95, 0x0b, 0xe8, 0x07,
	0x69, 0x59, 0xec, 0x9d, 0x07, 0xc8, 0xc7, 0x10, 0x39, 0x53, 0xd1, 0x61, 0x57, 0x03, 0xf8, 0x19,
	0xc4, 0xf5, 0x9f, 0x57, 0x5e, 0x74, 0x03, 0xbd, 0x47, 0xb5, 0xdb, 0x5b, 0xc4, 0x42, 0xf4, 0x52,
	0x26, 0x59, 0xd6, 0x00, 0x3e, 0x81, 0x81, 0xc6, 0x0f, 0x5a, 0x91, 0xb1, 0x22, 0x0a, 0x42, 0x8b,
	0xf9, 0x25, 0xfc, 0x2f, 0x95, 0xae, 0x08, 0xfd, 0xda, 0x3c, 0x1e, 0x4c, 0x71, 0xca, 0x64, 0x94,
	0xfd, 0x16, 0xda, 0x46, 0xfa, 0xa7, 0x8d, 0x14, 0xca, 0xe1, 0xa6, 0xee, 0x51, 0x0c, 0x9a, 0x73,
	0xb6, 0x04, 0x4f, 0x61, 0xd8, 0xec, 0x0a, 0x3d, 0x8b, 0x24, 0xe8, 0xdf, 0xa9, 0xd0, 0x93, 0x2a,
	0xd1, 0x53, 0x5e, 0x5a, 0x01, 0x29, 0x93, 0xdd, 0xec, 0x48, 0x4c, 0xfb, 0x10, 0xdd, 0x95, 0x96,
	0x76, 0xf3, 0x6b, 0x18, 0x3d, 0x84, 0x3b, 0x5b, 0xa1, 0xdb, 0xaa, 0x0d, 0x72, 0x09, 0xff, 0x32,
	0xdc, 0xa0, 0xda, 0xe2, 0xb1, 0x47, 0x98, 0xb5, 0xf3, 0x24, 0x9e, 0x85, 0xa0, 0x64, 0xf3, 0x05,
	0xc0, 0x6d, 0xde, 0xe6, 0x2e, 0x60, 0xb4, 0x42, 0x5d, 0xfc, 0x11, 0x7a, 0x8e, 0xc3, 0x93, 0x58,
	0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x0e, 0x05, 0x93, 0x2a, 0x02, 0x00, 0x00,
}