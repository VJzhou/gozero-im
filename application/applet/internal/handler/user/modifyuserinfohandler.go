package user

import (
	"net/http"

	"im/application/applet/internal/logic/user"
	"im/application/applet/internal/svc"
	"im/application/applet/internal/types"
	"im/pkg/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ModifyUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ModifyUserInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, err)
			return
		}

		l := user.NewModifyUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.ModifyUserInfo(&req)
		response.Response(w, resp, err)
	}
}
