// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package product

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goship/backend/internal/logic/store/product"
	"goship/backend/internal/svc"
)

// 获取前台已上架商品列表 (SSR展示)
func GetStoreProductsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := product.NewGetStoreProductsLogic(r.Context(), svcCtx)
		resp, err := l.GetStoreProducts()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
