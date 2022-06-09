package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

type AlbumApi struct {
	AlbumName   *string      `json:"album_name,omitempty"`
	RealeseDate *strfmt.Date `json:"realese_date-date,omitempty" gorm:"type:timestamptz" format:"date" swaggertype:"string"`
	ArtistID    *uuid.UUID   `json:"artist_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"` // model ID
	GendreID    *uuid.UUID   `json:"gendre_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"` // model ID

}
