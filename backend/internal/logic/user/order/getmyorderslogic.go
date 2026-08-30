// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package order

import (
	"context"

	"goship/backend/internal/svc"
	"goship/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyOrdersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 前台用户分页查询我的订单列表
func NewGetMyOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyOrdersLogic {
	return &GetMyOrdersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyOrdersLogic) GetMyOrders(req *types.UserOrderListReq) (resp *types.UserOrderListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
