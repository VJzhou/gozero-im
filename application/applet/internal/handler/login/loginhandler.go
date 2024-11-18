package login

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"im/application/applet/internal/logic/login"
	"im/application/applet/internal/svc"
	"im/application/applet/internal/types"
	"im/pkg/response"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, err)
			return
		}

		l := login.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Response(w, resp, err)
	}
}
