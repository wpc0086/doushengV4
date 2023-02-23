package main

import (
	"doushengV4/cmd/publish/dal"
	"doushengV4/cmd/publish/rpc"
	"doushengV4/cmd/publish/util"
	"doushengV4/kitex_gen/publish/publishservice"
	"doushengV4/pkg/consts"
	"doushengV4/pkg/mw"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
)

func Init() {
	dal.Init()
	util.InitMinio()
	rpc.Init()
}
func main() {
	r, err := etcd.NewEtcdRegistry([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr(consts.TCP, consts.PublishServiceAddr)
	if err != nil {
		panic(err)
	}
	Init()

	svr := publishservice.NewServer(new(PublishServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 1000}),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.PublishServiceName}),
	)
	err = svr.Run()
	if err != nil {
		return
	}
}
