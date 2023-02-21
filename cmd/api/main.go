package main

import (
	"doushengV4/cmd/api/routes"
	rpc "doushengV4/cmd/api/rpc"
	"doushengV4/cmd/api/service"
	"github.com/gin-gonic/gin"
)

func Init() {
	rpc.Init()
}
func main() {
	go service.RunMessageServer()
	Init()
	r := gin.Default()
	routes.InitRouter(r)
	errrun := r.Run()
	if errrun != nil {
		return
	}
}
