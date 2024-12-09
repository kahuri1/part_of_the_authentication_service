package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/part_of_the_authentication_service/iternal/domain"
	"github.com/kahuri1/part_of_the_authentication_service/iternal/model"
	"net/http"
)

func (h *Handler) UserSingUpInput(c *gin.Context) {
	var input model.UserSignUpInput

	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")

		return
	}

	if err := h.service.SignUpService(c.Request.Context(), input); err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			newResponse(c, http.StatusBadRequest, err.Error())
		}
		newResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.Status(http.StatusCreated)
}
