package restful

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/helmet/v2"
)

func FiberRouter(router *RestfulDriver) {
	router.App.Use(helmet.New())
	api := router.App.Group("/api")
	dataService := api.Group("/DataService")

	// this group middleware
	// dataService.Use()

	dataService.Post("/Token", fiberToken)
	dataService.Get("/MatchById", fiberMatchById)
	dataService.Get("/CompetitionById", fiberCompetitionById)
	dataService.Get("/Region", fiberRegion)
}

func fiberToken(c *fiber.Ctx) error {
	// req := adaptor.TokenRequestDto{}

	// if res, err := adaptor.Token(req); err != nil {
	// 	return c.Status(http.StatusBadRequest).JSON(res)
	// } else {
	// 	return c.JSON(res)
	// }

	return c.JSON("")
}

func fiberMatchById(c *fiber.Ctx) error {
	// req := adaptor.MatchByIdRequestDto{}

	// err := c.QueryParser(&req)
	// if err != nil {
	// 	return c.SendStatus(http.StatusBadRequest)
	// }

	// if res, err := adaptor.MatchById(req); err != nil {
	// 	return c.Status(http.StatusBadRequest).JSON(res)
	// } else {
	// 	return c.JSON(res)
	// }

	return c.JSON("")
}

func fiberCompetitionById(c *fiber.Ctx) error {
	// req := adaptor.CompetitionByIdRequestDto{}

	// err := c.QueryParser(&req)
	// if err != nil {
	// 	return c.SendStatus(http.StatusBadRequest)
	// }

	// if res, err := adaptor.CompetitionById(req); err != nil {
	// 	return c.Status(http.StatusBadRequest).JSON(res)
	// } else {
	// return c.JSON(res)
	// }

	return c.JSON("")
}

func fiberRegion(c *fiber.Ctx) error {
	// req := adaptor.RegionRequestDto{}

	// err := c.QueryParser(&req)
	// if err != nil {
	// 	return c.SendStatus(http.StatusBadRequest)
	// }

	// if res, err := adaptor.Region(req); err != nil {
	// 	return c.Status(http.StatusBadRequest).JSON(res)
	// } else {
	// 	return c.JSON(res)
	// }

	return c.JSON("")
}
