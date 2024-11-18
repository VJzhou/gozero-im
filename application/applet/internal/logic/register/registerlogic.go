package register

import (
	"context"

	"im/application/applet/internal/svc"
	"im/application/applet/internal/types"
	"im/application/email/rpc/email"
	"im/application/user/rpc/user"
	"im/pkg/jwt"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// TODO validate param

	// check email code
	_, err = l.svcCtx.EmailRPC.CheckEmail(l.ctx, &email.CheckEmailCodeRequest{
		Email: req.Email,
		Code:  req.Code,
	})
	if err != nil {
		return nil, err
	}

	existsUser, err := l.svcCtx.UserRPC.FindByEmail(l.ctx, &user.FindByEmailRequest{
		Email: req.Email,
	})
	if err != nil {
		logx.Errorf("FindByEmail error: %v", err)
		return nil, err
	}

	if existsUser != nil && existsUser.UserId > 0 {
		logx.Errorf("Email exists Email: %s", req.Email)
		return nil, err
	}

	res, err := l.svcCtx.UserRPC.Register(l.ctx, &user.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	if err != nil {
		logx.Errorf("Register err: %v", err)
		return nil, err
	}

	// generate token
	token, err := jwt.BuildTokens(jwt.TokenOptions{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		Fields: map[string]interface{}{
			"user_id": res.Id,
		},
	})
	if err != nil {
		logx.Errorf("BuildTokens error: %v", err)
		return nil, err
	}

	l.svcCtx.EmailRPC.DeleteEmailCode(l.ctx, &email.DeleteEmailCodeRequest{
		Email: req.Email,
	})

	resp = &types.RegisterResponse{
		UserId: res.Id,
		Token: types.Token{
			AccessToken:  token.AccessToken,
			AccessExpire: token.AccessExpire,
		},
	}

	return
}
