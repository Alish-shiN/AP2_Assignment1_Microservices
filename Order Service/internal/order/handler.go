package order

import (
    "context"
    "net/http"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
    service Service
}

func NewHandler(s Service) *Handler {
    return &Handler{service: s}
}

func (h *Handler) CreateOrder(c *gin.Context) {
    var o Order
    if err := c.ShouldBindJSON(&o); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    id, err := h.service.Create(context.TODO(), o)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"id": id.Hex()})
}

func (h *Handler) GetOrderByID(c *gin.Context) {
    id, _ := primitive.ObjectIDFromHex(c.Param("id"))
    o, err := h.service.GetByID(context.TODO(), id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }
    c.JSON(http.StatusOK, o)
}

func (h *Handler) UpdateOrderStatus(c *gin.Context) {
    id, _ := primitive.ObjectIDFromHex(c.Param("id"))
    var payload struct {
        Status string `json:"status"`
    }
    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := h.service.UpdateStatus(context.TODO(), id, payload.Status)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}

func (h *Handler) ListOrders(c *gin.Context) {
    orders, err := h.service.List(context.TODO())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, orders)
}
