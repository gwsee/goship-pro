// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package order

import (
	"context"

	"goship/backend/internal/svc"
	"goship/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckOrderStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 前台支付结果轮询(付款跳回后检查是否已到账)
func NewCheckOrderStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckOrderStatusLogic {
	return &CheckOrderStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckOrderStatusLogic) CheckOrderStatus(req *types.UserOrderStatusReq) (resp *types.UserOrderStatusResp, err error) {
	// todo: add your logic here and delete this line

	return
}
