// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: api/proto/anime_radio.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AnimeRadioService_SendAnimeRadioInfo_FullMethodName = "/AnimeRadioService/SendAnimeRadioInfo"
)

// AnimeRadioServiceClient is the client API for AnimeRadioService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnimeRadioServiceClient interface {
	SendAnimeRadioInfo(ctx context.Context, opts ...grpc.CallOption) (AnimeRadioService_SendAnimeRadioInfoClient, error)
}

type animeRadioServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnimeRadioServiceClient(cc grpc.ClientConnInterface) AnimeRadioServiceClient {
	return &animeRadioServiceClient{cc}
}

func (c *animeRadioServiceClient) SendAnimeRadioInfo(ctx context.Context, opts ...grpc.CallOption) (AnimeRadioService_SendAnimeRadioInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &AnimeRadioService_ServiceDesc.Streams[0], AnimeRadioService_SendAnimeRadioInfo_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &animeRadioServiceSendAnimeRadioInfoClient{stream}
	return x, nil
}

type AnimeRadioService_SendAnimeRadioInfoClient interface {
	Send(*YouTubeInfo) error
	CloseAndRecv() (*SlackResponse, error)
	grpc.ClientStream
}

type animeRadioServiceSendAnimeRadioInfoClient struct {
	grpc.ClientStream
}

func (x *animeRadioServiceSendAnimeRadioInfoClient) Send(m *YouTubeInfo) error {
	return x.ClientStream.SendMsg(m)
}

func (x *animeRadioServiceSendAnimeRadioInfoClient) CloseAndRecv() (*SlackResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SlackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AnimeRadioServiceServer is the server API for AnimeRadioService service.
// All implementations must embed UnimplementedAnimeRadioServiceServer
// for forward compatibility
type AnimeRadioServiceServer interface {
	SendAnimeRadioInfo(AnimeRadioService_SendAnimeRadioInfoServer) error
	mustEmbedUnimplementedAnimeRadioServiceServer()
}

// UnimplementedAnimeRadioServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAnimeRadioServiceServer struct {
}

func (UnimplementedAnimeRadioServiceServer) SendAnimeRadioInfo(AnimeRadioService_SendAnimeRadioInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method SendAnimeRadioInfo not implemented")
}
func (UnimplementedAnimeRadioServiceServer) mustEmbedUnimplementedAnimeRadioServiceServer() {}

// UnsafeAnimeRadioServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnimeRadioServiceServer will
// result in compilation errors.
type UnsafeAnimeRadioServiceServer interface {
	mustEmbedUnimplementedAnimeRadioServiceServer()
}

func RegisterAnimeRadioServiceServer(s grpc.ServiceRegistrar, srv AnimeRadioServiceServer) {
	s.RegisterService(&AnimeRadioService_ServiceDesc, srv)
}

func _AnimeRadioService_SendAnimeRadioInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AnimeRadioServiceServer).SendAnimeRadioInfo(&animeRadioServiceSendAnimeRadioInfoServer{stream})
}

type AnimeRadioService_SendAnimeRadioInfoServer interface {
	SendAndClose(*SlackResponse) error
	Recv() (*YouTubeInfo, error)
	grpc.ServerStream
}

type animeRadioServiceSendAnimeRadioInfoServer struct {
	grpc.ServerStream
}

func (x *animeRadioServiceSendAnimeRadioInfoServer) SendAndClose(m *SlackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *animeRadioServiceSendAnimeRadioInfoServer) Recv() (*YouTubeInfo, error) {
	m := new(YouTubeInfo)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AnimeRadioService_ServiceDesc is the grpc.ServiceDesc for AnimeRadioService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnimeRadioService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AnimeRadioService",
	HandlerType: (*AnimeRadioServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendAnimeRadioInfo",
			Handler:       _AnimeRadioService_SendAnimeRadioInfo_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "api/proto/anime_radio.proto",
}
