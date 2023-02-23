package redis

import (
	"context"
	"strconv"
)

func ReleaseLock(ctx context.Context, routine int64, resourceKey string) {
	routineMark, _ := RedisClient.Get(ctx, resourceKey).Result()
	if strconv.FormatInt(routine, 10) != routineMark {
		// 其它协程误删lock
		return
	}
	RedisClient.Del(ctx, resourceKey)
}
