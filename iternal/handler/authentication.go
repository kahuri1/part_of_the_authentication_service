package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/part_of_the_authentication_service/iternal/model"
	"net/http"
)

func (h *Handler) AuthenticationHandler(c *gin.Context) {
	var auth model.AuthenticationRequest
	//TODO Добавить ip в jwt token
	d, err := c.GetRawData()
	err = json.Unmarshal(d, &auth)
	if err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	auth.Ip = c.ClientIP()
	//if err := c.BindJSON(&auth.Uuid); err != nil {
	//	newResponse(c, http.StatusBadRequest, "invalid input body")
	//
	//	return
	//}

	tokens, err := h.service.Authentication(auth)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"tokens": tokens})
}

func (h *Handler) RefreshToken(c *gin.Context) {
	var token model.Tokens
	d, err := c.GetRawData()
	err = json.Unmarshal(d, &token)
	if err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	ip := c.ClientIP()
	tokens, err := h.service.RefreshToken(token, ip)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"tokens": tokens})
}
