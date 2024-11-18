package logic

import (
	"context"
	"errors"

	"im/application/friend/rpc/internal/model"
	"im/application/friend/rpc/internal/svc"
	"im/application/friend/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyRejectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewApplyRejectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyRejectLogic {
	return &ApplyRejectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ApplyRejectLogic) ApplyReject(in *service.ApplyHandlerRequest) (*service.ApplyHandlerResponse, error) {
	if in.Id == 0 {
		return nil, errors.New("")
	}
	if in.UserId == 0 {
		return nil, errors.New("")
	}

	err := l.svcCtx.ApplyFriendModel.UpdateStatus(l.ctx, in.Id, model.StatusReject)
	if err != nil {
		return nil, err
	}

	return &service.ApplyHandlerResponse{}, nil
}
