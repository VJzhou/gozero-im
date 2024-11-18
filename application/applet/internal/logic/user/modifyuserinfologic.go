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

type ModifyUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyUserInfoLogic {
	return &ModifyUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyUserInfoLogic) ModifyUserInfo(req *types.ModifyUserInfoRequest) (resp *types.ModifyUserInfoResponse, err error) {
	if len(req.Nickname) == 0 {
		return nil, code.NicknameEmpty
	}

	userId, err := util.UserId(l.ctx)
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.UserRPC.UpdateUserInfo(l.ctx, &user.UpdateUserInfoResquest{
		Avatar:   req.Avatar,
		Nickname: req.Nickname,
		Birthday: req.Birthday,
		Motto:    req.Motto,
		Gender:   int32(req.Gender),
		UserId:   userId,
	})
	if err != nil {
		return nil, code.ModifyUserInfoFailed
	}

	return resp, nil
}
