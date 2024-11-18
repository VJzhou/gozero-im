package logic

import (
	"context"
	"errors"

	"im/application/email/rpc/internal/svc"
	"im/application/email/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckEmailLogic {
	return &CheckEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckEmailLogic) CheckEmail(in *service.CheckEmailCodeRequest) (*service.CheckEmailCodeResponse, error) {

	code, err := l.svcCtx.BizRedis.GetCtx(l.ctx, parseEmailKey(in.Email))
	if err != nil || code != in.Code {
		return nil, errors.New("验证码错误")
	}

	return &service.CheckEmailCodeResponse{}, nil
}
