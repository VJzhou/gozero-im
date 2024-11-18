package register

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"im/application/applet/internal/logic/register"
	"im/application/applet/internal/svc"
	"im/application/applet/internal/types"
	"im/pkg/response"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, err)
			return
		}

		l := register.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		response.Response(w, resp, err)
	}
}
