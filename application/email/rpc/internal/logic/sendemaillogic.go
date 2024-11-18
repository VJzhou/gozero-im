package logic

import (
	"context"
	"fmt"

	"im/application/email/rpc/internal/svc"
	"im/application/email/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

const emailKey = "biz#email#%s" // TODO 添加到配置文件？
const emailCodeExpireTime = 600

type SendEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailLogic {
	return &SendEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendEmailLogic) SendEmail(in *service.SendEmailRequest) (*service.SendEmailResponse, error) {

	code := "123456"
	err := l.svcCtx.BizRedis.SetexCtx(l.ctx, parseEmailKey(in.Email), code, emailCodeExpireTime)
	if err != nil {
		return nil, err
	}

	return &service.SendEmailResponse{}, nil
}

func parseEmailKey(email string) string {
	return fmt.Sprintf(emailKey, email)
}
