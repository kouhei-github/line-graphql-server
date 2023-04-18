package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.29

import (
	"context"
	"fmt"

	"github.com/meetup/graph/model"
	"github.com/meetup/internal"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, data model.UserCreateInput) (*model.User, error) {
	//mutation CreateUser {
	//	createUser(data: {
	//		name: "kohei",
	//		email: "kouheidesu09@gmail.com"
	//	}) {
	//	id
	//	username
	//	email
	//}
	//}
	return r.Srv.CreateUserRecode(data.Name, data.Email)
}

// Meetups is the resolver for the meetups field.
func (r *queryResolver) Meetups(ctx context.Context) ([]*model.Meetup, error) {
	panic(fmt.Errorf("not implemented: Meetups - meetups"))
	//return r.Srv.GetUserByName("name")
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	//query Users{
	//	users {
	//		id
	//		username
	//		email
	//	}
	//}
	return r.Srv.GetUserLimited(3)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, name string) (*model.User, error) {
	//query {
	//	user(name: "nagamatsu") {
	//	id
	//	username
	//	email
	//}
	//}
	return r.Srv.GetUserByName(name)
}

// Mutation returns internal.MutationResolver implementation.
func (r *Resolver) Mutation() internal.MutationResolver { return &mutationResolver{r} }

// Query returns internal.QueryResolver implementation.
func (r *Resolver) Query() internal.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
