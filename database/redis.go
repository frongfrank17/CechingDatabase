package database

import (
	"fmt"

	redis "github.com/redis/go-redis/v9"
)

func RedisInit(uri string) *redis.Client {
	con := redis.NewClient(&redis.Options{
		Addr: uri,
	})
	fmt.Println(con)
	return con

}
