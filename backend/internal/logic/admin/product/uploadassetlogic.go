// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package product

import (
	"context"

	"goship/backend/internal/svc"
	"goship/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadAssetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 直传 Cloudflare R2 图片
func NewUploadAssetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadAssetLogic {
	return &UploadAssetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadAssetLogic) UploadAsset() (resp *types.AdminUploadResp, err error) {
	// todo: add your logic here and delete this line

	return
}
