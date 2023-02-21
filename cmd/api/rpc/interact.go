package rpc

import (
	"context"
	"doushengV4/kitex_gen/interact"
	"doushengV4/kitex_gen/interact/interactservice"
	"doushengV4/pkg/consts"
	"doushengV4/pkg/errno"
	"doushengV4/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var interactClient interactservice.Client

func initInteract() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	// build a new CBSuite
	//配置熔断器
	cbs := circuitbreak.NewCBSuite(circuitbreak.RPCInfo2Key)
	c, err := interactservice.NewClient(
		consts.InterActServiceName,
		client.WithResolver(r),
		client.WithCircuitBreaker(cbs),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	interactClient = c
}
func FavoriteAction(ctx context.Context, req *interact.FavoriteActionRequest) (*interact.FavoriteActionResponse, error) {
	resp, err := interactClient.FavoriteAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}
	return resp, nil
}

func FavoriteList(ctx context.Context, req *interact.FavoriteListRequest) (*interact.FavoriteListResponse, error) {
	resp, err := interactClient.FavoriteList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}
	return resp, nil
}

func CommentAction(ctx context.Context, req *interact.CommentActionRequest) (*interact.CommentActionResponse, error) {
	resp, err := interactClient.CommentAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}
	return resp, nil
}

func CommentList(ctx context.Context, req *interact.CommentListRequest) (*interact.CommentListResponse, error) {
	resp, err := interactClient.CommentList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(int64(resp.StatusCode), resp.StatusMsg)
	}
	return resp, nil
}
