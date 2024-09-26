package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"scraper/internal/models"

	"github.com/jmoiron/sqlx"
)

type LeagueRepository struct {
	DB *sqlx.DB
}

// NewLeagueRepository creates a new instance of LeagueRepository
func NewLeagueRepository(db *sqlx.DB) *LeagueRepository {
	return &LeagueRepository{DB: db}
}

// FindByID retrieves a league by its ID
func (r *LeagueRepository) FindByID(ctx context.Context, id int) (*models.League, error) {
	query := `SELECT id, name, country, created_at FROM leagues WHERE id = $1`
	var league models.League
	err := r.DB.GetContext(ctx, &league, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no league found with id: %d", id)
		}
		return nil, err
	}
	return &league, nil
}

// Insert adds a new league to the database
func (r *LeagueRepository) Insert(ctx context.Context, league *models.League) error {
	query := `INSERT INTO leagues (name, country, created_at) VALUES ($1, $2, $3) RETURNING id`
	err := r.DB.QueryRowContext(ctx, query, league.Name, league.Country, league.CreatedAt).Scan(&league.ID)
	if err != nil {
		return err
	}
	return nil
}
