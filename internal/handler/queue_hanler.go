package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rinuccia/transaction-system/internal/models"
	"github.com/rinuccia/transaction-system/internal/service"
	"net/http"
	"strconv"
)

type queueHandler struct {
	service service.Queue
}

func newQueueHandler(s service.Queue) *queueHandler {
	return &queueHandler{
		service: s,
	}
}

func (h *queueHandler) sendRequest(c *gin.Context) {
	request := models.UserRequest{}
	err := c.BindJSON(&request)
	validateErr := validate.Struct(request)
	if err != nil || validateErr != nil {
		c.JSON(http.StatusBadRequest, newErrResponse("invalid input body"))
		return
	}

	if request.AmountOfMoney < 0 {
		c.JSON(http.StatusBadRequest, newErrResponse("invalid input body"))
		return
	}

	num := fmt.Sprintf("%.2f", request.AmountOfMoney)
	result, err := strconv.ParseFloat(num, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, newErrResponse("invalid input body"))
		return
	}

	request.AmountOfMoney = result

	h.service.Handle(request)
}
