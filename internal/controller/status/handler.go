package status

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/teooliver/kanban/internal/repository/status"
)

type statusService interface {
	ListAllStatus(ctx context.Context) ([]status.Status, error)
	CreateStatus(ctx context.Context, status status.StatusForCreate) error
	DeleteStatus(ctx context.Context, status string) error
	UpdateStatus(ctx context.Context, status string, updatedStatus status.StatusForUpdate) error
}

type Handler struct {
	service statusService
}

func New(service statusService) Handler {
	return Handler{
		service: service,
	}
}

// TODO: Add pagination
type ListStatusResponse struct {
	StatusList []status.Status `json:"status_list"`
}

func (h Handler) ListStatuss(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	statusList, err := h.service.ListAllStatus(ctx)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Something went wrong: %v\n", err)))
	}
	StatusResponse := ListStatusResponse{
		StatusList: statusList,
	}

	jsonStatusList, err := json.Marshal(StatusResponse.StatusList)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Something went wrong: %v\n", err)))
	}

	w.Write([]byte(jsonStatusList))
}

func (h Handler) CreateStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var newStatus status.StatusForCreate
	err := json.NewDecoder(r.Body).Decode(&newStatus)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CreateStatus(ctx, newStatus)

}

func (h Handler) DeleteStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	statusID := chi.URLParam(r, "id")
	err := h.service.DeleteStatus(ctx, statusID)

	if err != nil {
		// Should return Error Not Found and 404
		slog.Info("ListID %+v\n", statusID, statusID)
	}

}

func (h Handler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	statusID := chi.URLParam(r, "id")

	var statusToUpdate status.StatusForUpdate
	err := json.NewDecoder(r.Body).Decode(&statusToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.UpdateStatus(ctx, statusID, statusToUpdate)

	if err != nil {
		// Should return Error Not Found and 404
		print(err)
	}

}
