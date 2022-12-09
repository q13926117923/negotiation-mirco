package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"user/conf"
	"user/core"
	pb "user/services/proto"
)

func main()  {
	conf.Init()
	etcdReg := etcd.NewRegistry( //选项模式
		registry.Addrs("127.0.0.1:2379"),
		)
	microService := micro.NewService(
		micro.Name("rpcUserService"),
		micro.Address("127.0.0.1:8082"),
		micro.Registry(etcdReg),
		)
	microService.Init()
	//服务注册
	_ = pb.RegisterUserServiceHandler(microService.Server(),new(core.UserService))
	//启动
	_ = microService.Run()
}
