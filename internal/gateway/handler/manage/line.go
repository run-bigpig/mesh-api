package manage

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/mesh-api/internal/data/entry/line"
	"github.com/run-bigpig/mesh-api/internal/data/entry/model"
	"github.com/run-bigpig/mesh-api/internal/errorcode"
	"github.com/run-bigpig/mesh-api/internal/gateway/handler"
	"github.com/run-bigpig/mesh-api/internal/gateway/response"
)

// AddLine godoc
// @Summary 添加线路
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

// UpdateLine godoc
// @Summary 更新线路
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

// FindLine godoc
// @Summary 查找线路
func FindLine(ctx *fiber.Ctx) error {
	var req handler.IdRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}
	if req.Id == 0 {
		return errorcode.ErrorCodeInvalidParam
	}
	data, err := line.FindOne(ctx.Context(), req.Id)
	if err != nil {
		return err
	}
	return response.Success(ctx, data)
}

// ListLine godoc
// @Summary 获取线路列表
func ListLine(ctx *fiber.Ctx) error {
	var req line.FindLineRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}
	list, err := line.FindAll(ctx.Context(), &req)
	if err != nil {
		return err
	}
	return response.Success(ctx, list)
}

// DeleteLine godoc
// @Summary 删除线路
func DeleteLine(ctx *fiber.Ctx) error {
	var req handler.IdRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}
	if req.Id == 0 {
		return errorcode.ErrorCodeInvalidParam
	}
	err = line.DeleteOne(ctx.Context(), req.Id)
	if err != nil {
		return err
	}
	return response.Success(ctx, nil)
}

// FindLineByModelId godoc
// @Summary 根据模型id查找线路
func FindLineByModelId(ctx *fiber.Ctx) error {
	var req handler.IdRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}
	if req.Id == 0 {
		return errorcode.ErrorCodeInvalidParam
	}
	lineIds, err := model.FindModelLine(ctx.Context(), req.Id)
	if err != nil {
		return err
	}
	list, err := line.FindLineByIds(ctx.Context(), lineIds)
	if err != nil {
		return err
	}
	return response.Success(ctx, list)
}
