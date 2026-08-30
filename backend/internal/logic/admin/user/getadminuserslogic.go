// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"goship/backend/internal/svc"
	"goship/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 后台查看注册用户列表
func NewGetAdminUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminUsersLogic {
	return &GetAdminUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminUsersLogic) GetAdminUsers(req *types.AdminUserListReq) (resp *types.AdminUserListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
