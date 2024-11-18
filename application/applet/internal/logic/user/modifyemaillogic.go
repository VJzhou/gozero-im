package user

import (
	"context"

	"im/application/applet/code"
	"im/application/applet/internal/svc"
	"im/application/applet/internal/types"
	"im/application/email/rpc/email"
	"im/application/user/rpc/user"
	"im/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyEmailLogic {
	return &ModifyEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyEmailLogic) ModifyEmail(req *types.ModifyEmailRequest) (resp *types.ModifyEmailResponse, err error) {
	if len(req.Password) == 0 {
		return nil, code.PasswordEmpty
	}
	if len(req.Email) == 0 {
		return nil, code.EmailEmpty
	}

	if len(req.Code) == 0 {
		return nil, code.CodeEmpty
	}

	userId, err := util.UserId(l.ctx)
	if err != nil {
		return nil, err
	}

	// TODO check password

	_, err = l.svcCtx.EmailRPC.CheckEmail(l.ctx, &email.CheckEmailCodeRequest{
		Email: req.Email,
		Code:  req.Code,
	})
	if err != nil {
		return nil, code.CodeIncorrect
	}

	_, err = l.svcCtx.UserRPC.ModifyEmail(l.ctx, &user.ModifyEmailRequest{
		UserId: userId,
		Email:  req.Email,
	})
	if err != nil {
		l.Errorf("ModifyEmail api err: %v", err)
		return nil, code.ChangeEmailFailed
	}

	return nil, nil
}
