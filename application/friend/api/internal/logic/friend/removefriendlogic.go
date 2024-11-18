package friend

import (
	"context"
	"errors"

	"im/application/friend/api/internal/svc"
	"im/application/friend/api/internal/types"
	"im/application/friend/rpc/friend"
	"im/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveFriendLogic {
	return &RemoveFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveFriendLogic) RemoveFriend(req *types.RemoveFriendRequest) (resp *types.RemoveFriendResponse, err error) {
	// todo: add your logic here and delete this line
	if req.FriendId == 0 {
		return nil, errors.New("")
	}

	userId, err := util.UserId(l.ctx)
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.FriendRPC.RemoveFriend(l.ctx, &friend.RemoveFriendRequest{
		UserId:   userId,
		FriendId: req.FriendId,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
