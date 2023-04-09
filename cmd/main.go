package main

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"

	configs "CachingDatabase/configs"
	database "CachingDatabase/database"
	public_route "CachingDatabase/route"
)

func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		fmt.Println("Error Configuration :")
		panic(err)
	}
	app := fiber.New()
	redisConnected := database.RedisInit(cfg.REDIS_URL)
	public_route.Public(app, redisConnected)

	app.Listen(":3000")

}
