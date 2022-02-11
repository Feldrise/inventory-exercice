package inventories

import (
	"log"

	"feldrise.com/inventory-exercice/graph/model"
	"feldrise.com/inventory-exercice/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Inventory struct {
	ID     primitive.ObjectID `bson:"_id"`
	Name   string             `bson:"name"`
	UserID primitive.ObjectID `bson:"user_id"`
}

func (inventory *Inventory) ToModel() *model.Inventory {
	return &model.Inventory{
		ID:     inventory.ID.Hex(),
		Name:   inventory.Name,
		UserID: inventory.UserID.Hex(),
	}
}

func Create(input model.NewInventory, userID string) *Inventory {
	userObjectID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		log.Fatal(err)
	}

	databaseInventory := Inventory{
		ID:     primitive.NewObjectID(),
		Name:   input.Name,
		UserID: userObjectID,
	}

	_, err = database.CollectionInventories.InsertOne(database.MongoContext, databaseInventory)

	if err != nil {
		log.Fatal(err)
	}

	return &databaseInventory
}

func GetAll() ([]Inventory, error) {
	filter := bson.D{{}}

	return GetFiltered(filter)
}

func GetAllForUser(userID string) ([]Inventory, error) {
	userObjectID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return nil, err
	}

	filter := bson.D{
		primitive.E{
			Key:   "user_id",
			Value: userObjectID,
		},
	}

	return GetFiltered(filter)
}

func GetById(id string) (*Inventory, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.D{
		primitive.E{
			Key:   "_id",
			Value: objectId,
		},
	}

	inventories, err := GetFiltered(filter)

	if err != nil {
		return nil, err
	}

	if len(inventories) == 0 {
		return nil, nil
	}

	return &inventories[0], nil
}

func GetFiltered(filter interface{}) ([]Inventory, error) {
	inventories := []Inventory{}

	cursor, err := database.CollectionInventories.Find(database.MongoContext, filter)

	if err != nil {
		return inventories, err
	}

	for cursor.Next(database.MongoContext) {
		var inventory Inventory

		err := cursor.Decode(&inventory)

		if err != nil {
			return inventories, err
		}

		inventories = append(inventories, inventory)
	}

	if err := cursor.Err(); err != nil {
		return inventories, err
	}

	return inventories, nil
}
