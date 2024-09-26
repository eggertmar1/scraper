-- V1__create_tables.sql
CREATE TABLE leagues (
     id             SERIAL PRIMARY KEY,
     name           VARCHAR(255) NOT NULL,
     country        VARCHAR(100),
     created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE teams (
   id               SERIAL PRIMARY KEY,
   name             VARCHAR(255) NOT NULL,
   league_id        INT REFERENCES leagues(id),
   created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE games (
   id               SERIAL PRIMARY KEY,
   home_team_id     INT REFERENCES teams(id),
   away_team_id     INT REFERENCES teams(id),
   league_id        INT REFERENCES leagues(id),
   match_date       TIMESTAMP,
   home_team_score  INT,
   away_team_score  INT,
   created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
