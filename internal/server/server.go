package server

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/radityacandra/besart-gallery/internal/core"
	"github.com/radityacandra/besart-gallery/pkg/validator"

	"go.uber.org/zap"
)

func InitServer(ctx context.Context, deps *core.Dependency) {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Validator = validator.NewValidator()

	e.Use(middleware.CORS())

	deps.Echo = e

	deps.Logger.Info("Web server ready", zap.Int("port", 8080))
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			deps.Logger.Fatal("Failed to start web server", zap.Error(err))
		}
	}()
}
