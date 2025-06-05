package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/radityacandra/besart-gallery/api/product"
	"github.com/radityacandra/besart-gallery/pkg/util"
)

func (h *Handler) ProductDetailGet(c echo.Context, productId product.ProductIdPathParams) error {
	if err := uuid.Validate(productId); err != nil {
		return util.ReturnBadRequest(c, err, h.Logger)
	}

	reqCtx := c.Request().Context()
	output, err := h.Service.ProductDetail(reqCtx, productId)
	if err != nil {
		return util.ReturnError(c, err, h.Logger)
	}

	return c.JSON(http.StatusOK, output)
}
