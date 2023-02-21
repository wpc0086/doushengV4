// Code generated by Kitex v0.4.4. DO NOT EDIT.

package interactservice

import (
	"context"
	interact "doushengV4/kitex_gen/interact"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return interactServiceServiceInfo
}

var interactServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "InteractService"
	handlerType := (*interact.InteractService)(nil)
	methods := map[string]kitex.MethodInfo{
		"FavoriteAction": kitex.NewMethodInfo(favoriteActionHandler, newInteractServiceFavoriteActionArgs, newInteractServiceFavoriteActionResult, false),
		"FavoriteList":   kitex.NewMethodInfo(favoriteListHandler, newInteractServiceFavoriteListArgs, newInteractServiceFavoriteListResult, false),
		"CommentAction":  kitex.NewMethodInfo(commentActionHandler, newInteractServiceCommentActionArgs, newInteractServiceCommentActionResult, false),
		"CommentList":    kitex.NewMethodInfo(commentListHandler, newInteractServiceCommentListArgs, newInteractServiceCommentListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "interact",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func favoriteActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*interact.InteractServiceFavoriteActionArgs)
	realResult := result.(*interact.InteractServiceFavoriteActionResult)
	success, err := handler.(interact.InteractService).FavoriteAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newInteractServiceFavoriteActionArgs() interface{} {
	return interact.NewInteractServiceFavoriteActionArgs()
}

func newInteractServiceFavoriteActionResult() interface{} {
	return interact.NewInteractServiceFavoriteActionResult()
}

func favoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*interact.InteractServiceFavoriteListArgs)
	realResult := result.(*interact.InteractServiceFavoriteListResult)
	success, err := handler.(interact.InteractService).FavoriteList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newInteractServiceFavoriteListArgs() interface{} {
	return interact.NewInteractServiceFavoriteListArgs()
}

func newInteractServiceFavoriteListResult() interface{} {
	return interact.NewInteractServiceFavoriteListResult()
}

func commentActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*interact.InteractServiceCommentActionArgs)
	realResult := result.(*interact.InteractServiceCommentActionResult)
	success, err := handler.(interact.InteractService).CommentAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newInteractServiceCommentActionArgs() interface{} {
	return interact.NewInteractServiceCommentActionArgs()
}

func newInteractServiceCommentActionResult() interface{} {
	return interact.NewInteractServiceCommentActionResult()
}

func commentListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*interact.InteractServiceCommentListArgs)
	realResult := result.(*interact.InteractServiceCommentListResult)
	success, err := handler.(interact.InteractService).CommentList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newInteractServiceCommentListArgs() interface{} {
	return interact.NewInteractServiceCommentListArgs()
}

func newInteractServiceCommentListResult() interface{} {
	return interact.NewInteractServiceCommentListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) FavoriteAction(ctx context.Context, req *interact.FavoriteActionRequest) (r *interact.FavoriteActionResponse, err error) {
	var _args interact.InteractServiceFavoriteActionArgs
	_args.Req = req
	var _result interact.InteractServiceFavoriteActionResult
	if err = p.c.Call(ctx, "FavoriteAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteList(ctx context.Context, req *interact.FavoriteListRequest) (r *interact.FavoriteListResponse, err error) {
	var _args interact.InteractServiceFavoriteListArgs
	_args.Req = req
	var _result interact.InteractServiceFavoriteListResult
	if err = p.c.Call(ctx, "FavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentAction(ctx context.Context, req *interact.CommentActionRequest) (r *interact.CommentActionResponse, err error) {
	var _args interact.InteractServiceCommentActionArgs
	_args.Req = req
	var _result interact.InteractServiceCommentActionResult
	if err = p.c.Call(ctx, "CommentAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentList(ctx context.Context, req *interact.CommentListRequest) (r *interact.CommentListResponse, err error) {
	var _args interact.InteractServiceCommentListArgs
	_args.Req = req
	var _result interact.InteractServiceCommentListResult
	if err = p.c.Call(ctx, "CommentList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
