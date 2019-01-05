// +build portal

package openapi

import (
	"github.com/alimy/gin-music/api/v1"
	"github.com/alimy/gin-music/pkg/portal"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	assetFS := &assetfs.AssetFS{
		Asset:     portal.Asset,
		AssetDir:  portal.AssetDir,
		AssetInfo: portal.AssetInfo}

	mainHandler := createStaticHandler("/", assetFS)
	staticHandler := createStaticHandler("/static", assetFS)

	api.Register(api.ApiGetMainPage, mainHandler)
	api.Register(api.ApiHeadMainPage, mainHandler)
	api.Register(api.ApiGetStaticAssets, staticHandler)
	api.Register(api.ApiHeadStaticAssets, staticHandler)
}

func createStaticHandler(path string, fs http.FileSystem) gin.HandlerFunc {
	fileServer := http.StripPrefix(path, http.FileServer(fs))
	return func(c *gin.Context) {
		fileServer.ServeHTTP(c.Writer, c.Request)
	}
}
