package order

import (
    "context"
    "time"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
    Create(ctx context.Context, order Order) (primitive.ObjectID, error)
    GetByID(ctx context.Context, id primitive.ObjectID) (Order, error)
    UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error
    List(ctx context.Context) ([]Order, error)
}

type repo struct {
    collection *mongo.Collection
}

func NewRepository(db *mongo.Database) Repository {
    return &repo{collection: db.Collection("orders")}
}

func (r *repo) Create(ctx context.Context, o Order) (primitive.ObjectID, error) {
    o.CreatedAt = time.Now().Unix()
    res, err := r.collection.InsertOne(ctx, o)
    if err != nil {
        return primitive.NilObjectID, err
    }
    return res.InsertedID.(primitive.ObjectID), nil
}

func (r *repo) GetByID(ctx context.Context, id primitive.ObjectID) (Order, error) {
    var o Order
    err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&o)
    return o, err
}

func (r *repo) UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error {
    _, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"status": status}})
    return err
}

func (r *repo) List(ctx context.Context) ([]Order, error) {
    cursor, err := r.collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var orders []Order
    for cursor.Next(ctx) {
        var o Order
        if err := cursor.Decode(&o); err != nil {
            return nil, err
        }
        orders = append(orders, o)
    }
    return orders, nil
}
