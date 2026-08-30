// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package config

import (
	"goship/backend/pkgx/database"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth  Auth            `json:"Auth,optional"`
	Redis redis.RedisConf `json:"Redis,optional"`
	DbCfg database.Config `json:"DbCfg,optional"`
}
type Auth struct {
	AccessSecret string
	AccessExpire int64
}
