package user

import (
	"context"
	"errors"

	"github.com/YoonDongGeun/go-backend-practice/internal/store"
	"github.com/jackc/pgx/v5"
)

var ErrNotFound = errors.New("user not found")

type Repository interface {
	Create(ctx context.Context, input CreateUserInput) (User, error)
	GetByID(ctx context.Context, id int64) (User, error)
	List(ctx context.Context, limit, offset int32) ([]User, error)
}

type PostgresRepository struct {
	queries *store.Queries
}

func NewPostgresRepository(queries *store.Queries) *PostgresRepository {
	return &PostgresRepository{queries: queries}
}

func (r *PostgresRepository) Create(ctx context.Context, input CreateUserInput) (User, error) {
	row, err := r.queries.CreateUser(ctx, store.CreateUserParams{
		Email: input.Email,
		Name:  input.Name,
	})
	if err != nil {
		return User{}, err
	}
	return toUser(row), nil
}

func (r *PostgresRepository) GetByID(ctx context.Context, id int64) (User, error) {
	row, err := r.queries.GetUser(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return User{}, ErrNotFound
	}
	if err != nil {
		return User{}, err
	}
	return toUser(row), nil
}

func (r *PostgresRepository) List(ctx context.Context, limit, offset int32) ([]User, error) {
	rows, err := r.queries.ListUsers(ctx, store.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	users := make([]User, 0, len(rows))
	for _, row := range rows {
		users = append(users, toUser(row))
	}
	return users, nil
}

func toUser(row store.User) User {
	return User{
		ID:        row.ID,
		Email:     row.Email,
		Name:      row.Name,
		CreatedAt: row.CreatedAt.Time,
	}
}
