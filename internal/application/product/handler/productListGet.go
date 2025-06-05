package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/besart-gallery/api"
	"github.com/radityacandra/besart-gallery/api/product"
	"github.com/radityacandra/besart-gallery/internal/application/product/types"
	"github.com/radityacandra/besart-gallery/pkg/util"
)

func (h *Handler) ProductListGet(c echo.Context, params product.ProductListGetParams) error {
	reqCtx := c.Request().Context()
	if err := c.Validate(params); err != nil {
		return util.ReturnBadRequest(c, err, h.Logger)
	}

	page := 1
	if params.Page != nil {
		page = *params.Page
	}

	output, err := h.Service.ProductList(reqCtx, types.ProductListInput{
		Page:     page,
		PageSize: 10,
	})
	if err != nil {
		return util.ReturnError(c, err, h.Logger)
	}

	response := api.ProductListGetResponse{}
	for _, item := range output {
		rating := 0
		if item.Rating != nil {
			rating = int(*item.Rating)
		}

		discountedPrice := 0
		if item.DiscountedPrice != nil {
			discountedPrice = int(*item.DiscountedPrice)
		}

		response = append(response, api.ProductListGetResponseItem{
			Id:              item.Id,
			Name:            item.Name,
			Image:           item.Image,
			Rating:          rating,
			DiscountedPrice: discountedPrice,
			OriginalPrice:   int(item.OriginalPrice),
		})
	}

	return c.JSON(http.StatusOK, response)
}
