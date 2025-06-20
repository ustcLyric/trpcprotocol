// Code generated by trpc-go/trpc-cmdline v1.0.9. DO NOT EDIT.
// source: polar_bear_user.proto

package polar_bear_user

import (
	"context"
	"errors"
	"fmt"

	_ "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/codec"
	_ "trpc.group/trpc-go/trpc-go/http"
	"trpc.group/trpc-go/trpc-go/server"
)

// START ======================================= Server Service Definition ======================================= START

// PolarBearUserServiceService defines service.
type PolarBearUserServiceService interface {
	// UserRegister 用户注册
	UserRegister(ctx context.Context, req *UserRegisterRequest) (*UserRegisterResponse, error)
	// UserLogin 用户登陆
	UserLogin(ctx context.Context, req *UserLoginRequest) (*UserLoginResponse, error)
	// GetUserInfo 获取用户信息
	GetUserInfo(ctx context.Context, req *GetUserInfoReq) (*GetUserInfoResp, error)
	// GetPermissions 获取权限菜单信息
	GetPermissions(ctx context.Context, req *GetPermissionsReq) (*GetPermissionsResp, error)
	// GetRoles 获取角色信息
	GetRoles(ctx context.Context, req *GetRolesReq) (*GetRolesResp, error)
}

func PolarBearUserServiceService_UserRegister_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &UserRegisterRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(PolarBearUserServiceService).UserRegister(ctx, reqbody.(*UserRegisterRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func PolarBearUserServiceService_UserLogin_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &UserLoginRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(PolarBearUserServiceService).UserLogin(ctx, reqbody.(*UserLoginRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func PolarBearUserServiceService_GetUserInfo_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &GetUserInfoReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(PolarBearUserServiceService).GetUserInfo(ctx, reqbody.(*GetUserInfoReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func PolarBearUserServiceService_GetPermissions_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &GetPermissionsReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(PolarBearUserServiceService).GetPermissions(ctx, reqbody.(*GetPermissionsReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func PolarBearUserServiceService_GetRoles_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &GetRolesReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(PolarBearUserServiceService).GetRoles(ctx, reqbody.(*GetRolesReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// PolarBearUserServiceServer_ServiceDesc descriptor for server.RegisterService.
var PolarBearUserServiceServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "trpc.polarBear.user.PolarBearUserService",
	HandlerType: ((*PolarBearUserServiceService)(nil)),
	Methods: []server.Method{
		{
			Name: "/trpc.polarBear.user.PolarBearUserService/UserRegister",
			Func: PolarBearUserServiceService_UserRegister_Handler,
		},
		{
			Name: "/trpc.polarBear.user.PolarBearUserService/UserLogin",
			Func: PolarBearUserServiceService_UserLogin_Handler,
		},
		{
			Name: "/trpc.polarBear.user.PolarBearUserService/GetUserInfo",
			Func: PolarBearUserServiceService_GetUserInfo_Handler,
		},
		{
			Name: "/trpc.polarBear.user.PolarBearUserService/GetPermissions",
			Func: PolarBearUserServiceService_GetPermissions_Handler,
		},
		{
			Name: "/trpc.polarBear.user.PolarBearUserService/GetRoles",
			Func: PolarBearUserServiceService_GetRoles_Handler,
		},
	},
}

// RegisterPolarBearUserServiceService registers service.
func RegisterPolarBearUserServiceService(s server.Service, svr PolarBearUserServiceService) {
	if err := s.Register(&PolarBearUserServiceServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("PolarBearUserService register error:%v", err))
	}
}

// START --------------------------------- Default Unimplemented Server Service --------------------------------- START

type UnimplementedPolarBearUserService struct{}

// UserRegister 用户注册
func (s *UnimplementedPolarBearUserService) UserRegister(ctx context.Context, req *UserRegisterRequest) (*UserRegisterResponse, error) {
	return nil, errors.New("rpc UserRegister of service PolarBearUserService is not implemented")
}

// UserLogin 用户登陆
func (s *UnimplementedPolarBearUserService) UserLogin(ctx context.Context, req *UserLoginRequest) (*UserLoginResponse, error) {
	return nil, errors.New("rpc UserLogin of service PolarBearUserService is not implemented")
}

// GetUserInfo 获取用户信息
func (s *UnimplementedPolarBearUserService) GetUserInfo(ctx context.Context, req *GetUserInfoReq) (*GetUserInfoResp, error) {
	return nil, errors.New("rpc GetUserInfo of service PolarBearUserService is not implemented")
}

// GetPermissions 获取权限菜单信息
func (s *UnimplementedPolarBearUserService) GetPermissions(ctx context.Context, req *GetPermissionsReq) (*GetPermissionsResp, error) {
	return nil, errors.New("rpc GetPermissions of service PolarBearUserService is not implemented")
}

// GetRoles 获取角色信息
func (s *UnimplementedPolarBearUserService) GetRoles(ctx context.Context, req *GetRolesReq) (*GetRolesResp, error) {
	return nil, errors.New("rpc GetRoles of service PolarBearUserService is not implemented")
}

// END --------------------------------- Default Unimplemented Server Service --------------------------------- END

// END ======================================= Server Service Definition ======================================= END

// START ======================================= Client Service Definition ======================================= START

// PolarBearUserServiceClientProxy defines service client proxy
type PolarBearUserServiceClientProxy interface {
	// UserRegister 用户注册
	UserRegister(ctx context.Context, req *UserRegisterRequest, opts ...client.Option) (rsp *UserRegisterResponse, err error)
	// UserLogin 用户登陆
	UserLogin(ctx context.Context, req *UserLoginRequest, opts ...client.Option) (rsp *UserLoginResponse, err error)
	// GetUserInfo 获取用户信息
	GetUserInfo(ctx context.Context, req *GetUserInfoReq, opts ...client.Option) (rsp *GetUserInfoResp, err error)
	// GetPermissions 获取权限菜单信息
	GetPermissions(ctx context.Context, req *GetPermissionsReq, opts ...client.Option) (rsp *GetPermissionsResp, err error)
	// GetRoles 获取角色信息
	GetRoles(ctx context.Context, req *GetRolesReq, opts ...client.Option) (rsp *GetRolesResp, err error)
}

type PolarBearUserServiceClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewPolarBearUserServiceClientProxy = func(opts ...client.Option) PolarBearUserServiceClientProxy {
	return &PolarBearUserServiceClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *PolarBearUserServiceClientProxyImpl) UserRegister(ctx context.Context, req *UserRegisterRequest, opts ...client.Option) (*UserRegisterResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.polarBear.user.PolarBearUserService/UserRegister")
	msg.WithCalleeServiceName(PolarBearUserServiceServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("polarBear")
	msg.WithCalleeServer("user")
	msg.WithCalleeService("PolarBearUserService")
	msg.WithCalleeMethod("UserRegister")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &UserRegisterResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *PolarBearUserServiceClientProxyImpl) UserLogin(ctx context.Context, req *UserLoginRequest, opts ...client.Option) (*UserLoginResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.polarBear.user.PolarBearUserService/UserLogin")
	msg.WithCalleeServiceName(PolarBearUserServiceServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("polarBear")
	msg.WithCalleeServer("user")
	msg.WithCalleeService("PolarBearUserService")
	msg.WithCalleeMethod("UserLogin")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &UserLoginResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *PolarBearUserServiceClientProxyImpl) GetUserInfo(ctx context.Context, req *GetUserInfoReq, opts ...client.Option) (*GetUserInfoResp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.polarBear.user.PolarBearUserService/GetUserInfo")
	msg.WithCalleeServiceName(PolarBearUserServiceServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("polarBear")
	msg.WithCalleeServer("user")
	msg.WithCalleeService("PolarBearUserService")
	msg.WithCalleeMethod("GetUserInfo")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &GetUserInfoResp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *PolarBearUserServiceClientProxyImpl) GetPermissions(ctx context.Context, req *GetPermissionsReq, opts ...client.Option) (*GetPermissionsResp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.polarBear.user.PolarBearUserService/GetPermissions")
	msg.WithCalleeServiceName(PolarBearUserServiceServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("polarBear")
	msg.WithCalleeServer("user")
	msg.WithCalleeService("PolarBearUserService")
	msg.WithCalleeMethod("GetPermissions")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &GetPermissionsResp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *PolarBearUserServiceClientProxyImpl) GetRoles(ctx context.Context, req *GetRolesReq, opts ...client.Option) (*GetRolesResp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.polarBear.user.PolarBearUserService/GetRoles")
	msg.WithCalleeServiceName(PolarBearUserServiceServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("polarBear")
	msg.WithCalleeServer("user")
	msg.WithCalleeService("PolarBearUserService")
	msg.WithCalleeMethod("GetRoles")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &GetRolesResp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

// END ======================================= Client Service Definition ======================================= END
