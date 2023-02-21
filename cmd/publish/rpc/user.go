package rpc

import (
	"context"
	"doushengV4/kitex_gen/user"
	"doushengV4/kitex_gen/user/userservice"
	"doushengV4/pkg/consts"
	"doushengV4/pkg/errno"
	"doushengV4/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func initUser() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	// build a new CBSuite
	//配置熔断器
	cbs := circuitbreak.NewCBSuite(circuitbreak.RPCInfo2Key)
	c, err := userservice.NewClient(
		consts.UserServiceName,
		client.WithResolver(r),
		client.WithCircuitBreaker(cbs),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

func RegisterUser(ctx context.Context, req *user.RegisterUserRequest) (*user.RegisterResp, error) {
	resp, err := userClient.RegisterUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}
	return resp, nil
}

func LoginUser(ctx context.Context, req *user.LoginUserRequest) (*user.LoginResp, error) {
	resp, err := userClient.LoginUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}
	return resp, nil
}

func InfoUser(ctx context.Context, req *user.InfoUserRequest) (*user.InfoUserResponse, error) {
	resp, err := userClient.InforUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}
	return resp, nil
}
