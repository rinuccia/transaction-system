package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rinuccia/transaction-system/internal/service"
)

const userURL = "/user"

var validate = validator.New()

type Handler struct {
	*userHandler
	*queueHandler
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		newUserHandler(service.User),
		newQueueHandler(service.Queue),
	}
}

func (h *Handler) InitRoutes(router *gin.Engine) {
	router.GET(userURL+"/:id", h.getUserByID)
	router.POST(userURL+"/new", h.createUser)
	router.PUT(userURL+"/request", h.sendRequest)
}
