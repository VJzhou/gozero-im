package logic

import (
	"context"

	"im/application/user/rpc/code"
	"im/application/user/rpc/internal/svc"
	"im/application/user/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(in *service.UpdateUserInfoResquest) (*service.UpdateUserInfoResponse, error) {
	// validate param
	if in.UserId == 0 {
		return nil, code.IdEmpty
	}

	user, err := l.svcCtx.UserModel.FindById(l.ctx, in.UserId)
	if err != nil {
		l.Errorf("UpdateUserInfo FindById id: %d, err: %v", in.UserId, err)
		return nil, err
	}

	user.Avatar = in.Avatar
	user.Birthday = in.Birthday
	user.Gender = int(in.Gender)
	user.Motto = in.Motto
	user.Nickname = in.Nickname

	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		l.Errorf("UpdateUserInfo failed req: %v, err: %v", user, err)
		return nil, code.UpdateUserInfoFailed
	}

	return &service.UpdateUserInfoResponse{}, nil
}
