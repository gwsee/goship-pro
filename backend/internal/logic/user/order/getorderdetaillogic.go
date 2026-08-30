// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package order

import (
	"context"

	"goship/backend/internal/svc"
	"goship/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 前台用户查看单个订单详情(含物流与收件地址)
func NewGetOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderDetailLogic {
	return &GetOrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderDetailLogic) GetOrderDetail(req *types.UserOrderDetailReq) (resp *types.UserOrderDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
