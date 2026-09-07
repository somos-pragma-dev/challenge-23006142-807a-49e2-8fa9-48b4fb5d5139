package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"src/services"
	"src/utils"
)

type ProductController struct {
	productService *services.ProductService
}

func NewProductController(productService *services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var productDTO services.CreateProductDTO
	if err := c.ShouldBindJSON(&productDTO); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := utils.ValidateProductDTO(productDTO); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := pc.productService.CreateProduct(productDTO)
	if err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (pc *ProductController) GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := pc.productService.GetProduct(id)
	if err!= nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var productDTO services.UpdateProductDTO
	if err := c.ShouldBindJSON(&productDTO); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := utils.ValidateProductDTO(productDTO); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := pc.productService.UpdateProduct(id, productDTO)
	if err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := pc.productService.DeleteProduct(id); err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}