package setting

import (
	"net/http"

	"im/application/applet/internal/logic/setting"
	"im/application/applet/internal/svc"
	"im/pkg/response"
)

func UserSettingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := setting.NewUserSettingLogic(r.Context(), svcCtx)
		resp, err := l.UserSetting()
		response.Response(w, resp, err)
	}
}
