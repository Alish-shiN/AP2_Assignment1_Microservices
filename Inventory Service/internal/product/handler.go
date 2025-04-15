package product

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

func (h *Handler) CreateProduct(c *gin.Context) {
    var p Product
    if err := c.ShouldBindJSON(&p); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    id, err := h.service.Create(context.TODO(), p)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"id": id.Hex()})
}

func (h *Handler) GetProductByID(c *gin.Context) {
    id, _ := primitive.ObjectIDFromHex(c.Param("id"))
    p, err := h.service.GetByID(context.TODO(), id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }
    c.JSON(http.StatusOK, p)
}

func (h *Handler) UpdateProduct(c *gin.Context) {
    id, _ := primitive.ObjectIDFromHex(c.Param("id"))
    var p Product
    if err := c.ShouldBindJSON(&p); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := h.service.Update(context.TODO(), id, p)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
    id, _ := primitive.ObjectIDFromHex(c.Param("id"))
    err := h.service.Delete(context.TODO(), id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}

func (h *Handler) ListProducts(c *gin.Context) {
    products, err := h.service.List(context.TODO())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, products)
}
