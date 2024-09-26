package models

import "time"

type Game struct {
	ID            int       `db:"id"`
	HomeTeamID    int       `db:"home_team_id"`
	AwayTeamID    int       `db:"away_team_id"`
	LeagueID      int       `db:"league_id"`
	MatchDate     time.Time `db:"match_date"`
	HomeTeamScore int       `db:"home_team_score"`
	AwayTeamScore int       `db:"away_team_score"`
	CreatedAt     time.Time `db:"created_at"`
}
