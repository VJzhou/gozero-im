package apply

import (
	"net/http"

	"im/application/friend/api/internal/logic/apply"
	"im/application/friend/api/internal/svc"
	"im/pkg/response"
)

func ApplyFriendListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := apply.NewApplyFriendListLogic(r.Context(), svcCtx)
		resp, err := l.ApplyFriendList()
		response.Response(w, resp, err)
	}
}
