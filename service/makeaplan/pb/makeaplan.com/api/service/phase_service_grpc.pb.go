// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// PhaseQueryClient is the client API for PhaseQuery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PhaseQueryClient interface {
	// 查询计划的所有阶段
	QueryPlanPhaseSummary(ctx context.Context, in *QueryPlanPhaseSummaryRequest, opts ...grpc.CallOption) (*QueryPlanPhaseSummaryResponse, error)
	// 查询阶段
	QueryPhaseSummary(ctx context.Context, in *QueryPhaseSummaryRequest, opts ...grpc.CallOption) (*QueryPhaseSummaryResponse, error)
}

type phaseQueryClient struct {
	cc grpc.ClientConnInterface
}

func NewPhaseQueryClient(cc grpc.ClientConnInterface) PhaseQueryClient {
	return &phaseQueryClient{cc}
}

var phaseQueryQueryPlanPhaseSummaryStreamDesc = &grpc.StreamDesc{
	StreamName: "QueryPlanPhaseSummary",
}

func (c *phaseQueryClient) QueryPlanPhaseSummary(ctx context.Context, in *QueryPlanPhaseSummaryRequest, opts ...grpc.CallOption) (*QueryPlanPhaseSummaryResponse, error) {
	out := new(QueryPlanPhaseSummaryResponse)
	err := c.cc.Invoke(ctx, "/service.PhaseQuery/QueryPlanPhaseSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var phaseQueryQueryPhaseSummaryStreamDesc = &grpc.StreamDesc{
	StreamName: "QueryPhaseSummary",
}

func (c *phaseQueryClient) QueryPhaseSummary(ctx context.Context, in *QueryPhaseSummaryRequest, opts ...grpc.CallOption) (*QueryPhaseSummaryResponse, error) {
	out := new(QueryPhaseSummaryResponse)
	err := c.cc.Invoke(ctx, "/service.PhaseQuery/QueryPhaseSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PhaseQueryService is the service API for PhaseQuery service.
// Fields should be assigned to their respective handler implementations only before
// RegisterPhaseQueryService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type PhaseQueryService struct {
	// 查询计划的所有阶段
	QueryPlanPhaseSummary func(context.Context, *QueryPlanPhaseSummaryRequest) (*QueryPlanPhaseSummaryResponse, error)
	// 查询阶段
	QueryPhaseSummary func(context.Context, *QueryPhaseSummaryRequest) (*QueryPhaseSummaryResponse, error)
}

func (s *PhaseQueryService) queryPlanPhaseSummary(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPlanPhaseSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.QueryPlanPhaseSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/service.PhaseQuery/QueryPlanPhaseSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.QueryPlanPhaseSummary(ctx, req.(*QueryPlanPhaseSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *PhaseQueryService) queryPhaseSummary(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPhaseSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.QueryPhaseSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/service.PhaseQuery/QueryPhaseSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.QueryPhaseSummary(ctx, req.(*QueryPhaseSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterPhaseQueryService registers a service implementation with a gRPC server.
func RegisterPhaseQueryService(s grpc.ServiceRegistrar, srv *PhaseQueryService) {
	srvCopy := *srv
	if srvCopy.QueryPlanPhaseSummary == nil {
		srvCopy.QueryPlanPhaseSummary = func(context.Context, *QueryPlanPhaseSummaryRequest) (*QueryPlanPhaseSummaryResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method QueryPlanPhaseSummary not implemented")
		}
	}
	if srvCopy.QueryPhaseSummary == nil {
		srvCopy.QueryPhaseSummary = func(context.Context, *QueryPhaseSummaryRequest) (*QueryPhaseSummaryResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method QueryPhaseSummary not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "service.PhaseQuery",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "QueryPlanPhaseSummary",
				Handler:    srvCopy.queryPlanPhaseSummary,
			},
			{
				MethodName: "QueryPhaseSummary",
				Handler:    srvCopy.queryPhaseSummary,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "phase_service.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewPhaseQueryService creates a new PhaseQueryService containing the
// implemented methods of the PhaseQuery service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewPhaseQueryService(s interface{}) *PhaseQueryService {
	ns := &PhaseQueryService{}
	if h, ok := s.(interface {
		QueryPlanPhaseSummary(context.Context, *QueryPlanPhaseSummaryRequest) (*QueryPlanPhaseSummaryResponse, error)
	}); ok {
		ns.QueryPlanPhaseSummary = h.QueryPlanPhaseSummary
	}
	if h, ok := s.(interface {
		QueryPhaseSummary(context.Context, *QueryPhaseSummaryRequest) (*QueryPhaseSummaryResponse, error)
	}); ok {
		ns.QueryPhaseSummary = h.QueryPhaseSummary
	}
	return ns
}

// UnstablePhaseQueryService is the service API for PhaseQuery service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstablePhaseQueryService interface {
	// 查询计划的所有阶段
	QueryPlanPhaseSummary(context.Context, *QueryPlanPhaseSummaryRequest) (*QueryPlanPhaseSummaryResponse, error)
	// 查询阶段
	QueryPhaseSummary(context.Context, *QueryPhaseSummaryRequest) (*QueryPhaseSummaryResponse, error)
}

// PhaseModifyClient is the client API for PhaseModify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PhaseModifyClient interface {
	// 创建阶段
	CreatePhase(ctx context.Context, in *CreatePhaseRequest, opts ...grpc.CallOption) (*CreatePhaseResponse, error)
	// 删除阶段
	DeletePhase(ctx context.Context, in *DeletePhaseRequest, opts ...grpc.CallOption) (*DeletePhaseResponse, error)
	// 更新阶段
	UpdatePhase(ctx context.Context, in *UpdatePhaseRequest, opts ...grpc.CallOption) (*UpdatePhaseResponse, error)
}

type phaseModifyClient struct {
	cc grpc.ClientConnInterface
}

func NewPhaseModifyClient(cc grpc.ClientConnInterface) PhaseModifyClient {
	return &phaseModifyClient{cc}
}

var phaseModifyCreatePhaseStreamDesc = &grpc.StreamDesc{
	StreamName: "CreatePhase",
}

func (c *phaseModifyClient) CreatePhase(ctx context.Context, in *CreatePhaseRequest, opts ...grpc.CallOption) (*CreatePhaseResponse, error) {
	out := new(CreatePhaseResponse)
	err := c.cc.Invoke(ctx, "/service.PhaseModify/CreatePhase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var phaseModifyDeletePhaseStreamDesc = &grpc.StreamDesc{
	StreamName: "DeletePhase",
}

func (c *phaseModifyClient) DeletePhase(ctx context.Context, in *DeletePhaseRequest, opts ...grpc.CallOption) (*DeletePhaseResponse, error) {
	out := new(DeletePhaseResponse)
	err := c.cc.Invoke(ctx, "/service.PhaseModify/DeletePhase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var phaseModifyUpdatePhaseStreamDesc = &grpc.StreamDesc{
	StreamName: "UpdatePhase",
}

func (c *phaseModifyClient) UpdatePhase(ctx context.Context, in *UpdatePhaseRequest, opts ...grpc.CallOption) (*UpdatePhaseResponse, error) {
	out := new(UpdatePhaseResponse)
	err := c.cc.Invoke(ctx, "/service.PhaseModify/UpdatePhase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PhaseModifyService is the service API for PhaseModify service.
// Fields should be assigned to their respective handler implementations only before
// RegisterPhaseModifyService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type PhaseModifyService struct {
	// 创建阶段
	CreatePhase func(context.Context, *CreatePhaseRequest) (*CreatePhaseResponse, error)
	// 删除阶段
	DeletePhase func(context.Context, *DeletePhaseRequest) (*DeletePhaseResponse, error)
	// 更新阶段
	UpdatePhase func(context.Context, *UpdatePhaseRequest) (*UpdatePhaseResponse, error)
}

func (s *PhaseModifyService) createPhase(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePhaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.CreatePhase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/service.PhaseModify/CreatePhase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.CreatePhase(ctx, req.(*CreatePhaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *PhaseModifyService) deletePhase(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePhaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.DeletePhase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/service.PhaseModify/DeletePhase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeletePhase(ctx, req.(*DeletePhaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *PhaseModifyService) updatePhase(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePhaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.UpdatePhase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/service.PhaseModify/UpdatePhase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.UpdatePhase(ctx, req.(*UpdatePhaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterPhaseModifyService registers a service implementation with a gRPC server.
func RegisterPhaseModifyService(s grpc.ServiceRegistrar, srv *PhaseModifyService) {
	srvCopy := *srv
	if srvCopy.CreatePhase == nil {
		srvCopy.CreatePhase = func(context.Context, *CreatePhaseRequest) (*CreatePhaseResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method CreatePhase not implemented")
		}
	}
	if srvCopy.DeletePhase == nil {
		srvCopy.DeletePhase = func(context.Context, *DeletePhaseRequest) (*DeletePhaseResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method DeletePhase not implemented")
		}
	}
	if srvCopy.UpdatePhase == nil {
		srvCopy.UpdatePhase = func(context.Context, *UpdatePhaseRequest) (*UpdatePhaseResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method UpdatePhase not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "service.PhaseModify",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "CreatePhase",
				Handler:    srvCopy.createPhase,
			},
			{
				MethodName: "DeletePhase",
				Handler:    srvCopy.deletePhase,
			},
			{
				MethodName: "UpdatePhase",
				Handler:    srvCopy.updatePhase,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "phase_service.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewPhaseModifyService creates a new PhaseModifyService containing the
// implemented methods of the PhaseModify service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewPhaseModifyService(s interface{}) *PhaseModifyService {
	ns := &PhaseModifyService{}
	if h, ok := s.(interface {
		CreatePhase(context.Context, *CreatePhaseRequest) (*CreatePhaseResponse, error)
	}); ok {
		ns.CreatePhase = h.CreatePhase
	}
	if h, ok := s.(interface {
		DeletePhase(context.Context, *DeletePhaseRequest) (*DeletePhaseResponse, error)
	}); ok {
		ns.DeletePhase = h.DeletePhase
	}
	if h, ok := s.(interface {
		UpdatePhase(context.Context, *UpdatePhaseRequest) (*UpdatePhaseResponse, error)
	}); ok {
		ns.UpdatePhase = h.UpdatePhase
	}
	return ns
}

// UnstablePhaseModifyService is the service API for PhaseModify service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstablePhaseModifyService interface {
	// 创建阶段
	CreatePhase(context.Context, *CreatePhaseRequest) (*CreatePhaseResponse, error)
	// 删除阶段
	DeletePhase(context.Context, *DeletePhaseRequest) (*DeletePhaseResponse, error)
	// 更新阶段
	UpdatePhase(context.Context, *UpdatePhaseRequest) (*UpdatePhaseResponse, error)
}
