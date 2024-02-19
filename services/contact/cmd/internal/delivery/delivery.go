package delivery

import (
	"architecture_go/services/contact/cmd/internal/usecase"
	"net/http"
)

type ContactHandler struct {
	usecase usecase.ContactUsecase
}

func NewContactHandler(usecase *usecase.ContactUsecase) *ContactHandler {
	return &ContactHandler{usecase: *usecase}
}

func (h *ContactHandler) HandleContacts(w http.ResponseWriter, r *http.Request) {
	// Handle contacts logic here
}
