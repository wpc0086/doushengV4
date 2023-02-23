package service

import (
	"context"
	"doushengV4/cmd/publish/dal/db"
	"doushengV4/cmd/publish/dal/redis"
	"doushengV4/cmd/publish/pack"
	"doushengV4/cmd/publish/rpc"
	"doushengV4/kitex_gen/publish"
	"doushengV4/kitex_gen/user"
	"doushengV4/pkg/consts"
)

type FeedPublishService struct {
	ctx context.Context
}

func NewFeedPublishService(ctx context.Context) *FeedPublishService {
	return &FeedPublishService{ctx: ctx}
}

func (s *FeedPublishService) FeedPulish(req *publish.FeedRequest) ([]*publish.Video, error) {
	videos, err := redis.GetFeed(s.ctx, *req.LatestTime)
	if videos == nil { //未命中Redis,查询数据库，设置Redis锁，防止缓存击穿
		resourceKey := consts.RedisFeedLockPre
		routine := consts.Rountiue
		acquired, err := redis.RedisClient.SetNX(s.ctx, resourceKey, routine, consts.RedisVideoLockExp).Result()
		defer redis.ReleaseLock(s.ctx, int64(routine), resourceKey)
		if err != nil {
			return nil, err
		}
		if acquired {
			videos, err = db.GetFeed(*req.LatestTime)
			for _, video := range videos {
				redis.SaveVideoAndImage(s.ctx, video)
			}
		}
	}
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
