package route

import (
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
	redis "github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"

	repository "CachingDatabase/domain/repository"
	service "CachingDatabase/domain/service"
	handlers "CachingDatabase/handlers"
)

func Public(app *fiber.App, redis *redis.Client, mongo *mongo.Client) {
	route := app.Group("/api/v1")
	repo := repository.NewRepo(redis, mongo)
	serv := service.NewService(repo)
	handler := handlers.NewCommand(serv)
	route.Get("/get", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"statusCode": http.StatusOK,
			"Message":    "STATUS OK ! ",
		})
	})
	route.Get("/test", handler.CommandTest)
	route.Post("/set", handler.CommandSetData)
	route.Get("/get/:imei", handler.QueriseData)

}
