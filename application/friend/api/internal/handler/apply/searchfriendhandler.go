package apply

import (
	"net/http"

	"im/application/friend/api/internal/logic/apply"
	"im/application/friend/api/internal/svc"
	"im/application/friend/api/internal/types"
	"im/pkg/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SearchFriendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchFriendRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, err)
			return
		}

		l := apply.NewSearchFriendLogic(r.Context(), svcCtx)
		resp, err := l.SearchFriend(&req)
		response.Response(w, resp, err)
	}
}
