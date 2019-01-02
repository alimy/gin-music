package serve

import (
	"github.com/gin-gonic/gin"
	"github.com/unisx/logus"
	"net/http"
)

func getAlbums(ctx *gin.Context) {
	// TODO
	logus.Debug("get albums")
	ctx.String(http.StatusOK, "Albums")
}

func createAlbums(ctx *gin.Context) {
	// TODO
	logus.Debug("create albums")
	ctx.String(http.StatusCreated, "Albums item to update")
}

func updateAlbums(ctx *gin.Context) {
	// TODO
	logus.Debug("update albums")
	ctx.String(http.StatusCreated, "Albums item to update")
}

func getAlbumsById(ctx *gin.Context) {
	// TODO
	albumId := ctx.Param("albumId")
	logus.Debug("get albums by id", logus.String("albumId", albumId))
	ctx.String(http.StatusOK, "albums")
}

func deleteAlbumsById(ctx *gin.Context) {
	// TODO
	albumId := ctx.Param("albumId")
	logus.Info("delete albums", logus.String("albumId", albumId))
	ctx.String(http.StatusOK, "Item deleted")
}
