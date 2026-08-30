// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package order

import (
	"context"

	"goship/backend/internal/svc"
	"goship/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FulfillOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 后台实物订单发货 (回填运单号)
func NewFulfillOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FulfillOrderLogic {
	return &FulfillOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FulfillOrderLogic) FulfillOrder(req *types.AdminFulfillOrderReq) (resp *types.NilResp, err error) {
	// todo: add your logic here and delete this line

	return
}
