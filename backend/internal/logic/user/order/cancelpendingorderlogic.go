// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package order

import (
	"context"

	"goship/backend/internal/svc"
	"goship/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelPendingOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 前台用户取消未支付的 pending 订单
func NewCancelPendingOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelPendingOrderLogic {
	return &CancelPendingOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelPendingOrderLogic) CancelPendingOrder(req *types.CancelOrderReq) (resp *types.NilResp, err error) {
	// todo: add your logic here and delete this line

	return
}
