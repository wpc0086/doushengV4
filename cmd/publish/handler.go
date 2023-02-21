package main

import (
	"context"
	"doushengV4/cmd/publish/pack"
	"doushengV4/cmd/publish/service"
	publish "doushengV4/kitex_gen/publish"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// FeedPublish implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) FeedPublish(ctx context.Context, req *publish.FeedRequest) (resp *publish.FeedResponse, err error) {
	pulishes, err := service.NewFeedPublishService(ctx).FeedPulish(req)
	if err != nil {
		resp = pack.BuildFeedResp(err)
		return resp, nil
	}
	return &publish.FeedResponse{StatusCode: int32(publish.ErrCode_SuccessCode), VideoList: pulishes}, nil
}

// ActionPublish implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) ActionPublish(ctx context.Context, req *publish.ActionRequest) (resp *publish.ActionResp, err error) {
	err = service.NewActionPublishService(ctx).ActionPulish(req)
	if err != nil {
		resp = pack.BuildActionResp(err)
		return resp, nil
	}
	resp = &publish.ActionResp{StatusCode: int32(publish.ErrCode_SuccessCode)}
	return resp, nil
}

// ListPublish implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) ListPublish(ctx context.Context, req *publish.ListRequest) (resp *publish.ListResp, err error) {
	pulishes, err := service.NewListPublishService(ctx).ListPulish(req)
	if err != nil {
		resp = pack.BuildListResp(err)
		return resp, nil
	}
	return &publish.ListResp{StatusCode: int32(publish.ErrCode_SuccessCode), VideoList: pulishes}, nil
}
