package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/gibalmeida/go-jobs/ent"
	"github.com/gibalmeida/go-jobs/graph/generated"
	"github.com/gibalmeida/go-jobs/graph/model"
)

func (r *mutationResolver) CreateApplicant(ctx context.Context, user model.UserInput) (*ent.User, error) {
	client := ent.FromContext(ctx)
	return client.User.
		Create().
		SetName(user.Name).
		SetEmail(user.Email).
		SetPassword(user.Password).
		SetRole("APPLICANT").
		Save(ctx)
}

func (r *mutationResolver) CreateDepartment(ctx context.Context, department model.DepartmentInput) (*ent.Department, error) {
	client := ent.FromContext(ctx)
	return client.Department.
		Create().
		SetName(department.Name).
		Save(ctx)
}

func (r *mutationResolver) CreateJob(ctx context.Context, job *model.JobInput) (*ent.Job, error) {
	// client := ent.FromContext(ctx)
	return r.client.Job.
		Create().
		SetName(job.Name).
		SetDescription(job.Description).
		SetDepartmentID(job.Department).
		Save(ctx)
}

func (r *mutationResolver) UpdateDepartmentManager(ctx context.Context, department *int, manager *int) (*ent.Department, error) {
	return r.client.Department.UpdateOneID(*department).SetManagerID(*manager).Save(ctx)
}

func (r *mutationResolver) UpdateJobDepartment(ctx context.Context, job *int, department *int) (*ent.Job, error) {
	return r.client.Job.UpdateOneID(*job).SetDepartmentID(*department).Save(ctx)
}

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder) (*ent.UserConnection, error) {
	return r.client.User.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithUserOrder(orderBy),
		)
}

func (r *queryResolver) Departments(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.DepartmentOrder) (*ent.DepartmentConnection, error) {
	return r.client.Department.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithDepartmentOrder(orderBy),
		)
}

func (r *queryResolver) Jobs(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.JobOrder) (*ent.JobConnection, error) {
	return r.client.Job.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithJobOrder(orderBy),
		)
}

func (r *userResolver) Role(ctx context.Context, obj *ent.User) (model.UserRole, error) {
	return model.UserRole(obj.Role), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
