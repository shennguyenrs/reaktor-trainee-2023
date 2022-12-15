package router

import (
	"log"

	conf "reaktor_birdnest/src/config"
	ctrl "reaktor_birdnest/src/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
)

func StartServer() {
	// Load env
	conf.LoadEnv()

	// Start router
	r := fiber.New(fiber.Config{
		AppName:       "Birdnest API v1",
		CaseSensitive: true,
		StrictRouting: true,
	})

	// Middlewares
	r.Use(helmet.New())
	r.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	r.Use(recover.New())
	r.Use(compress.New())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowHeaders:     "*",
		AllowMethods:     "GET, POST",
		MaxAge:           300,
	}))

	// Routes
	reaktor := r.Group("/violations")
	reaktor.Get("/temp", ctrl.GetTemp)
	reaktor.Post("/temp/check", ctrl.Check)
	reaktor.Get("/persist", ctrl.GetPersistList)
	reaktor.Post("/persist", ctrl.UpdatePersistList)

	// Start server
	log.Fatal(r.Listen(":3001"))
}
