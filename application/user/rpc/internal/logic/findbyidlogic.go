package logic

import (
	"context"

	"im/application/user/rpc/code"
	"im/application/user/rpc/internal/svc"
	"im/application/user/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByIdLogic {
	return &FindByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByIdLogic) FindById(in *service.FindByIdRequest) (*service.FindByIdResponse, error) {
	if in.Id == 0 {
		return nil, code.IdEmpty
	}
	user, err := l.svcCtx.UserModel.FindById(l.ctx, in.Id)
	if err != nil {
		l.Errorf("FindById req: %d, err: %v", in.Id, err)
		return nil, err
	}
	return &service.FindByIdResponse{
		UserId:   user.Id,
		Email:    user.Email,
		Mobile:   user.Mobile,
		Motto:    user.Motto,
		Birthday: user.Birthday,
		Gender:   int64(user.Gender),
		Avatar:   user.Avatar,
		Nickname: user.Nickname,
	}, nil
}
