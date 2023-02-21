package service

import (
	"context"
	"doushengV4/cmd/publish/dal/db"
	"doushengV4/cmd/publish/pack"
	"doushengV4/cmd/publish/rpc"
	"doushengV4/kitex_gen/publish"
	"doushengV4/kitex_gen/user"
)

type ListPublishService struct {
	ctx context.Context
}

func NewListPublishService(ctx context.Context) *ListPublishService {
	return &ListPublishService{ctx: ctx}
}

func (s *ListPublishService) ListPulish(req *publish.ListRequest) ([]*publish.Video, error) {
	user_id := req.UserId
	//查询用户发布的视频列表
	videos, err := db.GetVideoListByUserId(user_id)
	if err != nil {
		return nil, err
	}
	//根据用户id查询视频作者信息
	userResp, err := rpc.InfoUser(s.ctx, &user.InfoUserRequest{UserId: user_id})
	if err != nil {
		userResp.User = nil
	}
	//对每个video赋值author
	publishes := pack.Publishs(videos)
	for _, p := range publishes {
		p.Author = pack.User(userResp.User)
	}

	return publishes, nil
}
