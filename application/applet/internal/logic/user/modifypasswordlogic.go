package user

import (
	"context"

	"im/application/applet/code"
	"im/application/applet/internal/svc"
	"im/application/applet/internal/types"
	"im/application/user/rpc/user"
	"im/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyPasswordLogic {
	return &ModifyPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyPasswordLogic) ModifyPassword(req *types.ModifyPasswordRequest) (resp *types.ModifyPasswordResponse, err error) {
	// TODO 引入validator？
	if len(req.Password) == 0 {
		return nil, code.PasswordEmpty
	}

	if len(req.NewPassword) == 0 {
		return nil, code.NewPasswordEmpty
	}

	userId, err := util.UserId(l.ctx)
	if err != nil {
		return nil, err
	}

	// TODO check password

	_, err = l.svcCtx.UserRPC.ModifyPassword(l.ctx, &user.ModifyPasswordRequest{
		UserId:   userId,
		Password: req.NewPassword,
	})
	if err != nil {
		l.Errorf("ModifyPassword api err: %v", err)
		return nil, code.ChangePasswordFailed
	}

	return resp, nil
}
