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

// 前台用户分页查询我的订单列表
func GetMyOrdersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserOrderListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order.NewGetMyOrdersLogic(r.Context(), svcCtx)
		resp, err := l.GetMyOrders(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
