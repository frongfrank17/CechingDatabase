package main

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"

	configs "CachingDatabase/configs"
	database "CachingDatabase/database"
	route "CachingDatabase/route"
)

func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		fmt.Println("Error Configuration :")
		panic(err)
	}
	app := fiber.New()

	mongoDb := database.NewConnected(cfg.MONNGO_URL)
	Connect, err := mongoDb.InitConnected()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(Connect)
	}
	//
	redisConnected := database.RedisInit(cfg.REDIS_URL)
	//public_route.Public(app, redisConnected)
	route.Public(app, redisConnected, Connect)
	app.Listen(":3000")

}
