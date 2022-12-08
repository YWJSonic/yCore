package myrestful

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
)

type RestfulDriver struct {
	*gin.Engine
	*fiber.App
}

func (r *RestfulDriver) Listen(addr string) error {
	if r.Engine != nil {
		return r.Run(addr)
	} else if r.App != nil {
		return r.App.Listen(addr)
	}
	return errors.New("")
}

// Gin Restful
func GinNew() *RestfulDriver {
	driverObj := &RestfulDriver{
		Engine: gin.Default(),
	}
	return driverObj
}

// Fiber Restful
func FiberNew() *RestfulDriver {
	api := fiber.New()
	api.Use(helmet.New(), cors.New(), logger.New())
	driverObj := &RestfulDriver{
		App: api,
	}
	return driverObj
}
