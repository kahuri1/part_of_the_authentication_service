package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/part_of_the_authentication_service/pkg/logger"
)

type dataResponse struct {
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}

type idResponse struct {
	ID interface{} `json:"id"`
}

type response struct {
	Message string `json:"message"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	logger.Error(message)
	c.AbortWithStatusJSON(statusCode, response{message})
}
