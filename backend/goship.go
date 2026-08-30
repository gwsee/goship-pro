// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package main

import (
	"flag"
	"goship/backend/internal/config"
	"goship/backend/internal/handler"
	"goship/backend/internal/svc"
	"os"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/goship.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.MustSetup(c.Log)
	if c.Mode == service.DevMode || c.Mode == service.TestMode {
		logx.AddWriter(logx.NewWriter(os.Stdout))
	}
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	logx.Infof("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
