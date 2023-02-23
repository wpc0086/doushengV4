package redis

import (
	"context"
	"doushengV4/cmd/publish/dal/db"
	"doushengV4/kitex_gen/publish"
	"doushengV4/pkg/consts"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"time"
)

func SaveVideoAndImage(ctx context.Context, video *db.Video) (err error) {
	// 序列化视频对象为 JSON 字符串
	videoJson, err := json.Marshal(video)
	if err != nil {
		return err
	}
	// 保存视频到 Redis ZSET，以 CreatedAt 作为 score
	zsetItem := &redis.Z{
		Score:  float64(video.CreatedAt),
		Member: videoJson,
	}
	err = RedisClient.ZAdd(ctx, consts.RedisFeedKey, zsetItem).Err()
	RedisClient.Expire(ctx, consts.RedisFeedKey, consts.RedisFeedExp)
	if err != nil {
		return err
	}

	return nil
}

func GetFeed(ctx context.Context, lastTime int64) (feedList []*db.Video, err error) {
	// 获取 Redis 中的视频列表
	videoJsons, err := RedisClient.ZRevRangeByScore(ctx, consts.RedisFeedKey, &redis.ZRangeBy{
		Min:    "0",
		Max:    fmt.Sprintf("%d", lastTime),
		Offset: 0,
		Count:  30,
	}).Result()
	if err != nil {
		return nil, err
	}

	// 将 JSON 转为 Video 对象
	for _, videoJson := range videoJsons {
		video := &db.Video{}
		if err := json.Unmarshal([]byte(videoJson), video); err != nil {
			return nil, err
		}
		feedList = append(feedList, video)
	}

	return feedList, nil
}
func RemovePublishList(ctx context.Context, userId int64) {
	key := fmt.Sprintf(consts.RedisVideoListPre, userId)
	RedisClient.Del(ctx, key)
}

func SavePublishList(ctx context.Context, video []*publish.Video, userId int64) (err error) {
	videoList := fmt.Sprintf(consts.RedisVideoListPre, userId)
	data, err := json.Marshal(&video)
	if err != nil {
		return err
	}
	RedisClient.Set(ctx, videoList, data, 0)
	//防止缓存雪崩
	RedisClient.Expire(ctx, videoList, time.Duration(int64(time.Second)*(11+rand.Int63n(10))))
	return nil
}

func GetPublishList(ctx context.Context, userId int64) (video []*publish.Video, err error) {
	videoList := fmt.Sprintf(consts.RedisVideoListPre, userId)
	videoJson, err := RedisClient.Get(ctx, videoList).Result()
	err = json.Unmarshal([]byte(videoJson), &video)
	if err != nil {
		return nil, err
	}
	//防止缓存雪崩
	return video, nil
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
