package logic

import (
	"context"

	"im/application/user/rpc/code"
	"im/application/user/rpc/internal/svc"
	"im/application/user/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyMobileLogic {
	return &ModifyMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyMobileLogic) ModifyMobile(in *service.ModifyMobileRequest) (*service.ModifyMobileResponse, error) {
	if len(in.Mobile) == 0 {
		return nil, code.MobileEmpty
	}
	if in.UserId == 0 {
		return nil, code.IdEmpty
	}

	user, err := l.svcCtx.UserModel.FindById(l.ctx, in.UserId)
	if err != nil {
		l.Errorf("ModifyMobile FindById id: %d, err: %v", in.UserId, err)
		return nil, err
	}

	if user.Mobile != in.Mobile {
		existMobileUser, err := l.svcCtx.UserModel.FindByMobile(l.ctx, in.Mobile)
		if err != nil {
			l.Errorf("ModifyMobile FindByMobile email: %d, err: %v", in.Mobile, err)
			return nil, err
		}

		if (existMobileUser != nil) && (existMobileUser.Id != user.Id) {
			return nil, code.MobileExist
		}
	} // TODO 与旧Mobile一致？

	user.Mobile = in.Mobile
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		l.Errorf("ModifyMobile Update mobile: %v, err: %v", in.Mobile, err)
		return nil, err
	}

	return &service.ModifyMobileResponse{}, nil
}
