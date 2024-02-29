package delivery

import (
	"architecture_go/services/contact/cmd/internal/usecase"
	"context"
	"net/http"
	"strconv"
)

type GroupHandler struct {
	usecase usecase.ContactUsecase
}

func NewGroupHandler(usecase usecase.ContactUsecase) *GroupHandler {
	return &GroupHandler{
		usecase: usecase,
	}
}

func (h *GroupHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid group ID", http.StatusBadRequest)
		return
	}

	ctx := context.WithValue(r.Context(), "ID", id)

	h.usecase.CreateGroup(ctx, w, r)
}

func (h *GroupHandler) GetGroup(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid group ID", http.StatusBadRequest)
		return
	}

	ctx := context.WithValue(r.Context(), "ID", id)
	h.usecase.GetGroup(ctx, w, r)
}
