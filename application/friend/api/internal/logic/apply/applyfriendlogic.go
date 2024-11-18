package apply

import (
	"context"
	"errors"

	"im/application/friend/api/internal/svc"
	"im/application/friend/api/internal/types"
	"im/application/friend/rpc/friend"
	"im/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplyFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyFriendLogic {
	return &ApplyFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplyFriendLogic) ApplyFriend(req *types.ApplyFriendRequest) (resp *types.ApplyFriendResponse, err error) {
	if req.ApplyId == 0 {
		return nil, errors.New("")
	}

	if len(req.Remark) == 0 {
		return nil, errors.New("")
	}

	userId, err := util.UserId(l.ctx)
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.FriendRPC.ApplyFriend(l.ctx, &friend.ApplyFriendRequest{
		UserId: userId,
		Fid:    req.ApplyId,
		Remark: req.Remark,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
