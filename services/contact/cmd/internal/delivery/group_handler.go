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

func (h *GroupHandler) HandleGroups(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetGroup(w, r)
	case http.MethodPost:
		h.CreateGroup(w, r)
	case http.MethodPut:
		h.AddContactToGroup(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *GroupHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	id := uuid.New().String()
	ctx := context.WithValue(context.Background(), "ID", id)

	name := r.FormValue("name")

	err := h.usecase.CreateGroup(ctx, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *GroupHandler) GetGroup(w http.ResponseWriter, r *http.Request) {
	id := uuid.New().String()
	ctx := context.WithValue(context.Background(), "ID", id)

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

func (h *GroupHandler) AddContactToGroup(w http.ResponseWriter, r *http.Request) {
	id := uuid.New().String()
	ctx := context.WithValue(context.Background(), "ID", id)

	contactIDStr := r.FormValue("contactID")
	contactID, err := strconv.Atoi(contactIDStr)
	if err != nil {
		http.Error(w, "invalid contact ID", http.StatusBadRequest)
		return
	}

	groupIDStr := r.FormValue("groupID")
	groupID, err := strconv.Atoi(groupIDStr)
	if err != nil {
		http.Error(w, "invalid group ID", http.StatusBadRequest)
		return
	}

	err = h.usecase.AddContactToGroup(ctx, contactID, groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
