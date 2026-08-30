// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package profile

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goship/backend/internal/logic/user/profile"
	"goship/backend/internal/svc"
	"goship/backend/internal/types"
)

// AI 流式生成演示 (SSE)
func AIGenerateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AIGenerateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := profile.NewAIGenerateLogic(r.Context(), svcCtx)
		err := l.AIGenerate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
