// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package profile

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goship/backend/internal/logic/user/profile"
	"goship/backend/internal/svc"
)

// 获取当前登录用户资料与点数
func GetProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := profile.NewGetProfileLogic(r.Context(), svcCtx)
		resp, err := l.GetProfile()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
