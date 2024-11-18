package main

import (
	"flag"
	"fmt"

	"im/application/email/rpc/internal/config"
	"im/application/email/rpc/internal/server"
	"im/application/email/rpc/internal/svc"
	"im/application/email/rpc/service"

	"github.com/zeromicro/go-zero/core/conf"
	zservice "github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/email.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		service.RegisterEmailServer(grpcServer, server.NewEmailServer(ctx))

		if c.Mode == zservice.DevMode || c.Mode == zservice.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	_ = consul.RegisterService(c.ListenOn, c.Consul)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
