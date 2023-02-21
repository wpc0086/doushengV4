package service

import (
	"context"
	"doushengV4/cmd/publish/dal/db"
	"doushengV4/cmd/publish/pack"
	"doushengV4/cmd/publish/rpc"
	"doushengV4/kitex_gen/publish"
	"doushengV4/kitex_gen/user"
)

type FeedPublishService struct {
	ctx context.Context
}

func NewFeedPublishService(ctx context.Context) *FeedPublishService {
	return &FeedPublishService{ctx: ctx}
}

func (s *FeedPublishService) FeedPulish(req *publish.FeedRequest) ([]*publish.Video, error) {
	videos, err := db.GetFeed(*req.LatestTime)
	if err != nil {
		return nil, err
	}
	publishes := pack.Publishs(videos)
	//根据用户id查询视频作者信息
	for i, v := range videos {
		userResp, err := rpc.InfoUser(s.ctx, &user.InfoUserRequest{UserId: v.AuthorId})
		if err != nil {
			continue
		}
		publishes[i].Author = pack.User(userResp.User)
	}

	return publishes, nil
}
