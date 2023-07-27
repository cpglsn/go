package main

import (
	"time"

	"github.com/cpglsn/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json="id"`
	CreatedAt time.Time `json="created_at"`
	UpdatedAt time.Time `json="updated_at"`
	Name      string    `json="name"`
	APIKey    string    `json="api_key"`
}

func userToUser(u database.User) User {
	return User{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Name:      u.Name,
		APIKey:    u.ApiKey,
	}
}
