package logic

import (
	"context"
	"errors"

	"im/application/friend/rpc/internal/svc"
	"im/application/friend/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type RemoveFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveFriendLogic {
	return &RemoveFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveFriendLogic) RemoveFriend(in *service.RemoveFriendRequest) (*service.RemoveFriendResponse, error) {
	// todo: add your logic here and delete this line
	if in.UserId == 0 {
		return nil, errors.New("")
	}
	if in.FriendId == 0 {
		return nil, errors.New("")
	}

	l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		contact, err := l.svcCtx.FriendModel.FindByUidAndFid(l.ctx, in.UserId, in.FriendId)
		if err != nil {
			return err
		}
		err = l.svcCtx.FriendModel.DeleteById(l.ctx, contact.Id)
		if err != nil {
			return err
		}
		return nil
	})
	return &service.RemoveFriendResponse{}, nil
}
