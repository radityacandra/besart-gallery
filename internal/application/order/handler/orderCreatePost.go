package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/besart-gallery/api"
	"github.com/radityacandra/besart-gallery/internal/application/order/types"
	"github.com/radityacandra/besart-gallery/pkg/jwt"
	"github.com/radityacandra/besart-gallery/pkg/util"
)

func (h *Handler) OrderCreatePost(ctx echo.Context) error {
	data := ctx.Get(jwt.CONTEXT_KEY).(map[string]interface{})

	var requestBody api.OrderCreatePostRequest
	if err := ctx.Bind(&requestBody); err != nil {
		return util.ReturnBadRequest(ctx, err, h.Logger)
	}

	if err := ctx.Validate(requestBody); err != nil {
		return util.ReturnBadRequest(ctx, err, h.Logger)
	}

	reqCtx := ctx.Request().Context()
	input := types.CreateOrderInput{
		UserId: data["sub"].(string),
		ShippingInfo: types.ShippingInfo{
			FullName: requestBody.Shipping.FullName,
			Address:  requestBody.Shipping.FullAddress,
			PhoneNo:  requestBody.Shipping.PhoneNumber,
			Notes:    requestBody.Shipping.Notes,
		},
	}
	for _, item := range requestBody.OrderItems {
		input.OrderItems = append(input.OrderItems, types.OrderItem{
			ProductId: item.ProductId,
			Quantity:  item.Amount,
		})
	}

	id, err := h.Service.CreateOrder(reqCtx, input)
	if err != nil {
		return util.ReturnError(ctx, err, h.Logger)
	}

	return ctx.JSON(http.StatusOK, api.OrderCreatePostResponse{
		Id: id,
	})
}
