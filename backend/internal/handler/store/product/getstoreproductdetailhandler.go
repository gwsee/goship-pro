// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package product

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goship/backend/internal/logic/store/product"
	"goship/backend/internal/svc"
	"goship/backend/internal/types"
)

// 获取单个商品详情页
func GetStoreProductDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StoreProductDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := product.NewGetStoreProductDetailLogic(r.Context(), svcCtx)
		resp, err := l.GetStoreProductDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
