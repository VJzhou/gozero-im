package main

import (
	"flag"
	"fmt"

	"im/application/friend/rpc/internal/config"
	"im/application/friend/rpc/internal/server"
	"im/application/friend/rpc/internal/svc"
	"im/application/friend/rpc/service"

	"github.com/zeromicro/go-zero/core/conf"
	zeroService "github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/friend.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		service.RegisterFriendServer(grpcServer, server.NewFriendServer(ctx))

		if c.Mode == zeroService.DevMode || c.Mode == zeroService.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	_ = consul.RegisterService(c.ListenOn, c.Consul)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
