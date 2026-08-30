// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package auth

import (
	"context"

	"goship/backend/internal/svc"
	"goship/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoogleAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Google OAuth 一键登录
func NewGoogleAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoogleAuthLogic {
	return &GoogleAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoogleAuthLogic) GoogleAuth(req *types.GoogleAuthReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	return
}
