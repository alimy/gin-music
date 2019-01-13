// +build portal

package openapi

import (
	"github.com/alimy/gin-music/api"
	"github.com/alimy/gin-music/api/v1"
	"github.com/alimy/gin-music/pkg/portal"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func portalOperations() api.OperationSlice {
	assetFS := &assetfs.AssetFS{
		Asset:     portal.Asset,
		AssetDir:  portal.AssetDir,
		AssetInfo: portal.AssetInfo}

	mainHandler := createStaticHandler("/", assetFS)
	staticHandler := createStaticHandler("/static", assetFS)

	return api.OperationSlice{
		v1.ApiGetMainPage.Get(mainHandler),
		v1.ApiHeadMainPage.Head(mainHandler),
		v1.ApiGetStaticAssets.Get(staticHandler),
		v1.ApiHeadStaticAssets.Head(staticHandler),
	}
}

func createStaticHandler(path string, fs http.FileSystem) gin.HandlerFunc {
	fileServer := http.StripPrefix(path, http.FileServer(fs))
	return func(c Context) {
		fileServer.ServeHTTP(c.Writer, c.Request)
	}
}
