package manage

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/mesh-api/internal/data/entry/model"
	"github.com/run-bigpig/mesh-api/internal/errorcode"
	"github.com/run-bigpig/mesh-api/internal/gateway/handler"
	"github.com/run-bigpig/mesh-api/internal/gateway/response"
)

// AddModel godoc
// @Summary 添加模型
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

// DeleteModel godoc
// @Summary 删除模型
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

// UpdateModel godoc
// @Summary 更新模型
func UpdateModel(ctx *fiber.Ctx) error {
	var req model.Model
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}
	if req.Name == "" || req.Class == 0 || req.Status == 0 || req.Id == 0 {
		return errorcode.ErrorCodeInvalidParam
	}
	err = model.UpdateOne(ctx.Context(), &req)
	if err != nil {
		return err
	}
	return response.Success(ctx, nil)
}

// FindModel godoc
// @Summary 查找模型
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

// ListModel godoc
// @Summary 列出模型
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

// SetModelLine godoc
// @Summary 设置模型线路
func SetModelLine(ctx *fiber.Ctx) error {
	var req model.SetModelLineRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}
	if len(req.ModelIds) == 0 || len(req.LineIds) == 0 {
		return errorcode.ErrorCodeInvalidParam
	}
	err = model.SetModelLine(ctx.Context(), &req)
	if err != nil {
		return err
	}
	return response.Success(ctx, nil)
}
