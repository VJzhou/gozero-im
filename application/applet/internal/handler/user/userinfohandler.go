package user

import (
	"net/http"

	"im/application/applet/internal/logic/user"
	"im/application/applet/internal/svc"
	"im/pkg/response"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		response.Response(w, resp, err)
	}
}
