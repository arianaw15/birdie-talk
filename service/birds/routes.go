package birds

import (
	"fmt"
	"net/http"

	"github.com/arianaw15/birdie-talk/types"
	"github.com/arianaw15/birdie-talk/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.BirdStore
}

func NewHandler(store types.BirdStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/birds/create", h.CreateBird).Methods("POST")
}

func (h *Handler) CreateBird(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterBirdPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	_, err := h.store.GetBirdByName(payload.Name)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	err = h.store.CreateBird(&types.Bird{
		Name:        payload.Name,
		Species:     payload.Species,
		Description: payload.Description,
		ImageURL:    payload.ImageURL,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "bird created successfully"})
}
