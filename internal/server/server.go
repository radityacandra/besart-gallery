package server

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/radityacandra/besart-gallery/api/order"
	"github.com/radityacandra/besart-gallery/api/product"
	orderHandler "github.com/radityacandra/besart-gallery/internal/application/order/handler"
	"github.com/radityacandra/besart-gallery/internal/application/product/handler"
	"github.com/radityacandra/besart-gallery/internal/core"
	"github.com/radityacandra/besart-gallery/pkg/jwt"
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

	ePrivate := e.Group("")
	ePrivate.Use(jwt.Authorize())

	handler := handler.NewHandler(deps)
	product.RegisterHandlers(e, handler)

	orderHandler := orderHandler.NewHandler(deps)
	order.RegisterHandlers(ePrivate, orderHandler)

	deps.Logger.Info("Web server ready", zap.Int("port", 9000))
	go func() {
		if err := e.Start(":9000"); err != nil && err != http.ErrServerClosed {
			deps.Logger.Fatal("Failed to start web server", zap.Error(err))
		}
	}()
}
