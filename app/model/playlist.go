package model

type PlayList struct {
	Base
	PlayListApi
	PlaylistType *PlaylistType `json:"playlist_type,omitempty" gorm:"foreignKey:PlaylistTypeID;references:ID"`
}
