package delivery

import (
	"architecture_go/services/contact/cmd/internal/usecase"
	"context"
	"net/http"
	"strconv"
)

type ContactHandler struct {
	usecase usecase.ContactUsecase
}

func NewContactHandler(usecase usecase.ContactUsecase) *ContactHandler {
	return &ContactHandler{
		usecase: usecase,
	}
}

func (h *ContactHandler) CreateContact(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid contact ID", http.StatusBadRequest)
		return
	}

	ctx := context.WithValue(r.Context(), "ID", id)
	h.usecase.CreateContact(ctx, w, r)
}

func (h *ContactHandler) GetContact(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid contact ID", http.StatusBadRequest)
		return
	}

	ctx := context.WithValue(r.Context(), "ID", id)
	h.usecase.GetContact(ctx, w, r)
}

func (h *ContactHandler) UpdateContact(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid contact ID", http.StatusBadRequest)
		return
	}

	ctx := context.WithValue(r.Context(), "ID", id)
	h.usecase.UpdateContact(ctx, w, r)
}

func (h *ContactHandler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid contact ID", http.StatusBadRequest)
		return
	}

	ctx := context.WithValue(r.Context(), "ID", id)
	h.usecase.DeleteContact(ctx, w, r)
}
