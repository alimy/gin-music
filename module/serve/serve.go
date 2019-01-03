package serve

import (
	"github.com/alimy/gin-music/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/unisx/logus"
	"net/http"
)

func init() {
	api.Register(api.ApiGetAlbums, getAlbums)
	api.Register(api.ApiCreateAlbums, createAlbums)
	api.Register(api.ApiUpdateAlbums, updateAlbums)
	api.Register(api.ApiGetAlbumsById, getAlbumsById)
	api.Register(api.ApiDeleteAlbumsById, deleteAlbumsById)
}

func getAlbums(context *gin.Context) {
	// TODO
	logus.Debug("get albums")
	context.String(http.StatusOK, "get albums")
}

func createAlbums(context *gin.Context) {
	// TODO
	logus.Debug("create albums")
	context.String(http.StatusCreated, "Albums item created")
}

func updateAlbums(context *gin.Context) {
	// TODO
	logus.Debug("update albums")
	context.String(http.StatusCreated, "Albums item updated")
}

func getAlbumsById(context *gin.Context) {
	// TODO
	albumId := context.Param("albumId")
	logus.Debug("get albums by id", logus.String("albumId", albumId))
	context.String(http.StatusOK, "get albums by id")
}

func deleteAlbumsById(context *gin.Context) {
	// TODO
	albumId := context.Param("albumId")
	logus.Info("delete albums", logus.String("albumId", albumId))
	context.String(http.StatusOK, "Albums item deleted")
}
