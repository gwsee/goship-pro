// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package profile

import (
	"context"

	"goship/backend/internal/svc"
	"goship/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AIGenerateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// AI 流式生成演示 (SSE)
func NewAIGenerateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AIGenerateLogic {
	return &AIGenerateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AIGenerateLogic) AIGenerate(req *types.AIGenerateReq) error {
	// todo: add your logic here and delete this line

	return nil
}
