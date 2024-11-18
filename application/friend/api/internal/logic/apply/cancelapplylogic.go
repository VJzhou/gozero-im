package apply

import (
	"context"

	"im/application/friend/api/internal/svc"
	"im/application/friend/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelApplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelApplyLogic {
	return &CancelApplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelApplyLogic) CancelApply(req *types.CancelApplyRequest) (resp *types.CancelApplyResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
