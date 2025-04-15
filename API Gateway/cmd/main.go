package main

import (
    "api-gateway/internal/proxy"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    inventory := r.Group("/products")
    {
        inventory.POST("/", proxy.ProxyToInventory)
        inventory.GET("/", proxy.ProxyToInventory)
        inventory.GET("/:id", proxy.ProxyToInventory)
        inventory.PATCH("/:id", proxy.ProxyToInventory)
        inventory.DELETE("/:id", proxy.ProxyToInventory)
    }

    orders := r.Group("/orders")
    {
        orders.POST("/", proxy.ProxyToOrders)
        orders.GET("/", proxy.ProxyToOrders)
        orders.GET("/:id", proxy.ProxyToOrders)
        orders.PATCH("/:id", proxy.ProxyToOrders)
    }

    r.Run(":8080")
}
