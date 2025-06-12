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
	query := "SELECT * FROM birds WHERE id = ?"
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
		return nil, fmt.Errorf("bird with id %v not found", birdId)
	}
	return bird, nil
}

func (s *Store) GetBirdByName(name string) (*types.Bird, error) {
	query := "SELECT * FROM birds WHERE commonName = ?"
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
	_, err := s.db.Exec("INSERT INTO birds (commonName, scientificName, description, imageUrl) VALUES (?, ?, ?, ?)", bird.CommonName, bird.ScientificName, bird.Description, bird.ImageURL)
	if err != nil {
		return fmt.Errorf("failed to create bird: %w", err)
	}
	return nil
}

func (s *Store) CreateInitialBirdList(birds []types.Bird) error {
	// loop through payload and create each bird
	for _, bird := range birds {
		_, err := s.db.Exec("INSERT INTO birds (commonName, scientificName, description, imageUrl) VALUES (?, ?, ?, ?)", bird.CommonName, bird.ScientificName, bird.Description, bird.ImageURL)
		if err != nil {
			return fmt.Errorf("failed to create bird: %w", err)
		}
	}
	return nil
}

func scanRowIntoBird(rows *sql.Rows) (*types.Bird, error) {
	bird := new(types.Bird)
	err := rows.Scan(&bird.ID, &bird.CommonName, &bird.ScientificName, &bird.Description, &bird.ImageURL, &bird.CreatedAt)
	if err != nil {
		return nil, err
	}
	return bird, nil
}
