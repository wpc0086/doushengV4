package rpc

import (
	"context"
	"doushengV4/kitex_gen/publish"
	"doushengV4/kitex_gen/publish/publishservice"
	"doushengV4/pkg/consts"
	"doushengV4/pkg/errno"
	"doushengV4/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/tool/internal_pkg/log"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var publishClient publishservice.Client

func initPublish() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		log.Info("==========initPublish error==================")
		panic(err)
	}
	// build a new CBSuite
	//配置熔断器
	cbs := circuitbreak.NewCBSuite(circuitbreak.RPCInfo2Key)
	c, err := publishservice.NewClient(
		consts.PublishServiceName,
		client.WithResolver(r),
		client.WithCircuitBreaker(cbs),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	publishClient = c
}

func ActionPublic(ctx context.Context, req *publish.ActionRequest) (*publish.ActionResp, error) {
	resp, err := publishClient.ActionPublish(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}
	return resp, nil
}

func ListPublish(ctx context.Context, req *publish.ListRequest) (*publish.ListResp, error) {
	resp, err := publishClient.ListPublish(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}
	return resp, nil
}

func FeedPublish(ctx context.Context, req *publish.FeedRequest) (*publish.FeedResponse, error) {
	resp, err := publishClient.FeedPublish(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}
	return resp, nil
}
