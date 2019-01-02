package serve

import (
	"github.com/gin-gonic/gin"
	"github.com/unisx/logus"
	"net/http"
)

func getAlbums(context *gin.Context) {
	// TODO
	logus.Debug("get albums")
	context.String(http.StatusOK, "Albums")
}

func createAlbums(context *gin.Context) {
	// TODO
	logus.Debug("create albums")
	context.String(http.StatusCreated, "Albums item to update")
}

func updateAlbums(context *gin.Context) {
	// TODO
	logus.Debug("update albums")
	context.String(http.StatusCreated, "Albums item to update")
}

func getAlbumsById(context *gin.Context) {
	// TODO
	albumId := context.Param("albumId")
	logus.Debug("get albums by id", logus.String("albumId", albumId))
	context.String(http.StatusOK, "albums")
}

func deleteAlbumsById(context *gin.Context) {
	// TODO
	albumId := context.Param("albumId")
	logus.Info("delete albums", logus.String("albumId", albumId))
	context.String(http.StatusOK, "Item deleted")
}
