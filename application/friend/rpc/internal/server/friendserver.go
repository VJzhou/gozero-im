// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: friend.proto

package server

import (
	"context"

	"im/application/friend/rpc/internal/logic"
	"im/application/friend/rpc/internal/svc"
	"im/application/friend/rpc/service"
)

type FriendServer struct {
	svcCtx *svc.ServiceContext
	service.UnimplementedFriendServer
}

func NewFriendServer(svcCtx *svc.ServiceContext) *FriendServer {
	return &FriendServer{
		svcCtx: svcCtx,
	}
}

func (s *FriendServer) ApplyFriend(ctx context.Context, in *service.ApplyFriendRequest) (*service.ApplyFriendResponse, error) {
	l := logic.NewApplyFriendLogic(ctx, s.svcCtx)
	return l.ApplyFriend(in)
}

func (s *FriendServer) ApplyFriendList(ctx context.Context, in *service.ApplyFriendListRequest) (*service.ApplyFriendListResponse, error) {
	l := logic.NewApplyFriendListLogic(ctx, s.svcCtx)
	return l.ApplyFriendList(in)
}

func (s *FriendServer) ApplyPass(ctx context.Context, in *service.ApplyHandlerRequest) (*service.ApplyHandlerResponse, error) {
	l := logic.NewApplyPassLogic(ctx, s.svcCtx)
	return l.ApplyPass(in)
}

func (s *FriendServer) ApplyReject(ctx context.Context, in *service.ApplyHandlerRequest) (*service.ApplyHandlerResponse, error) {
	l := logic.NewApplyRejectLogic(ctx, s.svcCtx)
	return l.ApplyReject(in)
}

func (s *FriendServer) FriendList(ctx context.Context, in *service.FriendListRequest) (*service.FriendListResponse, error) {
	l := logic.NewFriendListLogic(ctx, s.svcCtx)
	return l.FriendList(in)
}

func (s *FriendServer) RemoveFriend(ctx context.Context, in *service.RemoveFriendRequest) (*service.RemoveFriendResponse, error) {
	l := logic.NewRemoveFriendLogic(ctx, s.svcCtx)
	return l.RemoveFriend(in)
}