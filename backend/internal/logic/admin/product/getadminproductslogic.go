// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package product

import (
	"context"

	"goship/backend/internal/svc"
	"goship/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminProductsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 后台获取商品列表
func NewGetAdminProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminProductsLogic {
	return &GetAdminProductsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminProductsLogic) GetAdminProducts(req *types.AdminProductListReq) (resp *types.AdminProductListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
