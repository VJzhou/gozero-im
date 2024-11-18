package register

import (
	"context"
	"errors"

	"im/application/applet/internal/svc"
	"im/application/applet/internal/types"
	"im/application/email/rpc/email"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendCodeLogic) SendCode(req *types.SendCodeRequest) (resp *types.SendCodeResponse, err error) {
	// validate email
	_, err = l.svcCtx.EmailRPC.SendEmail(l.ctx, &email.SendEmailRequest{
		Email: req.Email,
	})
	if err != nil {
		return nil, errors.New("email 发送失败")
	}
	resp = &types.SendCodeResponse{
		Code: "123456",
	}
	return
}
