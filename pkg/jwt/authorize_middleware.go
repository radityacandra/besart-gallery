package jwt

import (
	"github.com/labstack/echo/v4"
	"github.com/radityacandra/besart-gallery/pkg/util"
)

func Authorize() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			data, err := AuthorizeToken(ctx, c.Request().Header.Get("Authorization"))
			if err != nil {
				return util.ReturnError(c, err, nil)
			}

			c.Set(CONTEXT_KEY, data)

			return next(c)
		}
	}
}
