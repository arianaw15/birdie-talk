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
	router.HandleFunc("/birds/initial", h.CreateInitialBirdList).Methods("POST")
	router.HandleFunc("/birds/{id}", h.GetBirdById).Methods("GET")
}

func (h *Handler) CreateBird(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateBirdPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to parse request body: %v", err))
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	_, err := h.store.GetBirdByName(payload.CommonName)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to get bird by common name: %v", err))
	}

	err = h.store.CreateBird(&types.Bird{
		CommonName:     payload.CommonName,
		ScientificName: payload.ScientificName,
		Description:    payload.Description,
		ImageURL:       payload.ImageURL,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "bird created successfully"})
}

// CreateInitialBirdList creates a list of birds from the provided payload. Data can be found in data/initial-data.json file
func (h *Handler) CreateInitialBirdList(w http.ResponseWriter, r *http.Request) {
	var payload = []types.CreateBirdPayload{}
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to parse request body: %v", err))
		return
	}

	// loop through payload
	for _, bird := range payload {
		// validate payload
		if err := utils.Validate.Struct(bird); err != nil {
			errors := err.(validator.ValidationErrors)
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
			return
		}

		existingBird, _ := h.store.GetBirdByName(bird.CommonName)

		if existingBird != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf(" skipping bird with common name %s as already exists", bird.CommonName))
		} else {
			err := h.store.CreateBird(&types.Bird{
				CommonName:     bird.CommonName,
				ScientificName: bird.ScientificName,
				Description:    bird.Description,
				ImageURL:       bird.ImageURL,
			})
			if err != nil {
				utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to create bird: %v", err))
			}
		}

	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "initial bird list created successfully"})
}

func (h *Handler) GetBirdById(w http.ResponseWriter, r *http.Request) {
	var id int

	bird, err := h.store.GetBirdById(id)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("bird with id %d not found: %v", id, err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, bird)

}
