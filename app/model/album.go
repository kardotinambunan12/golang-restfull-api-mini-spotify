package model

type Album struct {
	Base
	AlbumApi
	Gendre *Gendre `json:"gendre,omitempty" gorm:"foreignKey:GendreID;references:ID"`
	Artist *Artist `json:"artist,omitempty" gorm:"foreignKey:ArtistID;references:ID"`
}
