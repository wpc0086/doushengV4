package main

import (
	"context"
	"doushengV4/cmd/interact/dal/redis"
	"doushengV4/cmd/interact/pack"
	"doushengV4/cmd/interact/rpc"
	"doushengV4/cmd/interact/service"
	interact "doushengV4/kitex_gen/interact"
	"doushengV4/kitex_gen/user"
	"doushengV4/pkg/errno"
	"strconv"
)

// InteractServiceImpl implements the last service interface defined in the IDL.
type InteractServiceImpl struct{}

// FavoriteAction implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) FavoriteAction(ctx context.Context, req *interact.FavoriteActionRequest) (resp *interact.FavoriteActionResponse, err error) {
	err = service.NewActionFavoriteService(ctx).ActionFavorite(req)
	if err != nil {
		resp = pack.BuildFavoriteActionResp(err)
		return resp, nil
	}
	return &interact.FavoriteActionResponse{StatusCode: int32(interact.ErrCode_SuccessCode)}, nil
}

// FavoriteList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) FavoriteList(ctx context.Context, req *interact.FavoriteListRequest) (resp *interact.FavoriteListResponse, err error) {
	if redis.GetErr(ctx, req.Token) == true { //防止缓存穿透
		return nil, errno.DataErr
	}
	videos, err := service.NewListFavoriteService(ctx).ListFavorite(req)
	if err != nil {
		resp = pack.BuildFavoriteListResp(err)
		redis.SaveErr(ctx, req.Token) //防止缓存穿透
		return resp, nil
	}
	return &interact.FavoriteListResponse{StatusCode: int32(interact.ErrCode_SuccessCode), VideoList: videos}, nil
}

// CommentAction implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentAction(ctx context.Context, req *interact.CommentActionRequest) (resp *interact.CommentActionResponse, err error) {
	comment, err := service.NewActionCommentService(ctx).ActionComment(req)
	if err != nil {
		resp = pack.BuildCommentActionResp(err)
		return resp, nil
	}
	rComment := pack.Comment(comment)
	infoUser, err := rpc.InfoUser(ctx, &user.InfoUserRequest{UserId: comment.UserId})
	if err != nil {
		resp = pack.BuildCommentActionResp(err)
		return resp, nil
	}
	rComment.User = pack.User(infoUser.User)
	return &interact.CommentActionResponse{StatusCode: int32(interact.ErrCode_SuccessCode), Comment: rComment}, nil
}

// CommentList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentList(ctx context.Context, req *interact.CommentListRequest) (resp *interact.CommentListResponse, err error) {
	if redis.GetErr(ctx, strconv.FormatInt(req.VideoId, 10)) == true { //防止缓存穿透
		return nil, errno.DataErr
	}
	comment, err := service.NewListCommentService(ctx).ListComment(req)
	if err != nil {
		resp = pack.BuildCommentListResp(err)
		redis.SaveErr(ctx, strconv.FormatInt(req.VideoId, 10)) //防止缓存穿透
		return resp, nil
	}
	return &interact.CommentListResponse{StatusCode: int32(interact.ErrCode_SuccessCode), CommentList: comment}, nil
}
