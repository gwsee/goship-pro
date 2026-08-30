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

// 前台用户查看单个订单详情(含物流与收件地址)
func GetOrderDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserOrderDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order.NewGetOrderDetailLogic(r.Context(), svcCtx)
		resp, err := l.GetOrderDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
