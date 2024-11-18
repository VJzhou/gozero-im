package friend

import (
	"net/http"

	"im/application/friend/api/internal/logic/friend"
	"im/application/friend/api/internal/svc"
	"im/application/friend/api/internal/types"
	"im/pkg/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RemoveFriendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveFriendRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, err)
			return
		}

		l := friend.NewRemoveFriendLogic(r.Context(), svcCtx)
		resp, err := l.RemoveFriend(&req)
		response.Response(w, resp, err)
	}
}
