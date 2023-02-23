package service

import (
	"context"
	"doushengV4/cmd/user/dal/db"
	"doushengV4/cmd/user/dal/redis"
	"doushengV4/cmd/user/pack"
	"doushengV4/kitex_gen/user"
	"doushengV4/pkg/consts"
	"fmt"
)

type InfoUserService struct {
	ctx context.Context
}

func NewInfoUserService(ctx context.Context) *InfoUserService {
	return &InfoUserService{
		ctx: ctx,
	}
}

func (s *InfoUserService) InfoUser(req *user.InfoUserRequest) (user *user.User, err error) {
	redisUser, err := redis.GetInfo(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	//命中Redis
	if redisUser.Id != 0 {
		return redisUser, nil
	}
	//未命中查询数据库，设置Redis锁，防止缓存击穿
	resourceKey := fmt.Sprintf(consts.RedisUserLockPre, req.UserId)
	routine := consts.Rountiue
	acquired, err := redis.RedisClient.SetNX(s.ctx, resourceKey, routine, consts.RedisUserLockExp).Result()
	defer redis.ReleaseLock(s.ctx, int64(routine), resourceKey)
	if err != nil {
		return nil, err
	}
	if acquired {
		u, err := db.GetUserById(req.UserId)
		if err != nil {
			return nil, err
		}
		u.Password = ""
		infoUser := pack.User(u)
		//查询视频表获取作品数
		videos, err := db.GetVideoListByUserId(infoUser.Id)
		if err != nil {
			return infoUser, err
		}
		infoUser.WorkCount = int64(len(videos))
		//获赞数量
		infoUser.TotalFavorited = 0
		for _, video := range videos {
			infoUser.TotalFavorited += int64(video.FavoriteCount)
		}
		//查询喜欢表获取喜欢数
		favoiteCount, _ := db.GetFavoritesByUserId(infoUser.Id)
		infoUser.FavoriteCount = favoiteCount
		//保存到Redis
		err = redis.SaveInfo(s.ctx, infoUser, req.UserId)
		if err != nil {
			return nil, err
		}
		return infoUser, nil
	}
	return nil, nil
}
