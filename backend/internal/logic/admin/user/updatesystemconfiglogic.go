// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"goship/backend/internal/svc"
	"goship/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSystemConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 后台修改系统全局配置
func NewUpdateSystemConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSystemConfigLogic {
	return &UpdateSystemConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSystemConfigLogic) UpdateSystemConfig(req *types.AdminUpdateConfigReq) (resp *types.NilResp, err error) {
	// todo: add your logic here and delete this line

	return
}
