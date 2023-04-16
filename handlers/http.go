package handlers

import (
	"fmt"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"

	service "CachingDatabase/domain/service"
)

type HttpCommandHandler struct {
	serv service.DeviceInterface
}
type HttpCommandHandlerInterface interface {
	CommandSetData(*fiber.Ctx) error
	CommandTest(*fiber.Ctx) error
	QueriseData(*fiber.Ctx) error
}

func NewCommand(serv service.DeviceInterface) HttpCommandHandlerInterface {
	return HttpCommandHandler{serv: serv}
}
func (svc HttpCommandHandler) CommandTest(c *fiber.Ctx) error {
	data, _ := svc.serv.Test()
	fmt.Println(data)
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"statusCode": http.StatusCreated,
		"message":    "Created Data In Chaching",
	})
}
func (svc HttpCommandHandler) CommandSetData(c *fiber.Ctx) error {
	type RequestCommandSetData struct {
		Devicename string `json:"devicename"`
		Imei       string `json:"imei"`
	}
	requestBody := new(RequestCommandSetData)
	errBadRequest := c.BodyParser(&requestBody)
	fmt.Println(requestBody)
	if errBadRequest != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": http.StatusBadRequest,
			"message":    " Bad Request Values ",
		})
	}
	errSerivceInsert := svc.serv.InsertData(requestBody.Devicename, requestBody.Imei)
	if errSerivceInsert != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"statusCode": http.StatusInternalServerError,
			"message":    errSerivceInsert.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"statusCode": http.StatusCreated,
		"message":    "Created Data In Chaching",
	})
}
func (svc HttpCommandHandler) QueriseData(c *fiber.Ctx) error {
	imei := c.Params("imei")

	data, _ := svc.serv.GetData(imei)
	fmt.Println(data)
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"statusCode": http.StatusCreated,
		"message":    "Get Message ",
		"data":       data,
	})
}
