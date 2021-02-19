package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/gibalmeida/go-jobs/graph/generated"
	"github.com/gibalmeida/go-jobs/graph/model"
)

func (r *departmentResolver) Manager(ctx context.Context, obj *model.Department) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *jobResolver) Department(ctx context.Context, obj *model.Job) (*model.Department, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateAdmin(ctx context.Context, input model.NewAdmin) (*model.User, error) {
	admin := &model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Role:     model.RoleAdmin,
	}
	r.users = append(r.users, admin)
	return admin, nil
}

func (r *mutationResolver) CreateManager(ctx context.Context, input model.NewManager) (*model.User, error) {
	admin := &model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Role:     model.RoleManager,
	}
	r.users = append(r.users, admin)
	return admin, nil
}

func (r *mutationResolver) CreateApplicant(ctx context.Context, input model.NewApplicant) (*model.User, error) {
	applicant := &model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Role:     model.RoleApplicant,
	}
	r.users = append(r.users, applicant)
	return applicant, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// Department returns generated.DepartmentResolver implementation.
func (r *Resolver) Department() generated.DepartmentResolver { return &departmentResolver{r} }

// Job returns generated.JobResolver implementation.
func (r *Resolver) Job() generated.JobResolver { return &jobResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type departmentResolver struct{ *Resolver }
type jobResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) CreateManagwer(ctx context.Context, input model.NewManager) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
