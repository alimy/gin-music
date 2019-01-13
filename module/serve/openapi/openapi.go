package openapi

import (
	"github.com/alimy/gin-music/api"
	"github.com/alimy/gin-music/api/v1"
	"github.com/gin-gonic/gin"
)

// Context alias gin.Context type
type Context = *gin.Context

// RegisterApi setup v1 to global
func RegisterApi() {
	// Setup operations
	operations := api.OperationSlice{
		v1.ApiGetAppInfo.Get(getAppInfo),
		v1.ApiGetAlbums.Get(getAlbums),
		v1.ApiCreateAlbums.Put(createAlbums),
		v1.ApiUpdateAlbums.Post(updateAlbums),
		v1.ApiGetAlbumsById.Get(getAlbumsById),
		v1.ApiDeleteAlbumsById.Delete(deleteAlbumsById),
	}

	// Register v1.MusicGroup api to global
	api.Group(v1.Version, v1.MusicGroup).Register(operations...)

	// Register portal operations if need
	if portalOpts := portalOperations(); len(portalOpts) != 0 {
		api.Group(v1.Version).Register(portalOpts...)
	}
}
