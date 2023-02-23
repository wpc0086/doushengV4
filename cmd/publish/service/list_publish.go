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
	"fmt"
)

type ListPublishService struct {
	ctx context.Context
}

func NewListPublishService(ctx context.Context) *ListPublishService {
	return &ListPublishService{ctx: ctx}
}

func (s *ListPublishService) ListPulish(req *publish.ListRequest) ([]*publish.Video, error) {
	redisCache, _ := redis.GetPublishList(s.ctx, req.UserId)
	//命中redis
	if len(redisCache) != 0 {
		return redisCache, nil
	}
	user_id := req.UserId
	//未命中查询数据库，设置Redis锁，防止缓存击穿
	resourceKey := fmt.Sprintf(consts.RedisVideoLockPre, req.UserId)
	routine := consts.Rountiue
	acquired, err := redis.RedisClient.SetNX(s.ctx, resourceKey, routine, consts.RedisVideoLockExp).Result()
	if err != nil {
		return nil, err
	}
	defer redis.ReleaseLock(s.ctx, int64(routine), resourceKey)
	if acquired {
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
		err = redis.SavePublishList(s.ctx, publishes, req.UserId)

		return publishes, nil
	}
	return nil, nil
}
