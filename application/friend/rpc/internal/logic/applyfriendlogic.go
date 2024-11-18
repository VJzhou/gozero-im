package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"im/application/friend/rpc/internal/model"
	"im/application/friend/rpc/internal/svc"
	"im/application/friend/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

const applyPendingPrefix = "biz#friend#applypending#%d#%d"

func parseApplyPending(uid, fid int64) string {
	return fmt.Sprintf(applyPendingPrefix, uid, fid)
}

type ApplyFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewApplyFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyFriendLogic {
	return &ApplyFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ApplyFriendLogic) ApplyFriend(in *service.ApplyFriendRequest) (*service.ApplyFriendResponse, error) {
	if in.UserId == 0 {
		return nil, errors.New("miss user id")
	}
	if in.Fid == 0 {
		return nil, errors.New("miss friend id")
	}

	row, err := l.svcCtx.ApplyFriendModel.FindByUidAndFid(l.ctx, in.UserId, in.Fid)
	if err != nil {
		return nil, err
	}
	if row != nil && (row.Status == model.StatusPending) { // 已经申请
		l.Logger.Errorf("already apply friend, userId: %d, friendId: %d", in.UserId, in.Fid)
		return nil, nil
	}

	err = l.svcCtx.ApplyFriendModel.Insert(l.ctx, &model.ApplyFriend{
		UserId:    in.UserId,
		FriendId:  in.Fid,
		Remark:    in.Remark,
		Status:    model.StatusPending,
		CreatedAt: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &service.ApplyFriendResponse{}, nil
}
