// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goship/backend/internal/logic/admin/user"
	"goship/backend/internal/svc"
)

// 后台获取系统全局配置
func GetSystemConfigsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetSystemConfigsLogic(r.Context(), svcCtx)
		resp, err := l.GetSystemConfigs()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
