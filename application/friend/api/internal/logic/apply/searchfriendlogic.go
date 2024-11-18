package apply

import (
	"context"
	"errors"

	"im/application/friend/api/internal/svc"
	"im/application/friend/api/internal/types"
	"im/application/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchFriendLogic {
	return &SearchFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchFriendLogic) SearchFriend(req *types.SearchFriendRequest) (resp *types.SearchFriendResponse, err error) {

	if len(req.Mobile) == 0 {
		return nil, errors.New("")
	}

	existUser, err := l.svcCtx.UserRPC.FindByEmail(l.ctx, &user.FindByEmailRequest{
		Email: req.Mobile,
	})

	if err != nil {
		return nil, err
	}

	if existUser != nil {
		resp = &types.SearchFriendResponse{}
		return resp, nil
	}
	return nil, errors.New("用户不存在！")
}
