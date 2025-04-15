package main

import (
    "inventory-service/internal/database"
    "inventory-service/internal/product"
    "github.com/gin-gonic/gin"
)

func main() {
    db := database.ConnectMongo()
    r := gin.Default()

    productRepo := product.NewRepository(db)
    productService := product.NewService(productRepo)
    productHandler := product.NewHandler(productService)

    api := r.Group("/products")
    {
        api.POST("/", productHandler.CreateProduct)
        api.GET("/:id", productHandler.GetProductByID)
        api.PATCH("/:id", productHandler.UpdateProduct)
        api.DELETE("/:id", productHandler.DeleteProduct)
        api.GET("/", productHandler.ListProducts)
    }

    r.Run(":8081")
}
