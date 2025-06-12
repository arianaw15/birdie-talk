package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(userId int) (*User, error)
	CreateUser(user *User) error
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=64"`
}
type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserBirdList struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	Email     string `json:"email"`
	Birds     []Bird `json:"birds"`
}

type BirdStore interface {
	GetBirdById(birdId int) (*Bird, error)
	GetBirdByName(name string) (*Bird, error)
	CreateBird(bird *Bird) error
}

type Bird struct {
	ID             int       `json:"id"`
	CommonName     string    `json:"commonName"`
	ScientificName string    `json:"scientificName"`
	Description    string    `json:"description"`
	ImageURL       string    `json:"imageUrl"`
	CreatedAt      time.Time `json:"createdAt"`
}

type RegisterBirdPayload struct {
	CommonName     string `json:"commonName"`
	ScientificName string `json:"scientificName"`
	Description    string `json:"description"`
	ImageURL       string `json:"imageUrl"`
}
