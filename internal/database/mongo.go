package database

import (
	"context"
	"log"

	"feldrise.com/inventory-exercice/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoContext = context.TODO()

var CollectionInventories *mongo.Collection
var CollectionInventoryItems *mongo.Collection
var CollectionUsers *mongo.Collection

func Init() {
	connectionString := config.Cfg.Database.ConnectionString
	databaseName := config.Cfg.Database.Name

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(MongoContext, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(MongoContext, nil)

	if err != nil {
		log.Fatal(err)
	}

	CollectionInventories = client.Database(databaseName).Collection(config.Cfg.Collections.Inventories)
	CollectionInventoryItems = client.Database(databaseName).Collection(config.Cfg.Collections.InventoryItems)
	CollectionUsers = client.Database(databaseName).Collection(config.Cfg.Collections.Users)
}
