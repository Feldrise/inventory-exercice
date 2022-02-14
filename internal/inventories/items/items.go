package items

import (
	"log"

	"feldrise.com/inventory-exercice/graph/model"
	"feldrise.com/inventory-exercice/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InventoryItem struct {
	ID          primitive.ObjectID `bson:"_id"`
	InventoryID primitive.ObjectID `bson:"inventory_id"`
	Name        string             `bson:"name"`
	Quantity    int                `bons:"quantity"`
}

func (item *InventoryItem) ToModel() *model.InventoryItem {
	return &model.InventoryItem{
		ID:       item.ID.Hex(),
		Name:     item.Name,
		Quantity: item.Quantity,
	}
}

func Create(input model.NewInventoryItem) *InventoryItem {
	inventoryObjectID, err := primitive.ObjectIDFromHex(input.InventoryID)

	if err != nil {
		log.Fatal(err)
	}

	itemQuantity := 1
	if input.Quantity != nil {
		itemQuantity = *input.Quantity
	}

	databaseItem := InventoryItem{
		ID:          primitive.NewObjectID(),
		InventoryID: inventoryObjectID,
		Name:        input.Name,
		Quantity:    itemQuantity,
	}

	_, err = database.CollectionInventoryItems.InsertOne(database.MongoContext, databaseItem)

	if err != nil {
		log.Fatal(err)
	}

	return &databaseItem
}

func Update(changes *InventoryItem) {
	filter := bson.D{
		primitive.E{
			Key:   "_id",
			Value: changes.ID,
		},
	}

	_, err := database.CollectionInventoryItems.ReplaceOne(database.MongoContext, filter, changes)

	if err != nil {
		log.Fatal(err)
	}
}

func (item InventoryItem) IsLast() bool {
	filter := bson.D{
		primitive.E{
			Key:   "inventory_id",
			Value: item.InventoryID,
		},
	}

	opts := options.FindOptions{}
	opts.SetLimit(1)
	opts.SetSort(bson.D{
		primitive.E{
			Key: "_id", Value: -1,
		},
	})

	lastItem, err := GetFiltered(filter, &opts)

	if err != nil || len(lastItem) <= 0 {
		return false
	}

	return lastItem[0].ID == item.ID
}

func GetAll() ([]InventoryItem, error) {
	filter := bson.D{{}}

	return GetFiltered(filter, nil)
}

func GetAllForInventory(inventoryID string) ([]InventoryItem, error) {
	inventoryObjectID, err := primitive.ObjectIDFromHex(inventoryID)

	if err != nil {
		return nil, err
	}

	filter := bson.D{
		primitive.E{
			Key:   "inventory_id",
			Value: inventoryObjectID,
		},
	}

	return GetFiltered(filter, nil)
}

func GetById(id string) (*InventoryItem, error) {
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

	inventoryItems, err := GetFiltered(filter, nil)

	if err != nil {
		return nil, err
	}

	if len(inventoryItems) == 0 {
		return nil, nil
	}

	return &inventoryItems[0], nil
}

func GetPaginated(inventoryID string, startValue *string, first int) ([]InventoryItem, error) {
	inventoryObjectID, err := primitive.ObjectIDFromHex(inventoryID)

	if err != nil {
		return nil, err
	}

	var filter interface{}

	if startValue != nil {
		objectId, err := primitive.ObjectIDFromHex(*startValue)

		if err != nil {
			return nil, err
		}

		filter = bson.M{
			"inventory_id": inventoryObjectID,
			"_id": bson.M{
				"$gt": objectId,
			},
		}
	} else {
		filter = bson.D{
			primitive.E{
				Key:   "inventory_id",
				Value: inventoryObjectID,
			},
		}
	}

	opts := options.Find()
	opts.SetLimit(int64(first))

	return GetFiltered(filter, opts)
}

func GetFiltered(filter interface{}, options *options.FindOptions) ([]InventoryItem, error) {
	inventoryItems := []InventoryItem{}

	cursor, err := database.CollectionInventoryItems.Find(database.MongoContext, filter, options)

	if err != nil {
		return inventoryItems, err
	}

	for cursor.Next(database.MongoContext) {
		var inventoryItem InventoryItem

		err := cursor.Decode(&inventoryItem)

		if err != nil {
			return inventoryItems, err
		}

		inventoryItems = append(inventoryItems, inventoryItem)
	}

	if err := cursor.Err(); err != nil {
		return inventoryItems, err
	}

	return inventoryItems, nil
}
