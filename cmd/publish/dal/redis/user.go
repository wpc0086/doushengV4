package redis

import (
	"context"
	"doushengV4/pkg/consts"
	"fmt"
)

func UpdateInfo(ctx context.Context, userId int64) (err error) {
	userInfo := fmt.Sprintf(consts.RedisUserInfoPre, userId)
	RedisClient.HIncrBy(ctx, userInfo, "work_count", 1)
	return nil
}

func GetInfo(ctx context.Context, userId int64) (bool, error) {
	userInfo := fmt.Sprintf(consts.RedisUserInfoPre, userId)
	data, err := RedisClient.HGetAll(ctx, userInfo).Result()
	if err != nil {
		return false, err
	}
	if data == nil || len(data) == 0 {
		return false, nil
	}
	return true, nil
}
