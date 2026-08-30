// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package product

import (
	"context"

	"goship/backend/internal/svc"
	"goship/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ToggleProductStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 后台上下架切换
func NewToggleProductStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ToggleProductStatusLogic {
	return &ToggleProductStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ToggleProductStatusLogic) ToggleProductStatus(req *types.AdminToggleStatusReq) (resp *types.NilResp, err error) {
	// todo: add your logic here and delete this line

	return
}
