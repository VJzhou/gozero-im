package logic

import (
	"context"

	"im/application/friend/rpc/internal/svc"
	"im/application/friend/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendListLogic) FriendList(in *service.FriendListRequest) (*service.FriendListResponse, error) {
	// todo: add your logic here and delete this line

	return &service.FriendListResponse{}, nil
}
