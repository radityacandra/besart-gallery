package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/besart-gallery/api"
	"github.com/radityacandra/besart-gallery/internal/application/order/types"
	"github.com/radityacandra/besart-gallery/pkg/jwt"
	"github.com/radityacandra/besart-gallery/pkg/util"
)

func (h *Handler) OrderListGet(ctx echo.Context) error {
	data := ctx.Get(jwt.CONTEXT_KEY).(map[string]interface{})
	userId := data["sub"].(string)

	reqCtx := ctx.Request().Context()
	output, err := h.Service.ListOrder(reqCtx, types.ListOrderInput{
		UserId:   userId,
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		return util.ReturnError(ctx, err, h.Logger)
	}

	items := []api.OrderListGetDetail{}
	for _, item := range output.Data {
		items = append(items, api.OrderListGetDetail{
			Id:          item.Id,
			Status:      item.Status,
			TotalAmount: item.TotalAmount,
		})
	}

	response := api.OrderListGetResponse{
		Data: items,
		Pagination: api.PaginationSchema{
			Page:      output.Pagination.Page,
			PageSize:  output.Pagination.PageSize,
			TotalData: output.Pagination.TotalData,
		},
	}

	return ctx.JSON(http.StatusOK, response)
}
