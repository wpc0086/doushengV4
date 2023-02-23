package redis

import (
	"context"
	"doushengV4/kitex_gen/interact"
	"doushengV4/pkg/consts"
	"encoding/json"
	"fmt"
	"time"
)

func ModifyFavorite(ctx context.Context, userId int64, num int64) {
	userInfo := fmt.Sprintf(consts.RedisUserInfoPre, userId)
	_, err := RedisClient.HGet(ctx, userInfo, "favorite_count").Result()
	if err != nil {
		return
	}
	RedisClient.HIncrBy(ctx, userInfo, "favorite_count", num)
}
func RemoveFavoriteList(ctx context.Context, userId int64) {
	key := fmt.Sprintf(consts.RedisFavoriteListPre, userId)
	RedisClient.Del(ctx, key)
}
func SaveFavoriteList(ctx context.Context, video []*interact.Video, userId int64) (err error) {
	videoList := fmt.Sprintf(consts.RedisFavoriteListPre, userId)
	data, err := json.Marshal(&video)
	if err != nil {
		return err
	}
	RedisClient.Set(ctx, videoList, data, 0)
	//防止缓存雪崩
	RedisClient.Expire(ctx, videoList, time.Duration(consts.RedisFavoriteListExp))
	return nil
}

func GetFavoriteList(ctx context.Context, userId int64) (video []*interact.Video, err error) {
	videoList := fmt.Sprintf(consts.RedisFavoriteListPre, userId)
	videoJson, err := RedisClient.Get(ctx, videoList).Result()
	err = json.Unmarshal([]byte(videoJson), &video)
	if err != nil {
		return nil, err
	}
	//防止缓存雪崩
	return video, nil
}

func RemoveCommentList(ctx context.Context, userId int64) {
	key := fmt.Sprintf(consts.RedisCommentListPre, userId)
	RedisClient.Del(ctx, key)
}
func SaveCommentList(ctx context.Context, comments []*interact.Comment, userId int64) (err error) {
	commentsList := fmt.Sprintf(consts.RedisCommentListPre, userId)
	data, err := json.Marshal(&comments)
	if err != nil {
		return err
	}
	RedisClient.Set(ctx, commentsList, data, 0)
	//防止缓存雪崩
	RedisClient.Expire(ctx, commentsList, time.Duration(consts.RedisCommentListExp))
	return nil
}

func GetCommentList(ctx context.Context, commentsId int64) (comments []*interact.Comment, err error) {
	commentsList := fmt.Sprintf(consts.RedisCommentListPre, commentsId)
	commentsJson, err := RedisClient.Get(ctx, commentsList).Result()
	err = json.Unmarshal([]byte(commentsJson), &comments)
	if err != nil {
		return nil, err
	}
	//防止缓存雪崩
	return comments, nil
}
func SaveErr(ctx context.Context, parameter string) {
	key := fmt.Sprintf(consts.RediseErrorPre, parameter)
	RedisClient.Set(ctx, key, parameter, consts.RediseErrorExp)
}

func GetErr(ctx context.Context, parameter string) bool {
	key := fmt.Sprintf(consts.RediseErrorPre, parameter)
	result, _ := RedisClient.Get(ctx, key).Result()
	if len(result) != 0 {
		return true
	}
	return false
}
