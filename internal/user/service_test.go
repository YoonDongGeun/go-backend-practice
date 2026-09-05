package user

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestServiceCreate(t *testing.T) {
	repo := &fakeRepository{
		createFunc: func(ctx context.Context, command CreateUserCommand) (UserEntity, error) {
			if command.Email != "yoon@example.com" {
				t.Fatalf("expected email yoon@example.com, got %s", command.Email)
			}
			if command.Name != "Yoon" {
				t.Fatalf("expected name Yoon, got %s", command.Name)
			}

			return UserEntity{
				ID:        1,
				Email:     command.Email,
				Name:      command.Name,
				CreatedAt: time.Date(2026, 9, 5, 0, 0, 0, 0, time.UTC),
			}, nil
		},
	}

	service := NewService(repo)
	got, err := service.Create(context.Background(), CreateUserCommand{
		Email: "yoon@example.com",
		Name:  "Yoon",
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got.ID != 1 {
		t.Fatalf("expected user id 1, got %d", got.ID)
	}
}

func TestServiceGetReturnsRepositoryError(t *testing.T) {
	repo := &fakeRepository{
		getByIDFunc: func(ctx context.Context, id int64) (UserEntity, error) {
			return UserEntity{}, ErrNotFound
		},
	}

	service := NewService(repo)
	_, err := service.Get(context.Background(), 999)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

type fakeRepository struct {
	createFunc  func(context.Context, CreateUserCommand) (UserEntity, error)
	getByIDFunc func(context.Context, int64) (UserEntity, error)
	listFunc    func(context.Context, int32, int32) ([]UserEntity, error)
}

func (r *fakeRepository) Create(ctx context.Context, command CreateUserCommand) (UserEntity, error) {
	if r.createFunc != nil {
		return r.createFunc(ctx, command)
	}
	return UserEntity{}, nil
}

func (r *fakeRepository) GetByID(ctx context.Context, id int64) (UserEntity, error) {
	if r.getByIDFunc != nil {
		return r.getByIDFunc(ctx, id)
	}
	return UserEntity{}, nil
}

func (r *fakeRepository) List(ctx context.Context, limit, offset int32) ([]UserEntity, error) {
	if r.listFunc != nil {
		return r.listFunc(ctx, limit, offset)
	}
	return []UserEntity{}, nil
}
