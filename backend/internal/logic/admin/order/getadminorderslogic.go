// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package order

import (
	"context"

	"goship/backend/internal/svc"
	"goship/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminOrdersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 后台获取订单列表
func NewGetAdminOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminOrdersLogic {
	return &GetAdminOrdersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminOrdersLogic) GetAdminOrders(req *types.AdminOrderListReq) (resp *types.AdminOrderListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
