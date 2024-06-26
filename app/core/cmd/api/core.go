package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zero-zone/app/core/cmd/api/internal/config"
	"zero-zone/app/core/cmd/api/internal/handler"
	"zero-zone/app/core/cmd/api/internal/svc"
	"zero-zone/app/pkg/errorx"
	"zero-zone/app/pkg/zlog"
)

var configFile = flag.String("f", "etc/core-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.OldRegisterHandlers(server, ctx)
	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	if c.Mode == "dev" {
		//logx.DisableStat()
	}

	io := zlog.NewRedisIoWriter("redis.io.writer", redis.RedisConf{
		Host:        c.Redis.Host,
		Type:        c.Redis.Type,
		Pass:        c.Redis.Pass,
		Tls:         false,
		NonBlock:    false,
		PingTimeout: 0,
	})
	logx.SetWriter(logx.NewWriter(io))

	// -------------------------------- info --------------------------

	logx.Info("测试", "info")
	logx.Infof("测试  %v", "info")

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
