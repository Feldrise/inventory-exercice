package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"feldrise.com/inventory-exercice/graph/generated"
	"feldrise.com/inventory-exercice/graph/model"
	"feldrise.com/inventory-exercice/internal/users"
	"feldrise.com/inventory-exercice/pkg/jwt"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *inventoryResolver) User(ctx context.Context, obj *model.Inventory) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *inventoryResolver) Items(ctx context.Context, obj *model.Inventory) ([]*model.InventoryItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateInventoryItem(ctx context.Context, input model.NewInventoryItem) (*model.InventoryItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateInventory(ctx context.Context, input *model.NewInventory) (*model.Inventory, error) {
	panic(fmt.Errorf("not implemented"))
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

func (r *queryResolver) InventoryItems(ctx context.Context, inventory string) ([]*model.InventoryItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) InventoryItem(ctx context.Context, id string) (*model.InventoryItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Inventories(ctx context.Context) ([]*model.Inventory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Inventory(ctx context.Context, id string) (*model.Inventory, error) {
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
