package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"feldrise.com/inventory-exercice/graph/generated"
	"feldrise.com/inventory-exercice/graph/model"
)

func (r *mutationResolver) CreateInventoryItem(ctx context.Context, input model.NewInventoryItem) (*model.InventoryItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateInventory(ctx context.Context, input *model.NewInventory) (*model.Inventory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
