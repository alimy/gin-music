package openapi

import (
	"github.com/alimy/gin-music/models"
	"github.com/alimy/gin-music/models/core"
	"github.com/unisx/logus"
	"net/http"
)

func getAppInfo(c Context) {
	// TODO
	logus.Debug("get application information")
	c.String(http.StatusOK, "get application information")
}

func getAlbums(c Context) {
	if albums, ok := core.Model(models.IdAlbums).(*models.Albums); ok {
		logus.Debug("getAlbums")
		if err := core.Retrieve(models.IdAlbums, albums); err == nil {
			c.JSON(http.StatusOK, albums)
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	} else {
		c.String(http.StatusNotFound, "get albums")
	}
}

func createAlbums(c Context) {
	// TODO
	logus.Debug("create albums")
	c.String(http.StatusCreated, "Albums item created")
}

func updateAlbums(c Context) {
	if album, ok := core.Model(models.IdAlbum).(*models.Album); ok {
		logus.Debug("updateAlbums")
		album.Id = c.Param("albumId")
		if err := core.Update(models.IdAlbum, album); err == nil {
			c.String(http.StatusCreated, "Albums item updated")
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}

	} else {
		c.String(http.StatusNotFound, "update albums failure")
	}
}

func getAlbumsById(c Context) {
	if album, ok := core.Model(models.IdAlbum).(*models.Album); ok {
		logus.Debug("getAlbumsById")
		album.Id = c.Param("albumId")
		if err := core.Retrieve(models.IdAlbum, album); err == nil {
			c.JSON(http.StatusOK, album)
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	} else {
		c.String(http.StatusNotFound, "update albums failure")
	}
}

func deleteAlbumsById(c Context) {
	if album, ok := core.Model(models.IdAlbum).(*models.Album); ok {
		logus.Debug("deleteAlbumsById")
		album.Id = c.Param("albumId")
		if err := core.Delete(models.IdAlbum, album); err == nil {
			c.String(http.StatusOK, "Albums item deleted")
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	} else {
		c.String(http.StatusNotFound, "delete album failure")
	}
}
