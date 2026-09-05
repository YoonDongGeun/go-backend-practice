package user

import (
	"context"

	"github.com/YoonDongGeun/go-backend-practice/internal/store"
)

type Service struct {
	queries *store.Queries
}

func NewService(queries *store.Queries) *Service {
	return &Service{queries: queries}
}

func (s *Service) Create(ctx context.Context, params store.CreateUserParams) (store.User, error) {
	return s.queries.CreateUser(ctx, params)
}

func (s *Service) Get(ctx context.Context, id int64) (store.User, error) {
	return s.queries.GetUser(ctx, id)
}

func (s *Service) List(ctx context.Context, limit, offset int32) ([]store.User, error) {
	return s.queries.ListUsers(ctx, store.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
}
