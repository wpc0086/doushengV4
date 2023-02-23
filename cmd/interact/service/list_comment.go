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

type ListCommentService struct {
	ctx context.Context
}

func NewListCommentService(ctx context.Context) *ListCommentService {
	return &ListCommentService{ctx: ctx}
}

func (s *ListCommentService) ListComment(req *interact.CommentListRequest) ([]*interact.Comment, error) {
	redisCache, _ := redis.GetCommentList(s.ctx, req.VideoId)
	//命中redis
	if redisCache != nil {
		return redisCache, nil
	}
	//未命中查询数据库，设置Redis锁，防止缓存击穿
	resourceKey := fmt.Sprintf(consts.RedisCommentLockPre, req.VideoId)
	routine := consts.Rountiue
	acquired, err := redis.RedisClient.SetNX(s.ctx, resourceKey, routine, consts.RedisCommentLockExp).Result()
	if err != nil {
		return nil, err
	}
	defer redis.ReleaseLock(s.ctx, int64(routine), resourceKey)

	if acquired {
		cs, err := db.GetCommonsByVideoID(req.VideoId)
		if err != nil {
			return nil, err
		}
		comments := pack.Comments(cs)
		for i, c := range cs {
			infoUser, err := rpc.InfoUser(s.ctx, &user.InfoUserRequest{UserId: c.UserId})
			if err != nil {
				continue
			}
			comments[i].User = pack.User(infoUser.User)
		}
		err = redis.SaveCommentList(s.ctx, comments, req.VideoId)
		return comments, nil
	}
	return nil, nil
}
