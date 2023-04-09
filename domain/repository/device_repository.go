package repository

import (
	"context"
	"fmt"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type redisDeviceCheching struct {
	redis *redis.Client
}

var ctx = context.Background()

func NewRepo(redis *redis.Client) redisDeviceCheching {
	return redisDeviceCheching{redis: redis}
}

func (rdb redisDeviceCheching) SetData(devicename string, imei string) error {
	redis_key := imei
	fmt.Println("Repository : ", imei, " Values : ", devicename)
	err := rdb.redis.SetNX(ctx, redis_key, devicename, 300*time.Second).Err()
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}
func (rdb redisDeviceCheching) GetData(imei string) (string, error) {
	redis_key := imei
	result, err := rdb.redis.Get(ctx, redis_key).Result()
	if err != nil {
		return "", err
	}
	return result, nil

}
