// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.19.6
// source: document.proto

package document

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	NewDocument_InsertDocument_FullMethodName = "/document.NewDocument/InsertDocument"
)

// NewDocumentClient is the client API for NewDocument service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NewDocumentClient interface {
	InsertDocument(ctx context.Context, in *DocumentRequest, opts ...grpc.CallOption) (*DocumentResponse, error)
}

type newDocumentClient struct {
	cc grpc.ClientConnInterface
}

func NewNewDocumentClient(cc grpc.ClientConnInterface) NewDocumentClient {
	return &newDocumentClient{cc}
}

func (c *newDocumentClient) InsertDocument(ctx context.Context, in *DocumentRequest, opts ...grpc.CallOption) (*DocumentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DocumentResponse)
	err := c.cc.Invoke(ctx, NewDocument_InsertDocument_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewDocumentServer is the server API for NewDocument service.
// All implementations must embed UnimplementedNewDocumentServer
// for forward compatibility
type NewDocumentServer interface {
	InsertDocument(context.Context, *DocumentRequest) (*DocumentResponse, error)
	mustEmbedUnimplementedNewDocumentServer()
}

// UnimplementedNewDocumentServer must be embedded to have forward compatible implementations.
type UnimplementedNewDocumentServer struct {
}

func (UnimplementedNewDocumentServer) InsertDocument(context.Context, *DocumentRequest) (*DocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertDocument not implemented")
}
func (UnimplementedNewDocumentServer) mustEmbedUnimplementedNewDocumentServer() {}

// UnsafeNewDocumentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewDocumentServer will
// result in compilation errors.
type UnsafeNewDocumentServer interface {
	mustEmbedUnimplementedNewDocumentServer()
}

func RegisterNewDocumentServer(s grpc.ServiceRegistrar, srv NewDocumentServer) {
	s.RegisterService(&NewDocument_ServiceDesc, srv)
}

func _NewDocument_InsertDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewDocumentServer).InsertDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewDocument_InsertDocument_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewDocumentServer).InsertDocument(ctx, req.(*DocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NewDocument_ServiceDesc is the grpc.ServiceDesc for NewDocument service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NewDocument_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "document.NewDocument",
	HandlerType: (*NewDocumentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertDocument",
			Handler:    _NewDocument_InsertDocument_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "document.proto",
}

const (
	RemoveDocument_DeleteDocument_FullMethodName = "/document.RemoveDocument/DeleteDocument"
)

// RemoveDocumentClient is the client API for RemoveDocument service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemoveDocumentClient interface {
	DeleteDocument(ctx context.Context, in *DeleteDocumentRequest, opts ...grpc.CallOption) (*DeleteDocumentResponse, error)
}

type removeDocumentClient struct {
	cc grpc.ClientConnInterface
}

func NewRemoveDocumentClient(cc grpc.ClientConnInterface) RemoveDocumentClient {
	return &removeDocumentClient{cc}
}

func (c *removeDocumentClient) DeleteDocument(ctx context.Context, in *DeleteDocumentRequest, opts ...grpc.CallOption) (*DeleteDocumentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteDocumentResponse)
	err := c.cc.Invoke(ctx, RemoveDocument_DeleteDocument_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoveDocumentServer is the server API for RemoveDocument service.
// All implementations must embed UnimplementedRemoveDocumentServer
// for forward compatibility
type RemoveDocumentServer interface {
	DeleteDocument(context.Context, *DeleteDocumentRequest) (*DeleteDocumentResponse, error)
	mustEmbedUnimplementedRemoveDocumentServer()
}

// UnimplementedRemoveDocumentServer must be embedded to have forward compatible implementations.
type UnimplementedRemoveDocumentServer struct {
}

func (UnimplementedRemoveDocumentServer) DeleteDocument(context.Context, *DeleteDocumentRequest) (*DeleteDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDocument not implemented")
}
func (UnimplementedRemoveDocumentServer) mustEmbedUnimplementedRemoveDocumentServer() {}

// UnsafeRemoveDocumentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemoveDocumentServer will
// result in compilation errors.
type UnsafeRemoveDocumentServer interface {
	mustEmbedUnimplementedRemoveDocumentServer()
}

func RegisterRemoveDocumentServer(s grpc.ServiceRegistrar, srv RemoveDocumentServer) {
	s.RegisterService(&RemoveDocument_ServiceDesc, srv)
}

func _RemoveDocument_DeleteDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoveDocumentServer).DeleteDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoveDocument_DeleteDocument_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoveDocumentServer).DeleteDocument(ctx, req.(*DeleteDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RemoveDocument_ServiceDesc is the grpc.ServiceDesc for RemoveDocument service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RemoveDocument_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "document.RemoveDocument",
	HandlerType: (*RemoveDocumentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteDocument",
			Handler:    _RemoveDocument_DeleteDocument_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "document.proto",
}
