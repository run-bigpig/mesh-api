package manage

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/mesh-api/internal/data/entry/model"
	"github.com/run-bigpig/mesh-api/internal/errorcode"
	"github.com/run-bigpig/mesh-api/internal/gateway/handler"
	"github.com/run-bigpig/mesh-api/internal/gateway/response"
)

func AddModel(ctx *fiber.Ctx) error {
	var req model.Model
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}
	if req.Name == "" || req.Class == 0 || req.Status == 0 {
		return errorcode.ErrorCodeInvalidParam
	}
	err = model.InsertOne(ctx.Context(), &model.Model{
		Name:   req.Name,
		Class:  req.Class,
		Status: req.Status,
	})
	if err != nil {
		return err
	}
	return response.Success(ctx, nil)
}

func DeleteModel(ctx *fiber.Ctx) error {
	var req handler.IdRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}
	if req.Id == 0 {
		return errorcode.ErrorCodeInvalidParam
	}
	err = model.DeleteOne(ctx.Context(), req.Id)
	if err != nil {
		return err
	}
	return response.Success(ctx, nil)
}

func UpdateModel(ctx *fiber.Ctx) error {
	var req model.Model
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}
	if req.Name == "" || req.Class == 0 || req.Status == 0 {
		return errorcode.ErrorCodeInvalidParam
	}
	err = model.UpdateOne(ctx.Context(), &req)
	if err != nil {
		return err
	}
	return response.Success(ctx, nil)
}

func FindModel(ctx *fiber.Ctx) error {
	var req handler.IdRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}
	if req.Id == 0 {
		return errorcode.ErrorCodeInvalidParam
	}
	data, err := model.FindOne(ctx.Context(), req.Id)
	if err != nil {
		return err
	}
	return response.Success(ctx, data)
}

func ListModel(ctx *fiber.Ctx) error {
	var req model.FindModelRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}
	data, err := model.FindAll(ctx.Context(), &model.FindModelRequest{
		Name:   req.Name,
		Class:  req.Class,
		Status: req.Status,
	})
	if err != nil {
		return err
	}
	return response.Success(ctx, data)
}
