package user

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
	"github.com/teooliver/kanban/internal/repository/user"
	"github.com/teooliver/kanban/pkg/auth"
	"github.com/teooliver/kanban/pkg/postgresutils"
)

type userService interface {
	ListAllUsers(ctx context.Context, params *postgresutils.PageRequest) (postgresutils.Page[user.User], error)
	CreateUser(ctx context.Context, user user.UserForCreate) (string, error)
	DeleteUser(ctx context.Context, userID string) (string, error)
	UpdateUser(ctx context.Context, userID string, updatedUser user.UserForUpdate) error
	GetUserByEmail(ctx context.Context, email string) (user user.User, err error)
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

	users, err := h.service.ListAllUsers(ctx, input)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Something went wrong: %v\n", err)))
		return
	}

	userResponse := ListUserResponse{
		Users: users,
	}

	jsonUsers, err := json.Marshal(userResponse.Users)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Something went wrong: %v\n", err)))
		return
	}

	w.WriteHeader(http.StatusOK)
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
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id))

}

func (h Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := chi.URLParam(r, "id")
	id, err := h.service.DeleteUser(ctx, userID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		slog.Info("UserID %+v\n", userID, userID)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id))

}

// TODO: Return updated user
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
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

}

// Auth Functions
func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Testing using FormValue insteado of Json in this case
	// username is always email
	username := r.FormValue("username")
	password := r.FormValue("password")

	slog.Info("CONTROLER => GET USER BY EMAIL", username)

	u, err := h.service.GetUserByEmail(ctx, username)
	slog.Info("CONTROLER => ERROR", u, err)
	if err != nil || !auth.CheckPasswordHash(password, u.Login.HashedPassword) {
		err := http.StatusNotFound
		http.Error(w, "Invalid username or password", err)
		return
	}

	sessionToken := auth.GenerateToken(32)
	csrfToken := auth.GenerateToken(32)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
	})

	// // Store tokens in the database
	var userLogin = user.Login{
		HashedPassword: u.Login.HashedPassword,
		SessionToken:   sessionToken,
		CSRFToken:      csrfToken,
	}
	// u.Login.SessionToken = sessionToken
	// u.Login.CSRFToken = csrfToken

	var userToUpdate = user.UserForUpdate{
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Login:     userLogin,
	}

	// TODO: UpdateUser
	h.service.UpdateUser(ctx, u.ID, userToUpdate)
	// users[username] = user

	fmt.Fprintln(w, "Login successful!")

}
