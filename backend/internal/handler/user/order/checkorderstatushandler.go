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

// 前台支付结果轮询(付款跳回后检查是否已到账)
func CheckOrderStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserOrderStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order.NewCheckOrderStatusLogic(r.Context(), svcCtx)
		resp, err := l.CheckOrderStatus(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
