package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rinuccia/transaction-system/internal/models"
	"github.com/rinuccia/transaction-system/internal/service"
	"net/http"
)

type userHandler struct {
	service service.User
}

func newUserHandler(s service.User) *userHandler {
	return &userHandler{
		service: s,
	}
}

func (h *userHandler) getUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := h.service.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, newErrResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *userHandler) createUser(c *gin.Context) {
	user := models.User{}
	err := c.BindJSON(&user)
	validationErr := validate.Struct(user)
	if err != nil || validationErr != nil {
		c.JSON(http.StatusBadRequest, newErrResponse("invalid input body"))
		return
	}

	userID, err := h.service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, newErrResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"userID": userID})
}
