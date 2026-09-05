package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/YoonDongGeun/go-backend-practice/internal/httpx"
	"github.com/YoonDongGeun/go-backend-practice/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

type Handler struct {
	service *Service
}

type createUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", h.create)
		r.Get("/", h.list)
		r.Get("/{id}", h.get)
	})
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.Error(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if req.Email == "" || req.Name == "" {
		httpx.Error(w, http.StatusBadRequest, "email and name are required")
		return
	}

	user, err := h.service.Create(r.Context(), store.CreateUserParams{
		Email: req.Email,
		Name:  req.Name,
	})
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, "failed to create user")
		return
	}

	httpx.JSON(w, http.StatusCreated, user)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		httpx.Error(w, http.StatusBadRequest, "invalid user id")
		return
	}

	user, err := h.service.Get(r.Context(), id)
	if errors.Is(err, pgx.ErrNoRows) {
		httpx.Error(w, http.StatusNotFound, "user not found")
		return
	}
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, "failed to get user")
		return
	}

	httpx.JSON(w, http.StatusOK, user)
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	limit := parseQueryInt32(r, "limit", 20)
	offset := parseQueryInt32(r, "offset", 0)

	users, err := h.service.List(r.Context(), limit, offset)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, "failed to list users")
		return
	}

	httpx.JSON(w, http.StatusOK, users)
}

func parseQueryInt32(r *http.Request, key string, fallback int32) int32 {
	value := r.URL.Query().Get(key)
	if value == "" {
		return fallback
	}

	parsed, err := strconv.ParseInt(value, 10, 32)
	if err != nil || parsed < 0 {
		return fallback
	}
	return int32(parsed)
}
