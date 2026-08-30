// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package order

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goship/backend/internal/logic/user/order"
	"goship/backend/internal/svc"
	"goship/backend/internal/types"
)

// 【核心】前台用户下单并获取支付收银台链接
func CreateOrderCheckoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrderCheckoutReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order.NewCreateOrderCheckoutLogic(r.Context(), svcCtx)
		resp, err := l.CreateOrderCheckout(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
