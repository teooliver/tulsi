package status

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
	"github.com/teooliver/kanban/internal/repository/status"
	"github.com/teooliver/kanban/pkg/postgresutils"
)

type statusService interface {
	ListAllStatus(ctx context.Context, params *postgresutils.PageRequest) (postgresutils.Page[status.Status], error)
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

type ListStatusResponse struct {
	StatusList postgresutils.Page[status.Status] `json:"status_list"`
}

func (h Handler) ListAllStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	input := r.Context().Value(httpin.Input).(*postgresutils.PageRequest)
	slog.Info("HTTPin Input: ", input)

	statusList, err := h.service.ListAllStatus(ctx, input)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("STATUS HANDLER => Something went wrong: %v\n", err)))
	}
	StatusResponse := ListStatusResponse{
		StatusList: statusList,
	}

	jsonStatusList, err := json.Marshal(StatusResponse.StatusList)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("STATUS HANDLER MARSHAL =>Something went wrong: %v\n", err)))
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
