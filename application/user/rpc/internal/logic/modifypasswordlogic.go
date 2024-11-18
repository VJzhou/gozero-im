package logic

import (
	"context"
	"errors"

	"im/application/user/rpc/internal/svc"
	"im/application/user/rpc/service"
	"im/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyPasswordLogic {
	return &ModifyPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyPasswordLogic) ModifyPassword(in *service.ModifyPasswordRequest) (*service.ModifyPasswordResponse, error) {
	if len(in.Password) == 0 {
		return nil, errors.New("miss password")
	}

	user, err := l.svcCtx.UserModel.FindById(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	salt, err := util.GenerateSalt()
	if err != nil {
		return nil, err
	}

	passwordHash, err := util.HashPassword(in.Password, salt)
	if err != nil {
		return nil, err
	}

	user.Password = passwordHash
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		return nil, err
	}

	return &service.ModifyPasswordResponse{}, nil
}
