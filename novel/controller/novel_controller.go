package controller

import (
	"module-crud/novel/domain"
	"module-crud/novel/domain/model"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type NovelController struct {
	useCase domain.NovelUseCase
}

func NewNovelController(useCase domain.NovelUseCase) *NovelController {
	return &NovelController{useCase: useCase}
}

func (c *NovelController) CreateNovel(ctx *fiber.Ctx) error {
	var novel model.Novel
	if err := ctx.BodyParser(&novel); err != nil {
		return sendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
	}

	if novel.Author == "" || novel.Name == "" || novel.Description == "" {
		return sendErrorResponse(ctx, http.StatusBadRequest, "Missing required novel fields", nil)
	}

	if err := c.useCase.CreateNovel(novel); err != nil {
		return sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to create novel", err)
	}

	return sendSuccessResponse(ctx, http.StatusCreated, "Novel created successfully", nil)
}

func (c *NovelController) GetNovelById(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return sendErrorResponse(ctx, http.StatusBadRequest, "Invalid ID parameter", err)
	}

	novel, err := c.useCase.GetNovelById(id)
	if err != nil {
		return sendErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve novel", err)
	}

	return sendSuccessResponse(ctx, http.StatusOK, "Novel retrieved successfully", novel)
}

func sendErrorResponse(ctx *fiber.Ctx, statusCode int, message string, err error) error {
	resp := model.Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       nil,
	}
	if err != nil {
		resp.Data = err.Error() // Optionally include the error message in the response
	}
	return ctx.Status(statusCode).JSON(resp)
}

func sendSuccessResponse(ctx *fiber.Ctx, statusCode int, message string, data interface{}) error {
	resp := model.Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
	return ctx.Status(statusCode).JSON(resp)
}
