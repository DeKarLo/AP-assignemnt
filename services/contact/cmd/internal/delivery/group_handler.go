// group_handler.go
package delivery

import (
	"architecture_go/services/contact/cmd/internal/usecase"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

type GroupHandler struct {
	usecase usecase.ContactUsecase
}

func NewGroupHandler(usecase usecase.ContactUsecase) *GroupHandler {
	return &GroupHandler{
		usecase: usecase,
	}
}

type contextKey string

const (
	contextKeyID contextKey = "ID"
)

func (h *GroupHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), contextKeyID, uuid.New().String())

	name := r.FormValue("name")

	err := h.usecase.CreateGroup(ctx, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *GroupHandler) GetGroup(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("ID").(string)

	ctx := context.WithValue(r.Context(), "ID", id)

	idStr := r.FormValue("id")
	groupID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid group ID", http.StatusBadRequest)
		return
	}

	group, err := h.usecase.GetGroupByID(ctx, groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(group)
}
