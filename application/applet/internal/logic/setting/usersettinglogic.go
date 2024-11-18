package setting

import (
	"context"
	"encoding/json"

	"im/application/applet/internal/svc"
	"im/application/applet/internal/types"
	"im/application/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSettingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSettingLogic {
	return &UserSettingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserSettingLogic) UserSetting() (resp *types.UserSettingResponse, err error) {
	userId, err := l.ctx.Value("user_id").(json.Number).Int64()
	if err != nil {
		return nil, err
	}

	user, err := l.svcCtx.UserRPC.FindById(l.ctx, &user.FindByIdRequest{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.UserSettingResponse{
		UserInfo: types.UserInfo{
			UserId:   user.UserId,
			Email:    user.Email,
			Mobile:   user.Mobile,
			Avatar:   user.Avatar,
			Motto:    user.Motto,
			Birthday: user.Birthday,
			Gender:   int(user.Gender),
		},
		UserSetting: types.UserSetting{},
	}

	return
}
