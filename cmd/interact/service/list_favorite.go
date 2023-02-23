package service

import (
	"context"
	"doushengV4/cmd/interact/dal/db"
	"doushengV4/cmd/interact/dal/redis"
	"doushengV4/cmd/interact/pack"
	"doushengV4/cmd/interact/rpc"
	"doushengV4/kitex_gen/interact"
	"doushengV4/kitex_gen/user"
	"doushengV4/pkg/consts"
	"fmt"
)

type ListFavoriteService struct {
	ctx context.Context
}

func NewListFavoriteService(ctx context.Context) *ListFavoriteService {
	return &ListFavoriteService{ctx: ctx}
}

func (s *ListFavoriteService) ListFavorite(req *interact.FavoriteListRequest) ([]*interact.Video, error) {
	userId := req.UserId
	redisCache, _ := redis.GetFavoriteList(s.ctx, userId)
	//命中redis
	if redisCache != nil {
		return redisCache, nil
	}
	//未命中查询数据库，设置Redis锁，防止缓存击穿
	resourceKey := fmt.Sprintf(consts.RedisFavoriteLockPre, userId)
	routine := consts.Rountiue
	acquired, err := redis.RedisClient.SetNX(s.ctx, resourceKey, routine, consts.RedisFavoriteLockExp).Result()
	if err != nil {
		return nil, err
	}
	defer redis.ReleaseLock(s.ctx, int64(routine), resourceKey)
	if acquired {
		videos, err := db.GetVideosByUserId(userId)
		if err != nil {
			return nil, err
		}
		Favoritees := pack.Publishs(videos)
		for i, video := range videos {
			infoUser, err := rpc.InfoUser(s.ctx, &user.InfoUserRequest{UserId: video.AuthorId})
			if err != nil {
				continue
			}
			Favoritees[i].Author = pack.User(infoUser.User)
		}
		err = redis.SaveFavoriteList(s.ctx, Favoritees, req.UserId)
		return Favoritees, nil
	}
	return nil, nil
}
