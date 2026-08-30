// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"goship/backend/internal/svc"
	"goship/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdjustUserCreditsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 后台手动增减用户点数
func NewAdjustUserCreditsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdjustUserCreditsLogic {
	return &AdjustUserCreditsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdjustUserCreditsLogic) AdjustUserCredits(req *types.AdminAdjustCreditsReq) (resp *types.NilResp, err error) {
	// todo: add your logic here and delete this line

	return
}
