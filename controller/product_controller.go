package controller

import (
	"fmt"
	"go-sqlite-crud-product/model"
	"go-sqlite-crud-product/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService *service.ProductService
}

func NewProductController(productService *service.ProductService) *ProductController {
	return &ProductController{ProductService: productService}
}
func (controller *ProductController) CreateProduct(c *gin.Context) {
	fmt.Println("CreateProduct: Received request to create a product")

	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		fmt.Println("CreateProduct: Invalid JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	fmt.Printf("CreateProduct: Parsed product details: %+v\n", product)

	createdProduct, err := controller.ProductService.CreateProduct(&product)
	if err != nil {
		fmt.Printf("CreateProduct: Error creating product in service layer: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	fmt.Printf("CreateProduct: Successfully created product: %+v\n", createdProduct)
	c.JSON(http.StatusOK, createdProduct)
}

func (controller *ProductController) GetProduct(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	product, err := controller.ProductService.GetProduct(productID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (controller *ProductController) GetAllProducts(c *gin.Context) {
	products, err := controller.ProductService.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (controller *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.ID = productID
	updatedProduct, err := controller.ProductService.UpdateProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

func (controller *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = controller.ProductService.DeleteProduct(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
