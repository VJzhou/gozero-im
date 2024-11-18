package logic

import (
	"context"

	"im/application/friend/rpc/internal/svc"
	"im/application/friend/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewApplyFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyFriendListLogic {
	return &ApplyFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ApplyFriendListLogic) ApplyFriendList(in *service.ApplyFriendListRequest) (*service.ApplyFriendListResponse, error) {

	return &service.ApplyFriendListResponse{}, nil
}
