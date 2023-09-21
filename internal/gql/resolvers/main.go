package resolvers

import (
	"context"
	"github.com/Earl-Power/gqlgengin/internal/gql"
	"github.com/Earl-Power/gqlgengin/internal/gql/models"
)

type Resolver struct{}

func (r *Resolver) Mutation() gql.MutationResolver {
	return &MutationResolver(r)
}

func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (bool, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context, userID *string) ([]*models.User, error) {
	var tempID = "ec17af15-e354-440c-a09f-69715fc8b595"
	var tempEmail = "you@email.com"
	var tempUserID = "UserID-1"
	records := []*models.User{
		&models.User{
			ID:     &tempID,
			Email:  &tempEmail,
			UserID: &tempUserID,
		},
	}
	return records, nil
}
