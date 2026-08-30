// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"goship/backend/internal/svc"
	"goship/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSystemConfigsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 后台获取系统全局配置
func NewGetSystemConfigsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSystemConfigsLogic {
	return &GetSystemConfigsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSystemConfigsLogic) GetSystemConfigs() (resp *types.AdminConfigListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
