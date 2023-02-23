package redis

import (
	"context"
	"doushengV4/kitex_gen/user"
	"doushengV4/pkg/consts"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func SaveInfo(ctx context.Context, user *user.User, userId int64) (err error) {
	userInfo := fmt.Sprintf(consts.RedisUserInfoPre, userId)
	data, _ := json.Marshal(&user)
	m := make(map[string]interface{})
	err = json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	RedisClient.HMSet(ctx, userInfo, m)
	RedisClient.Expire(ctx, userInfo, time.Duration(consts.RedisUserInfoExp))
	return nil
}

func GetInfo(ctx context.Context, userId int64) (*user.User, error) {
	userInfo := fmt.Sprintf(consts.RedisUserInfoPre, userId)
	data, err := RedisClient.HGetAll(ctx, userInfo).Result()
	if err != nil {
		return nil, err
	}
	id, _ := strconv.ParseInt(data["id"], 10, 64)
	followCount, _ := strconv.ParseInt(data["follow_count"], 10, 64)
	followerCount, _ := strconv.ParseInt(data["follower_count"], 10, 64)
	isFollow, _ := strconv.ParseBool(data["is_follow"])
	totalFavorited, _ := strconv.ParseInt(data["total_favorited"], 10, 64)
	workCount, _ := strconv.ParseInt(data["work_count"], 10, 64)
	favoriteCount, _ := strconv.ParseInt(data["favorite_count"], 10, 64)
	return &user.User{Id: id, Name: data["name"], FavoriteCount: followCount, FollowerCount: followerCount, IsFollow: isFollow,
		Avatar: data["avatar"], BackgroundImage: data["background_image"], Signature: data["signature"], TotalFavorited: totalFavorited,
		WorkCount: workCount, FollowCount: favoriteCount,
	}, nil
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
