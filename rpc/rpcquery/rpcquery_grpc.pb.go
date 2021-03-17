// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rpcquery

import (
	context "context"

	acm "github.com/klye-dev/hsc-main/acm"
	names "github.com/klye-dev/hsc-main/execution/names"
	rpc "github.com/klye-dev/hsc-main/rpc"
	payload "github.com/klye-dev/hsc-main/txs/payload"
	types "github.com/tendermint/tendermint/proto/tendermint/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	Status(ctx context.Context, in *StatusParam, opts ...grpc.CallOption) (*rpc.ResultStatus, error)
	GetAccount(ctx context.Context, in *GetAccountParam, opts ...grpc.CallOption) (*acm.Account, error)
	GetMetadata(ctx context.Context, in *GetMetadataParam, opts ...grpc.CallOption) (*MetadataResult, error)
	GetStorage(ctx context.Context, in *GetStorageParam, opts ...grpc.CallOption) (*StorageValue, error)
	ListAccounts(ctx context.Context, in *ListAccountsParam, opts ...grpc.CallOption) (Query_ListAccountsClient, error)
	GetName(ctx context.Context, in *GetNameParam, opts ...grpc.CallOption) (*names.Entry, error)
	ListNames(ctx context.Context, in *ListNamesParam, opts ...grpc.CallOption) (Query_ListNamesClient, error)
	// GetNetworkRegistry returns for each validator address, the list of their identified node at the current state
	GetNetworkRegistry(ctx context.Context, in *GetNetworkRegistryParam, opts ...grpc.CallOption) (*NetworkRegistry, error)
	GetValidatorSet(ctx context.Context, in *GetValidatorSetParam, opts ...grpc.CallOption) (*ValidatorSet, error)
	GetValidatorSetHistory(ctx context.Context, in *GetValidatorSetHistoryParam, opts ...grpc.CallOption) (*ValidatorSetHistory, error)
	GetProposal(ctx context.Context, in *GetProposalParam, opts ...grpc.CallOption) (*payload.Ballot, error)
	ListProposals(ctx context.Context, in *ListProposalsParam, opts ...grpc.CallOption) (Query_ListProposalsClient, error)
	GetStats(ctx context.Context, in *GetStatsParam, opts ...grpc.CallOption) (*Stats, error)
	GetBlockHeader(ctx context.Context, in *GetBlockParam, opts ...grpc.CallOption) (*types.Header, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Status(ctx context.Context, in *StatusParam, opts ...grpc.CallOption) (*rpc.ResultStatus, error) {
	out := new(rpc.ResultStatus)
	err := c.cc.Invoke(ctx, "/rpcquery.Query/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetAccount(ctx context.Context, in *GetAccountParam, opts ...grpc.CallOption) (*acm.Account, error) {
	out := new(acm.Account)
	err := c.cc.Invoke(ctx, "/rpcquery.Query/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetMetadata(ctx context.Context, in *GetMetadataParam, opts ...grpc.CallOption) (*MetadataResult, error) {
	out := new(MetadataResult)
	err := c.cc.Invoke(ctx, "/rpcquery.Query/GetMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetStorage(ctx context.Context, in *GetStorageParam, opts ...grpc.CallOption) (*StorageValue, error) {
	out := new(StorageValue)
	err := c.cc.Invoke(ctx, "/rpcquery.Query/GetStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListAccounts(ctx context.Context, in *ListAccountsParam, opts ...grpc.CallOption) (Query_ListAccountsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Query_ServiceDesc.Streams[0], "/rpcquery.Query/ListAccounts", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryListAccountsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Query_ListAccountsClient interface {
	Recv() (*acm.Account, error)
	grpc.ClientStream
}

type queryListAccountsClient struct {
	grpc.ClientStream
}

func (x *queryListAccountsClient) Recv() (*acm.Account, error) {
	m := new(acm.Account)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *queryClient) GetName(ctx context.Context, in *GetNameParam, opts ...grpc.CallOption) (*names.Entry, error) {
	out := new(names.Entry)
	err := c.cc.Invoke(ctx, "/rpcquery.Query/GetName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListNames(ctx context.Context, in *ListNamesParam, opts ...grpc.CallOption) (Query_ListNamesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Query_ServiceDesc.Streams[1], "/rpcquery.Query/ListNames", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryListNamesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Query_ListNamesClient interface {
	Recv() (*names.Entry, error)
	grpc.ClientStream
}

type queryListNamesClient struct {
	grpc.ClientStream
}

func (x *queryListNamesClient) Recv() (*names.Entry, error) {
	m := new(names.Entry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *queryClient) GetNetworkRegistry(ctx context.Context, in *GetNetworkRegistryParam, opts ...grpc.CallOption) (*NetworkRegistry, error) {
	out := new(NetworkRegistry)
	err := c.cc.Invoke(ctx, "/rpcquery.Query/GetNetworkRegistry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetValidatorSet(ctx context.Context, in *GetValidatorSetParam, opts ...grpc.CallOption) (*ValidatorSet, error) {
	out := new(ValidatorSet)
	err := c.cc.Invoke(ctx, "/rpcquery.Query/GetValidatorSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetValidatorSetHistory(ctx context.Context, in *GetValidatorSetHistoryParam, opts ...grpc.CallOption) (*ValidatorSetHistory, error) {
	out := new(ValidatorSetHistory)
	err := c.cc.Invoke(ctx, "/rpcquery.Query/GetValidatorSetHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetProposal(ctx context.Context, in *GetProposalParam, opts ...grpc.CallOption) (*payload.Ballot, error) {
	out := new(payload.Ballot)
	err := c.cc.Invoke(ctx, "/rpcquery.Query/GetProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListProposals(ctx context.Context, in *ListProposalsParam, opts ...grpc.CallOption) (Query_ListProposalsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Query_ServiceDesc.Streams[2], "/rpcquery.Query/ListProposals", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryListProposalsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Query_ListProposalsClient interface {
	Recv() (*ProposalResult, error)
	grpc.ClientStream
}

type queryListProposalsClient struct {
	grpc.ClientStream
}

func (x *queryListProposalsClient) Recv() (*ProposalResult, error) {
	m := new(ProposalResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *queryClient) GetStats(ctx context.Context, in *GetStatsParam, opts ...grpc.CallOption) (*Stats, error) {
	out := new(Stats)
	err := c.cc.Invoke(ctx, "/rpcquery.Query/GetStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetBlockHeader(ctx context.Context, in *GetBlockParam, opts ...grpc.CallOption) (*types.Header, error) {
	out := new(types.Header)
	err := c.cc.Invoke(ctx, "/rpcquery.Query/GetBlockHeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	Status(context.Context, *StatusParam) (*rpc.ResultStatus, error)
	GetAccount(context.Context, *GetAccountParam) (*acm.Account, error)
	GetMetadata(context.Context, *GetMetadataParam) (*MetadataResult, error)
	GetStorage(context.Context, *GetStorageParam) (*StorageValue, error)
	ListAccounts(*ListAccountsParam, Query_ListAccountsServer) error
	GetName(context.Context, *GetNameParam) (*names.Entry, error)
	ListNames(*ListNamesParam, Query_ListNamesServer) error
	// GetNetworkRegistry returns for each validator address, the list of their identified node at the current state
	GetNetworkRegistry(context.Context, *GetNetworkRegistryParam) (*NetworkRegistry, error)
	GetValidatorSet(context.Context, *GetValidatorSetParam) (*ValidatorSet, error)
	GetValidatorSetHistory(context.Context, *GetValidatorSetHistoryParam) (*ValidatorSetHistory, error)
	GetProposal(context.Context, *GetProposalParam) (*payload.Ballot, error)
	ListProposals(*ListProposalsParam, Query_ListProposalsServer) error
	GetStats(context.Context, *GetStatsParam) (*Stats, error)
	GetBlockHeader(context.Context, *GetBlockParam) (*types.Header, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Status(context.Context, *StatusParam) (*rpc.ResultStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedQueryServer) GetAccount(context.Context, *GetAccountParam) (*acm.Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedQueryServer) GetMetadata(context.Context, *GetMetadataParam) (*MetadataResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetadata not implemented")
}
func (UnimplementedQueryServer) GetStorage(context.Context, *GetStorageParam) (*StorageValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStorage not implemented")
}
func (UnimplementedQueryServer) ListAccounts(*ListAccountsParam, Query_ListAccountsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAccounts not implemented")
}
func (UnimplementedQueryServer) GetName(context.Context, *GetNameParam) (*names.Entry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetName not implemented")
}
func (UnimplementedQueryServer) ListNames(*ListNamesParam, Query_ListNamesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListNames not implemented")
}
func (UnimplementedQueryServer) GetNetworkRegistry(context.Context, *GetNetworkRegistryParam) (*NetworkRegistry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkRegistry not implemented")
}
func (UnimplementedQueryServer) GetValidatorSet(context.Context, *GetValidatorSetParam) (*ValidatorSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValidatorSet not implemented")
}
func (UnimplementedQueryServer) GetValidatorSetHistory(context.Context, *GetValidatorSetHistoryParam) (*ValidatorSetHistory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValidatorSetHistory not implemented")
}
func (UnimplementedQueryServer) GetProposal(context.Context, *GetProposalParam) (*payload.Ballot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProposal not implemented")
}
func (UnimplementedQueryServer) ListProposals(*ListProposalsParam, Query_ListProposalsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListProposals not implemented")
}
func (UnimplementedQueryServer) GetStats(context.Context, *GetStatsParam) (*Stats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}
func (UnimplementedQueryServer) GetBlockHeader(context.Context, *GetBlockParam) (*types.Header, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHeader not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcquery.Query/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Status(ctx, req.(*StatusParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcquery.Query/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetAccount(ctx, req.(*GetAccountParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetadataParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcquery.Query/GetMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetMetadata(ctx, req.(*GetMetadataParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStorageParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcquery.Query/GetStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetStorage(ctx, req.(*GetStorageParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListAccounts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListAccountsParam)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServer).ListAccounts(m, &queryListAccountsServer{stream})
}

type Query_ListAccountsServer interface {
	Send(*acm.Account) error
	grpc.ServerStream
}

type queryListAccountsServer struct {
	grpc.ServerStream
}

func (x *queryListAccountsServer) Send(m *acm.Account) error {
	return x.ServerStream.SendMsg(m)
}

func _Query_GetName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNameParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcquery.Query/GetName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetName(ctx, req.(*GetNameParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListNames_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListNamesParam)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServer).ListNames(m, &queryListNamesServer{stream})
}

type Query_ListNamesServer interface {
	Send(*names.Entry) error
	grpc.ServerStream
}

type queryListNamesServer struct {
	grpc.ServerStream
}

func (x *queryListNamesServer) Send(m *names.Entry) error {
	return x.ServerStream.SendMsg(m)
}

func _Query_GetNetworkRegistry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkRegistryParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetNetworkRegistry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcquery.Query/GetNetworkRegistry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetNetworkRegistry(ctx, req.(*GetNetworkRegistryParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetValidatorSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValidatorSetParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetValidatorSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcquery.Query/GetValidatorSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetValidatorSet(ctx, req.(*GetValidatorSetParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetValidatorSetHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValidatorSetHistoryParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetValidatorSetHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcquery.Query/GetValidatorSetHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetValidatorSetHistory(ctx, req.(*GetValidatorSetHistoryParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProposalParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcquery.Query/GetProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetProposal(ctx, req.(*GetProposalParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListProposals_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListProposalsParam)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServer).ListProposals(m, &queryListProposalsServer{stream})
}

type Query_ListProposalsServer interface {
	Send(*ProposalResult) error
	grpc.ServerStream
}

type queryListProposalsServer struct {
	grpc.ServerStream
}

func (x *queryListProposalsServer) Send(m *ProposalResult) error {
	return x.ServerStream.SendMsg(m)
}

func _Query_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatsParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcquery.Query/GetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetStats(ctx, req.(*GetStatsParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetBlockHeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetBlockHeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcquery.Query/GetBlockHeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetBlockHeader(ctx, req.(*GetBlockParam))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpcquery.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Query_Status_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _Query_GetAccount_Handler,
		},
		{
			MethodName: "GetMetadata",
			Handler:    _Query_GetMetadata_Handler,
		},
		{
			MethodName: "GetStorage",
			Handler:    _Query_GetStorage_Handler,
		},
		{
			MethodName: "GetName",
			Handler:    _Query_GetName_Handler,
		},
		{
			MethodName: "GetNetworkRegistry",
			Handler:    _Query_GetNetworkRegistry_Handler,
		},
		{
			MethodName: "GetValidatorSet",
			Handler:    _Query_GetValidatorSet_Handler,
		},
		{
			MethodName: "GetValidatorSetHistory",
			Handler:    _Query_GetValidatorSetHistory_Handler,
		},
		{
			MethodName: "GetProposal",
			Handler:    _Query_GetProposal_Handler,
		},
		{
			MethodName: "GetStats",
			Handler:    _Query_GetStats_Handler,
		},
		{
			MethodName: "GetBlockHeader",
			Handler:    _Query_GetBlockHeader_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListAccounts",
			Handler:       _Query_ListAccounts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListNames",
			Handler:       _Query_ListNames_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListProposals",
			Handler:       _Query_ListProposals_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpcquery.proto",
}
