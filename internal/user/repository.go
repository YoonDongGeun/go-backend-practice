package user

import (
	"context"
	"errors"

	"github.com/YoonDongGeun/go-backend-practice/internal/store"
	"github.com/jackc/pgx/v5"
)

var ErrNotFound = errors.New("user not found")

type Repository interface {
	Create(ctx context.Context, command CreateUserCommand) (UserEntity, error)
	GetByID(ctx context.Context, id int64) (UserEntity, error)
	List(ctx context.Context, limit, offset int32) ([]UserEntity, error)
}

type PostgresRepository struct {
	queries *store.Queries
}

func NewPostgresRepository(queries *store.Queries) *PostgresRepository {
	return &PostgresRepository{queries: queries}
}

func (r *PostgresRepository) Create(ctx context.Context, command CreateUserCommand) (UserEntity, error) {
	row, err := r.queries.CreateUser(ctx, store.CreateUserParams{
		Email: command.Email,
		Name:  command.Name,
	})
	if err != nil {
		return UserEntity{}, err
	}
	return toUserEntity(row), nil
}

func (r *PostgresRepository) GetByID(ctx context.Context, id int64) (UserEntity, error) {
	row, err := r.queries.GetUser(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return UserEntity{}, ErrNotFound
	}
	if err != nil {
		return UserEntity{}, err
	}
	return toUserEntity(row), nil
}

func (r *PostgresRepository) List(ctx context.Context, limit, offset int32) ([]UserEntity, error) {
	rows, err := r.queries.ListUsers(ctx, store.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	users := make([]UserEntity, 0, len(rows))
	for _, row := range rows {
		users = append(users, toUserEntity(row))
	}
	return users, nil
}
