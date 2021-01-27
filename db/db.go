package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Instance is the mongo database instance
var Instance *mongo.Database

// StartDB sets default values for database instance and context
func StartDB() {
	client, _ := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DATABASE")))
	ctx := Ctx()
	_ = client.Connect(ctx)
	Instance = client.Database("senhor-bot")
}

// Ctx creates a context with timeout
func Ctx() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}
