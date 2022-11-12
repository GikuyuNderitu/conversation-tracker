// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: notes.proto

package api

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

// NotesServiceClient is the client API for NotesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotesServiceClient interface {
	GetNotes(ctx context.Context, in *GetNotesRequest, opts ...grpc.CallOption) (*GetNotesResponse, error)
	GetConversation(ctx context.Context, in *GetConversationRequest, opts ...grpc.CallOption) (*GetConversationResponse, error)
	ListConversations(ctx context.Context, in *ListConversationsRequest, opts ...grpc.CallOption) (*ListConversationsResponse, error)
	CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*CreateNoteResponse, error)
}

type notesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotesServiceClient(cc grpc.ClientConnInterface) NotesServiceClient {
	return &notesServiceClient{cc}
}

func (c *notesServiceClient) GetNotes(ctx context.Context, in *GetNotesRequest, opts ...grpc.CallOption) (*GetNotesResponse, error) {
	out := new(GetNotesResponse)
	err := c.cc.Invoke(ctx, "/notes.v1.api.NotesService/GetNotes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesServiceClient) GetConversation(ctx context.Context, in *GetConversationRequest, opts ...grpc.CallOption) (*GetConversationResponse, error) {
	out := new(GetConversationResponse)
	err := c.cc.Invoke(ctx, "/notes.v1.api.NotesService/GetConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesServiceClient) ListConversations(ctx context.Context, in *ListConversationsRequest, opts ...grpc.CallOption) (*ListConversationsResponse, error) {
	out := new(ListConversationsResponse)
	err := c.cc.Invoke(ctx, "/notes.v1.api.NotesService/ListConversations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesServiceClient) CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*CreateNoteResponse, error) {
	out := new(CreateNoteResponse)
	err := c.cc.Invoke(ctx, "/notes.v1.api.NotesService/CreateNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotesServiceServer is the server API for NotesService service.
// All implementations must embed UnimplementedNotesServiceServer
// for forward compatibility
type NotesServiceServer interface {
	GetNotes(context.Context, *GetNotesRequest) (*GetNotesResponse, error)
	GetConversation(context.Context, *GetConversationRequest) (*GetConversationResponse, error)
	ListConversations(context.Context, *ListConversationsRequest) (*ListConversationsResponse, error)
	CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error)
	mustEmbedUnimplementedNotesServiceServer()
}

// UnimplementedNotesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotesServiceServer struct {
}

func (UnimplementedNotesServiceServer) GetNotes(context.Context, *GetNotesRequest) (*GetNotesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotes not implemented")
}
func (UnimplementedNotesServiceServer) GetConversation(context.Context, *GetConversationRequest) (*GetConversationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversation not implemented")
}
func (UnimplementedNotesServiceServer) ListConversations(context.Context, *ListConversationsRequest) (*ListConversationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConversations not implemented")
}
func (UnimplementedNotesServiceServer) CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNote not implemented")
}
func (UnimplementedNotesServiceServer) mustEmbedUnimplementedNotesServiceServer() {}

// UnsafeNotesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotesServiceServer will
// result in compilation errors.
type UnsafeNotesServiceServer interface {
	mustEmbedUnimplementedNotesServiceServer()
}

func RegisterNotesServiceServer(s grpc.ServiceRegistrar, srv NotesServiceServer) {
	s.RegisterService(&NotesService_ServiceDesc, srv)
}

func _NotesService_GetNotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServiceServer).GetNotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.v1.api.NotesService/GetNotes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServiceServer).GetNotes(ctx, req.(*GetNotesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotesService_GetConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServiceServer).GetConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.v1.api.NotesService/GetConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServiceServer).GetConversation(ctx, req.(*GetConversationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotesService_ListConversations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConversationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServiceServer).ListConversations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.v1.api.NotesService/ListConversations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServiceServer).ListConversations(ctx, req.(*ListConversationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotesService_CreateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServiceServer).CreateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.v1.api.NotesService/CreateNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServiceServer).CreateNote(ctx, req.(*CreateNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotesService_ServiceDesc is the grpc.ServiceDesc for NotesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notes.v1.api.NotesService",
	HandlerType: (*NotesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNotes",
			Handler:    _NotesService_GetNotes_Handler,
		},
		{
			MethodName: "GetConversation",
			Handler:    _NotesService_GetConversation_Handler,
		},
		{
			MethodName: "ListConversations",
			Handler:    _NotesService_ListConversations_Handler,
		},
		{
			MethodName: "CreateNote",
			Handler:    _NotesService_CreateNote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notes.proto",
}
