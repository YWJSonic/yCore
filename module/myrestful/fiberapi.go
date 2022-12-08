package myrestful

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/helmet/v2"
)

func FiberRouter(router *RestfulDriver) {
	router.App.Use(helmet.New())
	api := router.App.Group("/api")
	dataService := api.Group("/DataService")

	// this group middleware
	// dataService.Use()

	dataService.Post("/POST_Request", fiberrequest)
	dataService.Get("/GET_Request", fiberrequest)
}

func fiberrequest(c *fiber.Ctx) error {

	var err error
	var res []byte
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(res)
	} else {
		return c.JSON(res)
	}
}
