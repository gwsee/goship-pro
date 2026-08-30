// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package order

import (
	"context"

	"goship/backend/internal/svc"
	"goship/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderCheckoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 【核心】前台用户下单并获取支付收银台链接
func NewCreateOrderCheckoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderCheckoutLogic {
	return &CreateOrderCheckoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderCheckoutLogic) CreateOrderCheckout(req *types.CreateOrderCheckoutReq) (resp *types.CreateOrderCheckoutResp, err error) {
	// todo: add your logic here and delete this line

	return
}
