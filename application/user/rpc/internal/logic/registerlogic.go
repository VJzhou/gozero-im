package logic

import (
	"context"
	"time"

	"im/application/user/rpc/code"
	"im/application/user/rpc/internal/model"
	"im/application/user/rpc/internal/svc"
	"im/application/user/rpc/service"
	"im/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *service.RegisterRequest) (*service.RegisterResponse, error) {
	if len(in.Email) == 0 {
		return nil, code.EmailEmpty
	}
	if len(in.Password) == 0 {
		return nil, code.PasswordEmpty
	}
	if len(in.Nickname) == 0 {
		return nil, code.NicknameEmpty
	}

	// insert
	salt, err := util.GenerateSalt()
	if err != nil {
		l.Errorf("register generate salt err: %v", err)
		return nil, code.GenerateSaltFailed
	}

	passwordHash, err := util.HashPassword(in.Password, salt)
	if err != nil {
		l.Errorf("register generate password hash err: %v", err)
		return nil, err
	}

	user := &model.User{
		Email:     in.Email,
		Password:  passwordHash,
		Salt:      salt,
		Nickname:  in.Nickname,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		l.Errorf("register failed req: %v err: %v", user, err)
		return nil, err
	}
	return &service.RegisterResponse{
		Id: user.Id,
	}, nil
}
