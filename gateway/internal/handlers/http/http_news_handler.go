package http

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"

	"gateway/internal/services"
)

type NewsHandler struct {
	newsSvc services.NewsService
}

func NewNewsHandler(newsSvc services.NewsService) *NewsHandler {
	return &NewsHandler{
		newsSvc: newsSvc,
	}
}

func (h *NewsHandler) GetRoutes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", h.get)

	return router
}

func (h *NewsHandler) get(writer http.ResponseWriter, request *http.Request) {
	users, err := h.newsSvc.Get(request.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	responseBytes, err := json.Marshal(users)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Content-Size", string(len(responseBytes)))
	writer.Write(responseBytes)
}
