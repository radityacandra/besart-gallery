package util

import (
	"errors"
	"net/http"
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/besart-gallery/api"
	jwtType "github.com/radityacandra/besart-gallery/pkg/jwt/types"
	"go.uber.org/zap"
)

var errorMap = map[error]int{}

const unknownError string = "unknown error"

func ReturnError(ctx echo.Context, err error, logger *zap.Logger) error {
	var code int
	var message string
	for registeredErr, registeredErrCode := range errorMap {
		if errors.Is(err, registeredErr) {
			code = registeredErrCode
			message = registeredErr.Error()
		}
	}

	if _, ok := err.(*jwtType.AuthorizationError); ok {
		code = http.StatusUnauthorized
		message = err.Error()
	}

	if code == 0 {
		code = http.StatusInternalServerError
	}

	if logger != nil {
		pc, file, line, _ := runtime.Caller(1)

		if code == http.StatusInternalServerError {
			logger.Error("error occured",
				zap.Error(err),
				zap.Any("invoker", runtime.FuncForPC(pc).Name()),
				zap.Any("file", file), zap.Any("line", line))

			message = unknownError
		} else {
			logger.Warn("responding with client error",
				zap.Error(err),
				zap.Any("invoker", runtime.FuncForPC(pc).Name()),
				zap.Any("file", file), zap.Any("line", line))
		}
	}

	return ctx.JSON(code, api.DefaultErrorResponse{
		Error: message,
	})
}

func ReturnBadRequest(ctx echo.Context, err error, logger *zap.Logger) error {
	pc, file, line, _ := runtime.Caller(1)
	logger.Warn("responding with client error",
		zap.Error(err),
		zap.Any("invoker", runtime.FuncForPC(pc).Name()),
		zap.Any("file", file), zap.Any("line", line))

	return ctx.JSON(http.StatusBadRequest, api.DefaultErrorResponse{
		Error: err.Error(),
	})
}
