package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/radityacandra/besart-gallery/api"
	"github.com/radityacandra/besart-gallery/api/order"
	"github.com/radityacandra/besart-gallery/internal/application/order/types"
	"github.com/radityacandra/besart-gallery/pkg/jwt"
	"github.com/radityacandra/besart-gallery/pkg/util"
)

func (h *Handler) OrderStatusPut(ctx echo.Context, orderId order.OrderIdPathParams) error {
	data := ctx.Get(jwt.CONTEXT_KEY).(map[string]interface{})

	if err := uuid.Validate(orderId); err != nil {
		return util.ReturnBadRequest(ctx, err, h.Logger)
	}

	var reqBody api.OrderStatusPutRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return util.ReturnBadRequest(ctx, err, h.Logger)
	}

	if err := ctx.Validate(reqBody); err != nil {
		return util.ReturnBadRequest(ctx, err, h.Logger)
	}

	reqCtx := ctx.Request().Context()
	err := h.Service.UpdateStatus(reqCtx, types.UpdateStatusInput{
		Status:  reqBody.Status,
		OrderId: orderId,
		UserId:  data["sub"].(string),
	})
	if err != nil {
		return util.ReturnError(ctx, err, h.Logger)
	}

	return ctx.JSON(http.StatusOK, api.OrderCreatePostResponse{
		Id: orderId,
	})
}
