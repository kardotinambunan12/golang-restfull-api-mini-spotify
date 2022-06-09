package model

import "github.com/google/uuid"

type SongApi struct {
	ArtistID    *uuid.UUID `json:"artist_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"` // model ID
	Name        *string    `json:"name,omitempty"`
	Gendre      *string    `json:"gendre,omitempty"`
	Year        *int       `json:"year,omitempty"`
	Description *string    `json:"description,omitempty"`
}
