package controller

import (
	"honya/backend/dto"
	"honya/backend/errors"
	"honya/backend/service"
	"honya/backend/utils"

	"github.com/gofiber/fiber/v2"
)

type UrlController interface {
	ProcessUrl(ctx *fiber.Ctx) error
}

type urlController struct {
	service service.UrlService
}

func NewUrlController(service service.UrlService) UrlController {
	return &urlController{service}
}

// ProcessUrl godoc
// @Summary Process a URL to get its redirection or canonical form
// @Description Process a given URL to retrieve its redirection URL, canonical URL, or both
// @Tags url
// @Accept json
// @Produce json
// @Param request body dto.ProcessUrlRequest true "Process URL Request"
// @Success 200 {object} dto.ProcessUrlResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /url/process-url [post]
func (c *urlController) ProcessUrl(ctx *fiber.Ctx) error {
	var req dto.ProcessUrlRequest
	if err := ctx.BodyParser(&req); err != nil {
		return errors.NewBadRequestError("Invalid JSON body")
	}

	err := utils.ValidateProcessUrlRequest(&req)
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	switch req.Operation {
	case utils.OpRedirection:
		return handleUrlRedirection(ctx, req.Url, c.service)
	case utils.OpCanonical:
		return handleUrlCanonical(ctx, req.Url, c.service)
	case utils.OpAll:
		return handleUrlAll(ctx, req.Url, c.service)
	default:
		return errors.NewBadRequestError("Invalid operation")
	}
}

func handleUrlRedirection(ctx *fiber.Ctx, url string, service service.UrlService) error {
	redirectionUrl, err := service.GetRedirectionUrl(url)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.ProcessUrlResponse{ProcessedUrl: redirectionUrl})
}

func handleUrlCanonical(ctx *fiber.Ctx, url string, service service.UrlService) error {
	canonicalUrl, err := service.GetCanonicalUrl(url)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.ProcessUrlResponse{ProcessedUrl: canonicalUrl})
}

func handleUrlAll(ctx *fiber.Ctx, url string, service service.UrlService) error {
	redirectionUrl, err := service.GetRedirectionUrl(url)
	if err != nil {
		return err
	}

	canonicalUrl, err := service.GetCanonicalUrl(redirectionUrl)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.ProcessUrlResponse{ProcessedUrl: canonicalUrl})
}
