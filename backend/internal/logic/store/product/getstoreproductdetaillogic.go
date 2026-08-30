// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package product

import (
	"context"

	"goship/backend/internal/svc"
	"goship/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStoreProductDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取单个商品详情页
func NewGetStoreProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStoreProductDetailLogic {
	return &GetStoreProductDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStoreProductDetailLogic) GetStoreProductDetail(req *types.StoreProductDetailReq) (resp *types.StoreProductItem, err error) {
	// todo: add your logic here and delete this line

	return
}
