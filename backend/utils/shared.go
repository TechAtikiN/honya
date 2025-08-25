package utils

import (
	"honya/backend/errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func ParseInt(val string, defaultValue int) int {
	if v, err := strconv.Atoi(val); err == nil {
		return v
	}
	return defaultValue
}

func ParseFloat(val string, defaultValue float64) float64 {
	if v, err := strconv.ParseFloat(val, 64); err == nil {
		return v
	}
	return defaultValue
}

func ParseUUIDParam(ctx *fiber.Ctx, param string) (uuid.UUID, error) {
	idStr := ctx.Params(param)
	if idStr == "" {
		return uuid.Nil, errors.NewBadRequestError("ID is required")
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.Nil, errors.NewBadRequestError("Invalid ID")
	}
	return id, nil
}
