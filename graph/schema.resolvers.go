package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"feldrise.com/inventory-exercice/graph/generated"
	"feldrise.com/inventory-exercice/graph/model"
	"feldrise.com/inventory-exercice/internal/auth"
	"feldrise.com/inventory-exercice/internal/inventories"
	"feldrise.com/inventory-exercice/internal/inventories/items"
	"feldrise.com/inventory-exercice/internal/users"
	"feldrise.com/inventory-exercice/pkg/jwt"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *inventoryResolver) User(ctx context.Context, obj *model.Inventory) (*model.User, error) {
	user, err := users.GetUserById(obj.UserID)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, gqlerror.Errorf("the user doesn't exist for this inventory")
	}

	return user.ToModel(), nil
}

func (r *inventoryResolver) Items(ctx context.Context, obj *model.Inventory) ([]*model.InventoryItem, error) {
	databaseItems, err := items.GetAllForInventory(obj.ID)

	if err != nil {
		return nil, err
	}

	items := []*model.InventoryItem{}

	for _, databaseItem := range databaseItems {
		items = append(items, databaseItem.ToModel())
	}

	return items, nil
}

func (r *mutationResolver) CreateInventoryItem(ctx context.Context, input model.NewInventoryItem) (*model.InventoryItem, error) {
	user := auth.ForContext(ctx)

	if user == nil {
		return nil, gqlerror.Errorf("access denied")
	}

	databaseInventory, err := inventories.GetById(input.InventoryID)

	if err != nil {
		return nil, err
	}

	if databaseInventory == nil {
		return nil, gqlerror.Errorf("the inventory doesn't exist")
	}

	if databaseInventory.UserID.Hex() != user.ID {
		return nil, gqlerror.Errorf("you don't own this inventory")
	}

	databaseItem := items.Create(input)

	return databaseItem.ToModel(), nil
}

func (r *mutationResolver) CreateInventory(ctx context.Context, input *model.NewInventory) (*model.Inventory, error) {
	user := auth.ForContext(ctx)

	if user == nil {
		return nil, gqlerror.Errorf("access denied")
	}

	databaseInventory := inventories.Create(*input, user.ID)

	return databaseInventory.ToModel(), nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	existingUser, _ := users.GetUserByEmail(input.Email)

	if existingUser != nil {
		return nil, gqlerror.Errorf("a user with this email already exists")
	}

	databaseUser := users.Create(input)

	return databaseUser.ToModel(), nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	isPasswordCorrect := users.Authenticate(input)

	if !isPasswordCorrect {
		return "", gqlerror.Errorf("wrong username or password")
	}

	user, err := users.GetUserByEmail(input.Email)

	if user == nil || err != nil {
		return "", err
	}

	token, err := jwt.GenerateToken(user.ID.Hex())

	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	id, err := jwt.ParseToken(input.Token)

	if err != nil {
		return "", gqlerror.Errorf("access denied")
	}

	token, err := jwt.GenerateToken(id)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *queryResolver) InventoryItem(ctx context.Context, id string) (*model.InventoryItem, error) {
	user := auth.ForContext(ctx)

	if user == nil {
		return nil, gqlerror.Errorf("access denied")
	}

	databaseItem, err := items.GetById(id)

	if err != nil {
		return nil, err
	}

	if databaseItem == nil {
		return nil, gqlerror.Errorf("there is no item matching this ID")
	}

	// We need to check the item actually belong to the user
	databaseInventory, err := inventories.GetById(databaseItem.InventoryID.Hex())

	if err != nil || databaseInventory == nil {
		return nil, gqlerror.Errorf("it appears that the item's inventory can't be get")
	}

	if databaseInventory.UserID.Hex() != user.ID {
		return nil, gqlerror.Errorf("you don't own this item")
	}

	return databaseItem.ToModel(), nil
}

func (r *queryResolver) Inventories(ctx context.Context) ([]*model.Inventory, error) {
	user := auth.ForContext(ctx)

	if user == nil {
		return nil, gqlerror.Errorf("access denied")
	}

	databaseInventories, err := inventories.GetAllForUser(user.ID)

	if err != nil {
		return nil, err
	}

	inventories := []*model.Inventory{}

	for _, databaseInventory := range databaseInventories {
		inventory := databaseInventory.ToModel()

		inventories = append(inventories, inventory)
	}

	return inventories, nil
}

func (r *queryResolver) Inventory(ctx context.Context, id string) (*model.Inventory, error) {
	user := auth.ForContext(ctx)

	if user == nil {
		return nil, gqlerror.Errorf("access denied")
	}

	databaseInventory, err := inventories.GetById(id)

	if err != nil {
		return nil, err
	}

	if databaseInventory == nil {
		return nil, gqlerror.Errorf("there is no inventory with this id")
	}

	if databaseInventory.UserID.Hex() != user.ID {
		return nil, gqlerror.Errorf("you don't own this inventory")
	}

	return databaseInventory.ToModel(), nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Inventory returns generated.InventoryResolver implementation.
func (r *Resolver) Inventory() generated.InventoryResolver { return &inventoryResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type inventoryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) InventoryItems(ctx context.Context, inventory string) ([]*model.InventoryItem, error) {
	user := auth.ForContext(ctx)

	if user == nil {
		return nil, gqlerror.Errorf("access denied")
	}

	databaseInventory, err := inventories.GetById(inventory)

	if err != nil {
		return nil, err
	}

	if databaseInventory == nil {
		return nil, gqlerror.Errorf("the inventory doesn't exist")
	}

	if databaseInventory.UserID.Hex() != user.ID {
		return nil, gqlerror.Errorf("You don't own this inventory")
	}

	databaseItems, err := items.GetAllForInventory(inventory)

	if err != nil {
		return nil, err
	}

	items := []*model.InventoryItem{}

	for _, item := range databaseItems {
		items = append(items, item.ToModel())
	}

	return items, nil
}
