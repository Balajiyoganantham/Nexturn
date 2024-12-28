package main

import (
	db "go-sqlite-crud-product/config"
	"go-sqlite-crud-product/controller"
	"go-sqlite-crud-product/middleware"
	"go-sqlite-crud-product/repository"
	"go-sqlite-crud-product/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection for products
	db.InitializeDatabase()
	db.InitializeDatabase1()
	db.InitializeDatabase_auth()

	// Create repository, service, and controller for products
	productRepo := repository.NewProductRepository(db.GetDB1())
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)
	userRepo := repository.NewUserRepository(db.GetDB())
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Initialize Gin router
	r := gin.Default()

	// Apply logging middleware globally
	r.Use(middleware.LoggingMiddleware())

	// Group routes and apply authentication middleware
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware(db.GetDB2()))

	// Routes for users
	api.POST("/users", userController.CreateUser)
	api.GET("/users/:id", userController.GetUser)
	api.GET("/users", userController.GetAllUsers)
	api.PUT("/users/:id", userController.UpdateUser)
	api.DELETE("/users/:id", userController.DeleteUser)

	// Routes for products
	api.POST("/products", productController.CreateProduct)
	api.GET("/products/:id", productController.GetProduct)
	api.GET("/products", productController.GetAllProducts)
	api.PUT("/products/:id", productController.UpdateProduct)
	api.DELETE("/products/:id", productController.DeleteProduct)

	// Start server on port 8080
	r.Run(":8080")
}
