// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package product

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goship/backend/internal/logic/admin/product"
	"goship/backend/internal/svc"
	"goship/backend/internal/types"
)

// 后台删除商品
func DeleteProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdminDeleteProductReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := product.NewDeleteProductLogic(r.Context(), svcCtx)
		resp, err := l.DeleteProduct(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
