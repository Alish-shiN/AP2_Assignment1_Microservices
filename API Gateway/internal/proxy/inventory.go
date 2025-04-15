package proxy

import (
    "github.com/gin-gonic/gin"
)

var inventoryURL = "http://localhost:8081/products"

func ProxyToInventory(c *gin.Context) {
    forwardRequest(c, inventoryURL)
}
