package model

type Song struct {
	Base
	SongApi
	Artist *Artist `json:"artist" gorm:"foreignKey:ArtistID;references:ID"`
}
