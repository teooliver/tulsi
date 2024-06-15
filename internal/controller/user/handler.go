package user

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/teooliver/kanban/internal/repository/user"
)

type userService interface {
	ListAllUsers(ctx context.Context) ([]user.User, error)
	CreateUser(ctx context.Context, user user.UserForCreate) error
	DeleteUser(ctx context.Context, userID string) error
	UpdateUser(ctx context.Context, userID string, updatedUser user.UserForUpdate) error
}

type Handler struct {
	service userService
}

func New(service userService) Handler {
	return Handler{
		service: service,
	}
}

// TODO: Add pagination
type ListUserResponse struct {
	Users []user.User `json:"users"`
}

func (h Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := h.service.ListAllUsers(ctx)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Something went wrong: %v\n", err)))
	}
	userResponse := ListUserResponse{
		Users: users,
	}

	jsonUsers, err := json.Marshal(userResponse.Users)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Something went wrong: %v\n", err)))
	}

	w.Write([]byte(jsonUsers))
}

func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var userToCreate user.UserForCreate
	err := json.NewDecoder(r.Body).Decode(&userToCreate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	slog.Info("User for CREATE %+v\n", "userToCreate", userToCreate)

	err = h.service.CreateUser(ctx, userToCreate)

}

func (h Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := chi.URLParam(r, "id")
	err := h.service.DeleteUser(ctx, userID)

	if err != nil {
		// Should return Error Not Found and 404
		slog.Info("UserID %+v\n", userID, userID)
	}

}

func (h Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := chi.URLParam(r, "id")
	var userToUpdate user.UserForUpdate
	err := json.NewDecoder(r.Body).Decode(&userToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.UpdateUser(ctx, userID, userToUpdate)

	if err != nil {
		// Should return Error Not Found and 404
		print(err)
	}

}
