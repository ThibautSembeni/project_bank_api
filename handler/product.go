package handler

import (
	"net/http"
	"project_api/product"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type productHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) *productHandler{
	return &productHandler{ productService}
}

func (h *productHandler) Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func (h *productHandler) Store(c *gin.Context)  {
	var input product.InputProduct
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := Response{
			Success: false,
			Message: "Verify your data format or structure",
			Data: err.Error(),
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

func (h *productHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input product.InputProduct
	err := c.ShouldBindJSON(&input)
	product, err := h.productService.Update(id, input)
	if err != nil {
		response := Response{
			Success: false,
			Message: "Verify your data format or structure",
			Data: err.Error(),
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

func (h *productHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.productService.Delete(id)
	if err != nil {
		response := Response{
			Success: false,
			Message: "Impossible to delete your product",
			Data: err.Error(),
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