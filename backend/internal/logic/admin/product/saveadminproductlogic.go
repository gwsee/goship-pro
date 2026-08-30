// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package product

import (
	"context"

	"goship/backend/internal/svc"
	"goship/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveAdminProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 后台新增/修改商品
func NewSaveAdminProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveAdminProductLogic {
	return &SaveAdminProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveAdminProductLogic) SaveAdminProduct(req *types.AdminSaveProductReq) (resp *types.NilResp, err error) {
	// todo: add your logic here and delete this line

	return
}
