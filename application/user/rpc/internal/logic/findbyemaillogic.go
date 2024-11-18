package logic

import (
	"context"

	"im/application/user/rpc/code"
	"im/application/user/rpc/internal/svc"
	"im/application/user/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByEmailLogic {
	return &FindByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByEmailLogic) FindByEmail(in *service.FindByEmailRequest) (*service.FindByEmailResponse, error) {
	if len(in.Email) == 0 {
		return nil, code.EmailEmpty
	}

	user, err := l.svcCtx.UserModel.FindByEmail(l.ctx, in.Email)
	if err != nil {
		l.Errorf("FindByEmail req: %v, err: %v", in.Email, err)
		return nil, err
	}
	if user == nil {
		return &service.FindByEmailResponse{}, code.NotFoundUser
	}

	return &service.FindByEmailResponse{
		UserId:   user.Id,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Mobile:   user.Mobile,
		Nickname: user.Nickname,
		Birthday: user.Birthday,
		Motto:    user.Motto,
		Gender:   int64(user.Gender),
	}, nil
}
