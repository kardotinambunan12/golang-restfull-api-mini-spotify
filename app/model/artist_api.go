package model

type ArtistApi struct {
	Name  *string `json:"name,omitempty"`
	About *string `json:"about,omitempty"`
}
