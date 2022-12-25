package handler

import (
	"fmt"
	"net/http"
	"project_api/adapter"
	"project_api/payment"
	service "project_api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type paymentHandler struct {
	paymentService payment.Service
}

func NewPaymentHandler(paymentService payment.Service) *paymentHandler {
	return &paymentHandler{paymentService}
}

func (h *paymentHandler) Store(c *gin.Context) {
	var input payment.InputPayment
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := Response{
			Success: false,
			Message: "Verify your data format or structure",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newPayment, err := h.paymentService.Store(input)
	if err != nil {
		response := Response{
			Success: false,
			Message: "something went wrong",
			Data:    err.Error(),
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}

	roomManager := service.NewRoomManager()
	adapterRest := adapter.NewGinAdapter(roomManager)
	adapterRest.Post(fmt.Sprintf("Un payement de %v a été réalisé pour le produit  %v", newPayment.PricePaid, newPayment.Product.Name))
	response := Response{
		Success: true,
		Message: "new payment successfully register",
		Data:    newPayment,
	}

	c.JSON(http.StatusOK, response)
}

func (h *paymentHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input payment.InputPayment
	err := c.ShouldBindJSON(&input)
	payment, err := h.paymentService.Update(id, input)
	if err != nil {
		response := Response{
			Success: false,
			Message: "Verify your data format or structure",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	roomManager := service.NewRoomManager()
	adapterRest := adapter.NewGinAdapter(roomManager)
	adapterRest.Post(fmt.Sprintf("Le prix de la transaction #%v à été modifié ! Nouveau montant : %v", payment.ID, payment.PricePaid))
	response := Response{
		Success: true,
		Message: "Payment successfully updated",
		Data:    payment,
	}

	c.JSON(http.StatusOK, response)
}

func (h *paymentHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.paymentService.Delete(id)
	if err != nil {
		response := Response{
			Success: false,
			Message: "Impossible to delete your product",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := Response{
		Success: true,
		Message: "Product successfully deleted",
	}

	c.JSON(http.StatusOK, response)
}

func (h *paymentHandler) FetchById(c *gin.Context) {
	id := c.Param("id")
	payment, err := h.paymentService.Find(id)
	if err != nil {
		response := Response{
			Success: false,
			Message: "something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := Response{
		Success: true,
		Data:    payment,
	}

	c.JSON(http.StatusOK, response)
}

func (handler *paymentHandler) List(c *gin.Context) {

	products, err := handler.paymentService.ListAll()
	if err != nil {
		response := Response{
			Success: false,
			Message: "something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := Response{
		Success: true,
		Data:    products,
	}

	c.JSON(http.StatusOK, response)
}
