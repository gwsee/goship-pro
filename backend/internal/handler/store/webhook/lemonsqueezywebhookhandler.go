// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package webhook

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goship/backend/internal/logic/store/webhook"
	"goship/backend/internal/svc"
)

// Lemon Squeezy 支付成功回调通知
func LemonSqueezyWebhookHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := webhook.NewLemonSqueezyWebhookLogic(r.Context(), svcCtx)
		err := l.LemonSqueezyWebhook()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
