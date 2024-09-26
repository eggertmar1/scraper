package models

import "time"

type Team struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	LeagueID  int       `db:"league_id"`
	CreatedAt time.Time `db:"created_at"`
}
