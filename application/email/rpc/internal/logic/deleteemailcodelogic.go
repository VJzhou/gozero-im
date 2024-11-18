package logic

import (
	"context"

	"im/application/email/rpc/internal/svc"
	"im/application/email/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteEmailCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEmailCodeLogic {
	return &DeleteEmailCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteEmailCodeLogic) DeleteEmailCode(in *service.DeleteEmailCodeRequest) (*service.DeleteEmailCodeResponse, error) {
	_, _ = l.svcCtx.BizRedis.DelCtx(l.ctx, parseEmailKey(in.Email))
	return &service.DeleteEmailCodeResponse{}, nil
}
