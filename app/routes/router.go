package routes

import (
	"final-project/app/controller"
	"final-project/app/controller/album"
	"final-project/app/controller/artist"
	"final-project/app/controller/gendre"
	"final-project/app/controller/playlist"
	"final-project/app/controller/playlist_type"
	"final-project/app/controller/song"
	"final-project/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

// Handle all request to route to controller
func Handle(app *fiber.App) {
	app.Use(cors.New())
	services.InitDatabase()
	// services.InitRedis()

	api := app.Group(viper.GetString("ENDPOINT"))

	api.Get("/", controller.GetAPIIndex)
	api.Get("/info.json", controller.GetAPIInfo)

	//routes artist endpoint
	api.Get("/artist", artist.ArtistGet)
	api.Get("/artist/:id", artist.ArtistGetById)
	api.Post("/artist", artist.ArtistPost)
	api.Put("/artist/:id", artist.ArtistPut)
	api.Delete("/artist/:id", artist.ArtistDeleted)

	//routes song endpoint
	api.Get("/song", song.SongGet)
	api.Get("/song/:id", song.SongGetById)
	api.Post("/song", song.SongPost)
	api.Put("/song/:id", song.SongPut)
	api.Delete("/song/:id", song.SongDelete)

	// routes gendre endpoint
	api.Get("/gendre", gendre.GendreGetAll)
	api.Get("/gendre/:id", gendre.GendreGetById)
	api.Post("/gendre", gendre.GendrePost)
	api.Put("/gendre/:id", gendre.GendrePut)
	api.Delete("/gendre/:id", gendre.GendreDelete)

	//routes album endpoint
	api.Get("/album", album.AlbumGetAll)
	api.Get("/album/:id", album.AlbumGetById)
	api.Post("/album", album.AlbumPost)
	api.Put("/album/:id", album.AlbumPut)
	api.Delete("/album/:id", album.AlbumDelete)

	//Routes playlist Type endpoint

	api.Get("/playlist_type", playlist_type.PlaylistTypeGet)
	api.Get("/playlist_type/:id", playlist_type.PlaylistTypeGetById)
	api.Post("/playlist_type", playlist_type.PlaylistTypePost)
	api.Put("/playlist_type/:id", playlist_type.PlaylistTypePut)
	api.Delete("/playlist_type/:id", playlist_type.PlaylistTypeDelete)

	//routes playlist endpoint
	api.Get("/playlist", playlist.PlaylistGetAll)
	api.Get("/playlist/:id", playlist.PlaylistGetById)
	api.Post("/playlist", playlist.PlaylistPost)
	api.Put("/playlist/:id", playlist.PlaylistPut)
	api.Delete("/playlist/:id", playlist.PlaylistDelete)

}
