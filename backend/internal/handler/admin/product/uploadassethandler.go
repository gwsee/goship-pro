// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package product

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goship/backend/internal/logic/admin/product"
	"goship/backend/internal/svc"
)

// 直传 Cloudflare R2 图片
func UploadAssetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := product.NewUploadAssetLogic(r.Context(), svcCtx)
		resp, err := l.UploadAsset()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
