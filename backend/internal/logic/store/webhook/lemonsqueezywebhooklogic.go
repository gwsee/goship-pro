// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package webhook

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"goship/backend/internal/svc"
)

type LemonSqueezyWebhookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Lemon Squeezy 支付成功回调通知
func NewLemonSqueezyWebhookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LemonSqueezyWebhookLogic {
	return &LemonSqueezyWebhookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LemonSqueezyWebhookLogic) LemonSqueezyWebhook() error {
	// todo: add your logic here and delete this line

	return nil
}
