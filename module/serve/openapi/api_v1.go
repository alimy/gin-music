package openapi

import (
	"github.com/alimy/gin-music/api/v1"
	"github.com/alimy/gin-music/models"
	"github.com/alimy/gin-music/models/core"
	"github.com/gin-gonic/gin"
	"github.com/unisx/logus"
	"net/http"
)

func init() {
	api.Register(api.ApiGetAppInfo, getAppInfo)
	api.Register(api.ApiGetAlbums, getAlbums)
	api.Register(api.ApiCreateAlbums, createAlbums)
	api.Register(api.ApiUpdateAlbums, updateAlbums)
	api.Register(api.ApiGetAlbumsById, getAlbumsById)
	api.Register(api.ApiDeleteAlbumsById, deleteAlbumsById)
}

func getAppInfo(context *gin.Context) {
	// TODO
	logus.Debug("get application information")
	context.String(http.StatusOK, "get application information")
}

func getAlbums(context *gin.Context) {
	if albums, ok := core.Model(models.IdAlbums).(*models.Albums); ok {
		logus.Debug("getAlbums")
		core.Retrieve(models.IdAlbums, albums)
		context.JSON(http.StatusOK, albums)
	} else {
		context.String(http.StatusNotFound, "get albums")
	}
}

func createAlbums(context *gin.Context) {
	// TODO
	logus.Debug("create albums")
	context.String(http.StatusCreated, "Albums item created")
}

func updateAlbums(context *gin.Context) {
	if album, ok := core.Model(models.IdAlbum).(*models.Album); ok {
		logus.Debug("updateAlbums")
		album.Id = context.Param("albumId")
		core.Update(models.IdAlbum, album)
		context.String(http.StatusCreated, "Albums item updated")
	} else {
		context.String(http.StatusNotFound, "update albums failure")
	}
}

func getAlbumsById(context *gin.Context) {
	if album, ok := core.Model(models.IdAlbum).(*models.Album); ok {
		logus.Debug("getAlbumsById")
		album.Id = context.Param("albumId")
		core.Retrieve(models.IdAlbum, album)
		context.JSON(http.StatusOK, album)
	} else {
		context.String(http.StatusNotFound, "update albums failure")
	}
}

func deleteAlbumsById(context *gin.Context) {
	if album, ok := core.Model(models.IdAlbum).(*models.Album); ok {
		logus.Debug("deleteAlbumsById")
		album.Id = context.Param("albumId")
		core.Delete(models.IdAlbum, album)
		context.String(http.StatusOK, "Albums item deleted")
	} else {
		context.String(http.StatusNotFound, "delete album failure")
	}
}
