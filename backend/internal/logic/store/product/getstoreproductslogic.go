// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package product

import (
	"context"

	"goship/backend/internal/svc"
	"goship/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStoreProductsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取前台已上架商品列表 (SSR展示)
func NewGetStoreProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStoreProductsLogic {
	return &GetStoreProductsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStoreProductsLogic) GetStoreProducts() (resp *types.StoreProductListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
