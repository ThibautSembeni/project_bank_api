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

// @BasePath /api/

func NewPaymentHandler(paymentService payment.Service) *paymentHandler {
	return &paymentHandler{paymentService}
}

//		@Summary		Payment Store
//		@Description	Create Payment
//	 	@Schemes
//		@Tags			payment
//		@Accept			json
//		@Param			create	body		payment.InputPayment	true	"Create payment"
//		@Produce		json
//		@Success		200	{object}	Response
//		@Security		ApiKeyAuth
//		@Router			/api/payment [post]
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

//		@Summary		Payment Update
//		@Description	Update Payment
//	 	@Schemes
//		@Tags			payment
//		@Accept			json
//		@Param			id		path		int			true	"Payment ID"
//		@Param			create	body		payment.InputPayment	true	"Update payment"
//		@Produce		json
//		@Success		200	{object}	Response
//		@Security		ApiKeyAuth
//		@Router			/api/payment/:id/update [put]
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

//		@Summary		Payment Delete
//		@Description	Delete Payment
//	 	@Schemes
//		@Tags			payment
//		@Accept			json
//		@Param			id		path		int			true	"Payment ID"
//		@Produce		json
//		@Success		200	{object}	Response
//		@Security		ApiKeyAuth
//		@Router			/api/payment/:id/delete [delete]
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

//		@Summary		Payment Search
//		@Description	Find Payment
//	 	@Schemes
//		@Tags			payment
//		@Accept			json
//		@Param			id		path		int			true	"Payment ID"
//		@Produce		json
//		@Success		200	{object}	Response
//		@Security		ApiKeyAuth
//		@Router			/api/payment/:id [get]
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

//		@Summary		All Payment Search
//		@Description	Search all Payment
//	 	@Schemes
//		@Tags			payment
//		@Accept			json
//		@Param			id		path		int			true	"Payment ID"
//		@Produce		json
//		@Success		200	{object}	Response
//		@Security		ApiKeyAuth
//		@Router			/api/payment [get]
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
