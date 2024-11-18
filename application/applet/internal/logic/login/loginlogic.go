package login

import (
	"context"
	"errors"

	"im/application/applet/internal/svc"
	"im/application/applet/internal/types"
	"im/application/user/rpc/user"
	"im/pkg/jwt"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// validate param
	existsUser, err := l.svcCtx.UserRPC.FindByEmail(l.ctx, &user.FindByEmailRequest{
		Email: req.Email,
	})
	if err != nil {
		logx.Errorf("FindByEmail error: %v", err)
		return nil, err
	}

	if existsUser == nil || existsUser.UserId == 0 {
		return nil, errors.New("账号密码错误")
	}

	// TODO chech password
	token, err := jwt.BuildTokens(jwt.TokenOptions{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		Fields: map[string]interface{}{
			"user_id": existsUser.UserId,
		},
	})
	if err != nil {
		return nil, err
	}
	resp = &types.LoginResponse{
		Token: types.Token{
			AccessToken:  token.AccessToken,
			AccessExpire: token.AccessExpire,
		},
	}
	return
}
