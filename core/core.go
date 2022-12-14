package main

import (
	"cherry-disk/core/common"
	"flag"
	"fmt"

	"cherry-disk/core/internal/config"
	"cherry-disk/core/internal/handler"
	"cherry-disk/core/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

// var configFile = flag.String("f", "etc/core-api.yaml", "the config file")
var configFile = flag.String("f", "/Users/edj/dev/go_project/git_project/cherry_disk/core/etc/core-api.yaml", "the config file")

func main() {
	//获取nacos配置
	common.GetConfig()
	//读取yaml配置
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	config.Banner()
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
