package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arianaw15/birdie-talk/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("GetUserByEmail - failure", func(t *testing.T) {
		// Invalid payload, missing email
		payload := types.RegisterUserPayload{
			FirstName: "Test",
			LastName:  "User",
			Email:     "",
			Password:  "password123",
		}
		marshalledPayload, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalledPayload))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.Register)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
	t.Run("GetUserByEmail - valid, user does not exist", func(t *testing.T) {
		// Invalid payload, missing email
		payload := types.RegisterUserPayload{
			FirstName: "Test",
			LastName:  "User",
			Email:     "test@test.com",
			Password:  "password123",
		}
		marshalledPayload, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalledPayload))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.Register)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return &types.User{
		ID:        1,
		FirstName: "Test",
		LastName:  "User",
		Email:     email,
		Password:  "hashedpassword",
	}, fmt.Errorf("user with email %s not found", email)
}
func (m *mockUserStore) GetUserById(userId int) (*types.User, error) {
	return &types.User{
		ID:        userId,
		FirstName: "Test",
		LastName:  "User",
		Email:     "test@email.com",
		Password:  "hashedpassword",
	}, nil
}
func (m *mockUserStore) CreateUser(user *types.User) error {
	return nil
}
