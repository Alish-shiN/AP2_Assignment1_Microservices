package proxy

import (
    "github.com/gin-gonic/gin"
)

var orderURL = "http://localhost:8082/orders"

func ProxyToOrders(c *gin.Context) {
    forwardRequest(c, orderURL)
}
