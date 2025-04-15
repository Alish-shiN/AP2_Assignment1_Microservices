package product

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
    Create(ctx context.Context, product Product) (primitive.ObjectID, error)
    GetByID(ctx context.Context, id primitive.ObjectID) (Product, error)
    Update(ctx context.Context, id primitive.ObjectID, product Product) error
    Delete(ctx context.Context, id primitive.ObjectID) error
    List(ctx context.Context) ([]Product, error)
}

type repo struct {
    collection *mongo.Collection
}

func NewRepository(db *mongo.Database) Repository {
    return &repo{collection: db.Collection("products")}
}

func (r *repo) Create(ctx context.Context, p Product) (primitive.ObjectID, error) {
    res, err := r.collection.InsertOne(ctx, p)
    if err != nil {
        return primitive.NilObjectID, err
    }
    return res.InsertedID.(primitive.ObjectID), nil
}

func (r *repo) GetByID(ctx context.Context, id primitive.ObjectID) (Product, error) {
    var p Product
    err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&p)
    return p, err
}

func (r *repo) Update(ctx context.Context, id primitive.ObjectID, p Product) error {
    _, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": p})
    return err
}

func (r *repo) Delete(ctx context.Context, id primitive.ObjectID) error {
    _, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
    return err
}

func (r *repo) List(ctx context.Context) ([]Product, error) {
    cursor, err := r.collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var products []Product
    for cursor.Next(ctx) {
        var p Product
        if err := cursor.Decode(&p); err != nil {
            return nil, err
        }
        products = append(products, p)
    }
    return products, nil
}
