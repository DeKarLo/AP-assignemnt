// contact_handler.go
package delivery

import (
	"architecture_go/services/contact/cmd/internal/usecase"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

type ContactHandler struct {
	usecase usecase.ContactUsecase
}

func NewContactHandler(usecase usecase.ContactUsecase) *ContactHandler {
	return &ContactHandler{
		usecase: usecase,
	}
}

func (h *ContactHandler) HandleContacts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetContact(w, r)
	case http.MethodPost:
		h.CreateContact(w, r)
	case http.MethodPut:
		h.UpdateContact(w, r)
	case http.MethodDelete:
		h.DeleteContact(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ContactHandler) CreateContact(w http.ResponseWriter, r *http.Request) {
	type contextKey string

	const (
		contextKeyID contextKey = "ID"
	)

	ctx := context.WithValue(r.Context(), contextKeyID, uuid.New().String())

	firstName := r.FormValue("firstName")
	middleName := r.FormValue("middleName")
	lastName := r.FormValue("lastName")
	phoneNumberStr := r.FormValue("phoneNumber")

	phoneNumber, err := strconv.Atoi(phoneNumberStr)
	if err != nil {
		http.Error(w, "invalid phone number", http.StatusBadRequest)
		return
	}

	err = h.usecase.CreateContact(ctx, firstName, middleName, lastName, phoneNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ContactHandler) GetContact(w http.ResponseWriter, r *http.Request) {
	type contextKey string

	const (
		contextKeyID contextKey = "ID"
	)

	id := r.Context().Value(contextKeyID).(string)

	ctx := context.WithValue(r.Context(), contextKeyID, id)

	idStr := r.FormValue("id")
	contactID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid contact ID", http.StatusBadRequest)
		return
	}

	contact, err := h.usecase.GetContactByID(ctx, contactID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(contact)
}

func (h *ContactHandler) UpdateContact(w http.ResponseWriter, r *http.Request) {
	type contextKey string

	const (
		contextKeyID contextKey = "ID"
	)

	id := r.Context().Value(contextKeyID).(string)

	ctx := context.WithValue(r.Context(), contextKeyID, id)

	idStr := r.FormValue("id")
	contactID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid contact ID", http.StatusBadRequest)
		return
	}

	firstName := r.FormValue("firstName")
	middleName := r.FormValue("middleName")
	lastName := r.FormValue("lastName")
	phoneNumberStr := r.FormValue("phoneNumber")

	phoneNumber, err := strconv.Atoi(phoneNumberStr)
	if err != nil {
		http.Error(w, "invalid phone number", http.StatusBadRequest)
		return
	}

	err = h.usecase.UpdateContact(ctx, contactID, firstName, middleName, lastName, phoneNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ContactHandler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	type contextKey string

	const (
		contextKeyID contextKey = "ID"
	)

	id := r.Context().Value(contextKeyID).(string)

	ctx := context.WithValue(r.Context(), contextKeyID, id)

	idStr := r.FormValue("id")
	contactID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid contact ID", http.StatusBadRequest)
		return
	}

	err = h.usecase.DeleteContact(ctx, contactID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
