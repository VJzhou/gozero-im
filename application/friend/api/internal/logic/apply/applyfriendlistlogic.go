package apply

import (
	"context"

	"im/application/friend/api/internal/svc"
	"im/application/friend/api/internal/types"
	"im/application/friend/rpc/friend"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplyFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyFriendListLogic {
	return &ApplyFriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplyFriendListLogic) ApplyFriendList() (resp *types.ApplyFriendListResponse, err error) {

	applyList, err := l.svcCtx.FriendRPC.ApplyFriendList(l.ctx, &friend.ApplyFriendListRequest{})
	if err != nil {
		return nil, err
	}

	if applyList == nil || len(applyList.Applys) == 0 {
		return resp, nil
	}

	// write to cache
	list := make([]types.ApplyFriendItem, 0, len(applyList.Applys))
	for _, item := range applyList.Applys {
		list = append(list, types.ApplyFriendItem{
			Id:        item.Id,
			UserId:    item.UserId,
			FriendId:  item.FriendId,
			Remark:    item.Remark,
			Nickname:  item.Nickname,
			Avater:    item.Avatar,
			CreatedAt: item.CreatedAt,
		})
	}

	return &types.ApplyFriendListResponse{
		Itmes: list,
	}, nil
}
