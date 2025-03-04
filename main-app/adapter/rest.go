package adapter

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/revandpratama/reflect/config"
	"github.com/revandpratama/reflect/helper"
	"github.com/revandpratama/reflect/internal/routes"
)

type RestOption struct {
	app *fiber.App
}

func (r *RestOption) Start(a *Adapter) error {
	helper.NewLog().Info("initializing REST server...").ToKafka()

	r.app = fiber.New()

	r.app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong"})
	})

	api := r.app.Group("/api")

	postHandler := routes.InitPostHandler(a.Postgres, a.MinioClient)
	routes.InitPostRoutes(api, postHandler)

	authHandler := routes.InitAuthHandler(a.GrcpClient)
	routes.InitAuthRoutes(api, authHandler)

	go func() {
		if err := r.app.Listen(fmt.Sprintf(":%v", config.ENV.RESTServerPort)); err != nil {
			helper.NewLog().Fatal(fmt.Sprintf("Failed to start REST server: %v", err)).ToKafka()
			os.Exit(1)
		}
	}()

	helper.NewLog().Info("REST server started").ToKafka()

	return nil
}

func (r *RestOption) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return r.app.ShutdownWithContext(ctx)
}
