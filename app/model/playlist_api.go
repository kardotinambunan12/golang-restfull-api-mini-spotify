package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

type PlayListApi struct {
	PlaylistName    *string      `json:"playlist_name,omitempty"`
	PlaylistDesc    *string      `json:"playlist_desc,omitempty"`
	PlaylistPicture *string      `json:"playlist_pic,omitempty"`
	CreationDate    *strfmt.Date `json:"creation_date,omitempty" gorm:"type:timestamptz" format:"date" swaggertype:"string"`
	PlaylistTypeID  *uuid.UUID   `json:"playlist_type_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`
}
