package user

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
)

func TestHandlerCreateUser(t *testing.T) {
	repo := &fakeRepository{
		createFunc: func(ctx context.Context, command CreateUserCommand) (UserEntity, error) {
			return UserEntity{
				ID:        1,
				Email:     command.Email,
				Name:      command.Name,
				CreatedAt: time.Date(2026, 9, 5, 0, 0, 0, 0, time.UTC),
			}, nil
		},
	}
	router := newTestRouter(repo)

	body := bytes.NewBufferString(`{"email":"yoon@example.com","name":"Yoon"}`)
	req := httptest.NewRequest(http.MethodPost, "/users/", body)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d: %s", http.StatusCreated, rec.Code, rec.Body.String())
	}

	var got UserResponseDTO
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if got.Email != "yoon@example.com" {
		t.Fatalf("expected email yoon@example.com, got %s", got.Email)
	}
}

func TestHandlerGetUserNotFound(t *testing.T) {
	repo := &fakeRepository{
		getByIDFunc: func(ctx context.Context, id int64) (UserEntity, error) {
			return UserEntity{}, ErrNotFound
		},
	}
	router := newTestRouter(repo)

	req := httptest.NewRequest(http.MethodGet, "/users/999", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}
}

func newTestRouter(repo Repository) http.Handler {
	router := chi.NewRouter()
	service := NewService(repo)
	handler := NewHandler(service)
	handler.RegisterRoutes(router)
	return router
}
