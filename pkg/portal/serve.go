package portal

import (
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InstallWith(r gin.IRouter) {
	handler := createStaticHandler("/",
		&assetfs.AssetFS{
			Asset:     Asset,
			AssetDir:  AssetDir,
			AssetInfo: AssetInfo})

	r.GET("/", handler)
	r.HEAD("/", handler)
}

func createStaticHandler(path string, fs http.FileSystem) gin.HandlerFunc {
	fileServer := http.StripPrefix(path, http.FileServer(fs))
	return func(c *gin.Context) {
		fileServer.ServeHTTP(c.Writer, c.Request)
	}
}
