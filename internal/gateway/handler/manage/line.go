package manage

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/mesh-api/internal/data/entry/line"
	"github.com/run-bigpig/mesh-api/internal/errorcode"
	"github.com/run-bigpig/mesh-api/internal/gateway/response"
)

func AddLine(ctx *fiber.Ctx) error {
	var req line.Line
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}
	if req.Name == "" || req.Host == "" || req.Status == 0 || req.Adapter == "" || req.IsProxy == 0 {
		return errorcode.ErrorCodeInvalidParam
	}
	err = line.InsertOne(ctx.Context(), &req)
	if err != nil {
		return err
	}
	return response.Success(ctx, nil)
}

func UpdateLine(ctx *fiber.Ctx) error {
	var req line.Line
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}
	if req.Id == 0 {
		return errorcode.ErrorCodeInvalidParam
	}
	err = line.UpdateOne(ctx.Context(), &req)
	if err != nil {
		return err
	}
	return response.Success(ctx, nil)
}
