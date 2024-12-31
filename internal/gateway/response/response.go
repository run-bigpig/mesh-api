package response

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/mesh-api/internal/errorcode"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Fail(ctx *fiber.Ctx, err error) error {
	var e *errorcode.ErrorCode
	if errors.As(err, &e) {
		return ctx.Status(500).JSON(&Response{
			Code:    e.GetCode(),
			Message: e.Error(),
			Data:    nil,
		})
	}
	if errors.Is(err, sql.ErrNoRows) {
		err = errors.New("record not found")
	}
	return ctx.Status(500).JSON(&Response{
		Code:    500,
		Message: err.Error(),
		Data:    nil,
	})
}

func Success(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(200).JSON(&Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}
