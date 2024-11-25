// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: user.proto

package server

import (
	"context"

	"im/application/user/rpc/internal/logic"
	"im/application/user/rpc/internal/svc"
	"im/application/user/rpc/service"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	service.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Register(ctx context.Context, in *service.RegisterRequest) (*service.RegisterResponse, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) FindByEmail(ctx context.Context, in *service.FindByEmailRequest) (*service.FindByEmailResponse, error) {
	l := logic.NewFindByEmailLogic(ctx, s.svcCtx)
	return l.FindByEmail(in)
}

func (s *UserServer) FindById(ctx context.Context, in *service.FindByIdRequest) (*service.FindByIdResponse, error) {
	l := logic.NewFindByIdLogic(ctx, s.svcCtx)
	return l.FindById(in)
}

func (s *UserServer) UpdateUserInfo(ctx context.Context, in *service.UpdateUserInfoResquest) (*service.UpdateUserInfoResponse, error) {
	l := logic.NewUpdateUserInfoLogic(ctx, s.svcCtx)
	return l.UpdateUserInfo(in)
}

func (s *UserServer) ModifyEmail(ctx context.Context, in *service.ModifyEmailRequest) (*service.ModifyEmailResponse, error) {
	l := logic.NewModifyEmailLogic(ctx, s.svcCtx)
	return l.ModifyEmail(in)
}

func (s *UserServer) ModifyMobile(ctx context.Context, in *service.ModifyMobileRequest) (*service.ModifyMobileResponse, error) {
	l := logic.NewModifyMobileLogic(ctx, s.svcCtx)
	return l.ModifyMobile(in)
}

func (s *UserServer) ModifyPassword(ctx context.Context, in *service.ModifyPasswordRequest) (*service.ModifyPasswordResponse, error) {
	l := logic.NewModifyPasswordLogic(ctx, s.svcCtx)
	return l.ModifyPassword(in)
}
