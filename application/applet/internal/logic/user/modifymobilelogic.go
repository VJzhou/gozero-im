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

type ModifyMobileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyMobileLogic {
	return &ModifyMobileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyMobileLogic) ModifyMobile(req *types.ModifyMobileRequest) (resp *types.ModifyMobileResponse, err error) {
	if len(req.Password) == 0 {
		return nil, code.PasswordEmpty
	}
	if len(req.Mobile) == 0 {
		return nil, code.MobileEmpty
	}

	if len(req.Code) == 0 {
		return nil, code.CodeEmpty
	}

	userId, err := util.UserId(l.ctx)
	if err != nil {
		return nil, err
	}

	// TODO check password

	_, err = l.svcCtx.UserRPC.ModifyMobile(l.ctx, &user.ModifyMobileRequest{
		UserId: userId,
		Mobile: req.Mobile,
	})
	if err != nil {
		l.Errorf("ModifyMobile api err: %v", err)
		return nil, code.ChangeMobileFailed
	}

	return
}
