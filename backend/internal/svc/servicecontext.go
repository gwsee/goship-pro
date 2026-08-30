// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"goship/backend/dao/query"
	"goship/backend/internal/config"
	"goship/backend/internal/middleware"
	"goship/backend/pkgx/database"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	AdminAuth rest.Middleware
	Loc       *time.Location `json:"Loc,optional"`
	Redis     *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	var err error
	c.DbCfg.Log = c.Log.Path
	gormDb, err := database.NewDatabase(&c.DbCfg)
	if err != nil {
		panic("数据库初始化失败" + err.Error())
	}
	query.SetDefault(gormDb)
	loc, err := time.LoadLocation("Local")
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:    c,
		Loc:       loc,
		AdminAuth: middleware.NewAdminAuthMiddleware().Handle,
		Redis: redis.NewClient(&redis.Options{
			Addr:        c.Redis.Host,
			Username:    c.Redis.User,
			Password:    c.Redis.Pass,
			DialTimeout: time.Second * 2,
		}),
	}
}
