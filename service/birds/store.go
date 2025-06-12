package birds

import (
	"database/sql"
	"fmt"

	"github.com/arianaw15/birdie-talk/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetBirdById(birdId int) (*types.Bird, error) {
	query := "SELECT * FROM birds WHERE birdId = ?"
	rows, err := s.db.Query(query, birdId)
	if err != nil {
		return nil, err
	}

	bird := new(types.Bird)
	for rows.Next() {
		bird, err = scanRowIntoBird(rows)
		if err != nil {
			return nil, err
		}
	}
	if bird.ID == 0 {
		return nil, fmt.Errorf("bird with birdId %v not found", birdId)
	}
	return bird, nil
}

func (s *Store) GetBirdByName(name string) (*types.Bird, error) {
	query := "SELECT * FROM birds WHERE name = ?"
	rows, err := s.db.Query(query, name)
	if err != nil {
		return nil, err
	}

	bird := new(types.Bird)
	for rows.Next() {
		bird, err = scanRowIntoBird(rows)
		if err != nil {
			return nil, err
		}
	}
	if bird.ID == 0 {
		return nil, fmt.Errorf("bird with name %v not found", name)
	}
	return bird, nil
}

func (s *Store) CreateBird(bird *types.Bird) error {
	_, err := s.db.Exec("INSERT INTO birds (name, species, description, imageUrl) VALUES (?, ?, ?, ?)", bird.Name, bird.Species, bird.Description, bird.ImageURL)
	if err != nil {
		return fmt.Errorf("failed to create bird: %w", err)
	}
	return nil
}

func scanRowIntoBird(rows *sql.Rows) (*types.Bird, error) {
	bird := new(types.Bird)
	err := rows.Scan(&bird.ID, &bird.Name, &bird.Species, &bird.Description, &bird.ImageURL, &bird.CreatedAt)
	if err != nil {
		return nil, err
	}
	return bird, nil
}
