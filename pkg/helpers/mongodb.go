package tools

import (
	"context"
	"log"

	"github.com/vkunssec/blockchain/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// StringToObjectId converts a string to a primitive.ObjectID
func StringToObjectId(id string) primitive.ObjectID {
	s, e := primitive.ObjectIDFromHex(id)
	if e != nil {
		log.Panic(e)
	}
	return s
}

// ArrayStringToObjectId converts an array of strings to an array of primitive.ObjectID
func ArrayStringToObjectId(arrString []string) []primitive.ObjectID {
	arrObjectId := make([]primitive.ObjectID, len(arrString))
	for i := range arrString {
		arrObjectId[i] = StringToObjectId(arrString[i])
	}
	return arrObjectId
}

// InsertOne inserts a single value into a collection
func InsertOne(ctx context.Context, collection string, values interface{}, options ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return database.MongoDB.
		Collection(collection).
		InsertOne(ctx, values, options...)
}

// UpdateOne updates a single value in a collection
func UpdateOne(ctx context.Context, collection string, filter primitive.M, values interface{}, options ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return database.MongoDB.
		Collection(collection).
		UpdateOne(ctx, filter, bson.M{"$set": values}, options...)
}

// Find finds multiple values in a collection
func Find(ctx context.Context, collection string, filters primitive.M, options ...*options.FindOptions) (*mongo.Cursor, error) {
	return database.MongoDB.
		Collection(collection).
		Find(ctx, filters, options...)
}

// FindOne finds a single value in a collection
func FindOne(ctx context.Context, collection string, filters primitive.M, options ...*options.FindOneOptions) *mongo.SingleResult {
	return database.MongoDB.
		Collection(collection).
		FindOne(ctx, filters, options...)
}

// DeleteOne deletes a single value from a collection
func DeleteOne(ctx context.Context, collection string, filters primitive.M, options ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return database.MongoDB.
		Collection(collection).
		DeleteOne(ctx, filters, options...)
}
