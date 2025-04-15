package main

import (
    "order-service/internal/database"
    "order-service/internal/order"
    "github.com/gin-gonic/gin"
)

func main() {
    db := database.ConnectMongo()
    r := gin.Default()

    repo := order.NewRepository(db)
    service := order.NewService(repo)
    handler := order.NewHandler(service)

    api := r.Group("/orders")
    {
        api.POST("/", handler.CreateOrder)
        api.GET("/:id", handler.GetOrderByID)
        api.PATCH("/:id", handler.UpdateOrderStatus)
        api.GET("/", handler.ListOrders)
    }

    r.Run(":8082")
}
