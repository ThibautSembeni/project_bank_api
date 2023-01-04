package handler

import (
	"net/http"
	"project_api/product"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type productHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) *productHandler {
	return &productHandler{productService}
}

// @BasePath /api/

func (h *productHandler) Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

//		@Summary		Product Store
//		@Description	Create Product
//	 	@Schemes
//		@Tags			product
//		@Accept			json
//		@Param			create	body		product.InputProduct	true	"Create product"
//		@Produce		json
//		@Success		200	{object}	Response
//		@Security		ApiKeyAuth
//		@Router			/api/product [post]
func (h *productHandler) Store(c *gin.Context) {
	var input product.InputProduct
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

	newProduct, err := h.productService.Store(input)
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
		Message: "new product successfully added",
		Data:    newProduct,
	}

	c.JSON(http.StatusOK, response)
}

//		@Summary		Product Update
//		@Description	Update Product
//	 	@Schemes
//		@Tags			product
//		@Accept			json
//		@Param			id		path		int			true	"Product ID"
//		@Param			create	body		product.InputProduct	true	"Update product"
//		@Produce		json
//		@Success		200	{object}	Response
//		@Security		ApiKeyAuth
//		@Router			/api/product/:id/update [put]
func (h *productHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input product.InputProduct
	err := c.ShouldBindJSON(&input)
	product, err := h.productService.Update(id, input)
	if err != nil {
		response := Response{
			Success: false,
			Message: "Verify your data format or structure",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := Response{
		Success: true,
		Message: "Product successfully updated",
		Data:    product,
	}

	c.JSON(http.StatusOK, response)
}

//		@Summary		Product Delete
//		@Description	Delete Product
//	 	@Schemes
//		@Tags			product
//		@Accept			json
//		@Param			id		path		int			true	"Product ID"
//		@Produce		json
//		@Success		200	{object}	Response
//		@Security		ApiKeyAuth
//		@Router			/api/product/:id/delete [delete]
func (h *productHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.productService.Delete(id)
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

//		@Summary		Product Search
//		@Description	Find Product
//	 	@Schemes
//		@Tags			product
//		@Accept			json
//		@Param			id		path		int			true	"Product ID"
//		@Produce		json
//		@Success		200	{object}	Response
//		@Security		ApiKeyAuth
//		@Router			/api/product/:id [get]
func (h *productHandler) FetchById(c *gin.Context) {
	id := c.Param("id")
	product, err := h.productService.Find(id)
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
		Data:    product,
	}

	c.JSON(http.StatusOK, response)
}

//		@Summary		All Products Search
//		@Description	Search all Products
//	 	@Schemes
//		@Tags			product
//		@Accept			json
//		@Param			id		path		int			true	"Product ID"
//		@Produce		json
//		@Success		200	{object}	Response
//		@Security		ApiKeyAuth
//		@Router			/api/products [get]
func (handler *productHandler) List(c *gin.Context) {

	products, err := handler.productService.ListAll()
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
