package migrations

import "final-project/app/model"

// "api/app/model"

// ModelMigrations models to automigrate
var ModelMigrations = []interface{}{
	&model.Artist{},
	&model.Song{},
	&model.Gendre{},
	&model.Album{},
	&model.PlaylistType{},
	&model.PlayList{},
	// &model.User{},
}
