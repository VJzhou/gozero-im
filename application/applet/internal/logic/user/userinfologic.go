package user

import (
	"context"

	"im/application/applet/internal/svc"
	"im/application/applet/internal/types"
	"im/application/user/rpc/user"
	"im/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	userId, err := util.UserId(l.ctx)
	if err != nil {
		return nil, err
	}

	user, err := l.svcCtx.UserRPC.FindById(l.ctx, &user.FindByIdRequest{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.UserInfoResponse{
		UserInfo: types.UserInfo{
			UserId:   user.UserId,
			Email:    user.Email,
			Mobile:   user.Mobile,
			Avatar:   user.Avatar,
			Motto:    user.Motto,
			Birthday: user.Birthday,
			Gender:   int(user.Gender),
			Nickname: user.Nickname,
		},
	}
	return
}
