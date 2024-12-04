package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/part_of_the_authentication_service/pkg/model"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type partAuthService interface {
}

type Handler struct {
	service partAuthService
}

func Newhandler(service partAuthService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	return r
}

func handlerError(c *gin.Context, err error, message string, statusCode int) {
	errorResponse := &model.Error{
		Code:    statusCode,
		Message: message,
	}

	log.WithFields(log.Fields{
		"error":   err.Error(),
		"context": c.Request.URL.Path,
	}).Error(message)

	c.JSON(statusCode, errorResponse)
}

func sendResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	response := model.Response{
		Status:  http.StatusText(statusCode),
		Message: message,
		Data:    data,
	}
	c.JSON(statusCode, response)
}
