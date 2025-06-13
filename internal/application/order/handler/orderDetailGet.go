package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/besart-gallery/api"
	"github.com/radityacandra/besart-gallery/api/order"
	"github.com/radityacandra/besart-gallery/pkg/jwt"
	"github.com/radityacandra/besart-gallery/pkg/util"
)

func (h *Handler) OrderDetailGet(ctx echo.Context, orderId order.OrderIdPathParams) error {
	data := ctx.Get(jwt.CONTEXT_KEY).(map[string]interface{})
	userId := data["sub"].(string)

	reqCtx := ctx.Request().Context()
	output, err := h.Service.DetailOrder(reqCtx, userId, orderId)
	if err != nil {
		return util.ReturnError(ctx, err, h.Logger)
	}

	response := api.OrderDetailGetResponse{
		Id:        output.Id,
		OrderTime: output.OrderTime,
		Status:    output.Status,
		Shipping: api.ShippingAddressRequest{
			FullName:    output.ShippingInfo.FullName,
			FullAddress: output.ShippingInfo.Address,
			PhoneNumber: output.ShippingInfo.PhoneNo,
			Notes:       output.ShippingInfo.Notes,
		},
		OrderItems: []api.OrderItemResponse{},
	}

	for _, item := range output.OrderItems {
		response.OrderItems = append(response.OrderItems, api.OrderItemResponse{
			Id:           item.Id,
			ProductName:  item.ProductName,
			ProductImage: item.ProductImage,
			Qty:          item.Quantity,
			ProductPrice: item.ProductPrice,
		})
	}

	return ctx.JSON(http.StatusOK, response)
}
