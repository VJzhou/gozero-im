package logic

import (
	"context"

	"im/application/user/rpc/code"
	"im/application/user/rpc/internal/svc"
	"im/application/user/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyEmailLogic {
	return &ModifyEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyEmailLogic) ModifyEmail(in *service.ModifyEmailRequest) (*service.ModifyEmailResponse, error) {
	if len(in.Email) == 0 {
		return nil, code.EmailEmpty
	}
	if in.UserId == 0 {
		return nil, code.IdEmpty
	}

	user, err := l.svcCtx.UserModel.FindById(l.ctx, in.UserId)
	if err != nil {
		l.Errorf("ModifyEmail FindById id: %d, err: %v", in.UserId, err)
		return nil, err
	}

	if user.Email != in.Email {
		existEmailUser, err := l.svcCtx.UserModel.FindByEmail(l.ctx, in.Email)
		if err != nil {
			l.Errorf("ModifyEmail FindByEmail email: %d, err: %v", in.Email, err)
			return nil, err
		}

		if (existEmailUser != nil) && (existEmailUser.Id != user.Id) {
			return nil, code.EmailExist
		}
	} // TODO 与旧邮箱一致？

	user.Email = in.Email
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		l.Errorf("ModifyEmail Update email: %v, err: %v", in.Email, err)
		return nil, err
	}

	return &service.ModifyEmailResponse{}, nil
}
