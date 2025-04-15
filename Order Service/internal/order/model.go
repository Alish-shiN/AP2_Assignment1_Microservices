package order

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    UserID    string             `bson:"user_id" json:"user_id"`
    Products  []OrderItem        `bson:"products" json:"products"`
    Status    string             `bson:"status" json:"status"` // e.g., "pending", "completed"
    CreatedAt int64              `bson:"created_at" json:"created_at"`
}

type OrderItem struct {
    ProductID string `bson:"product_id" json:"product_id"`
    Quantity  int    `bson:"quantity" json:"quantity"`
}
