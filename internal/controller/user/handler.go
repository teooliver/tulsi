package user

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
	"github.com/teooliver/kanban/internal/repository/user"
	"github.com/teooliver/kanban/pkg/postgresutils"
)

type userService interface {
	ListAllUsers(ctx context.Context, params *postgresutils.PageRequest) (postgresutils.Page[user.User], error)
	CreateUser(ctx context.Context, user user.UserForCreate) (string, error)
	DeleteUser(ctx context.Context, userID string) (string, error)
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
	Users postgresutils.Page[user.User] `json:"users"`
}

func (h Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	input := r.Context().Value(httpin.Input).(*postgresutils.PageRequest)
	slog.Info("HTTPin Input: ", input)

	users, err := h.service.ListAllUsers(ctx, input)
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

	id, err := h.service.CreateUser(ctx, userToCreate)
	w.Write([]byte(id))

}

func (h Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := chi.URLParam(r, "id")
	id, err := h.service.DeleteUser(ctx, userID)

	if err != nil {
		// Should return Error Not Found and 404
		slog.Info("UserID %+v\n", userID, userID)
	}

	w.Write([]byte(id))

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
