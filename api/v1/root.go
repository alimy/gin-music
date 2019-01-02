package api

import (
	"github.com/gin-gonic/gin"
)

var (
	// OperationIds for api handler
	ApiGetAlbums        = apiGet("albums")             // get albums
	ApiCreateAlbums     = apiPut("albums")             // create an albums
	ApiUpdateAlbums     = apiPost("albums")            // update an albums
	ApiGetAlbumsById    = apiGet("albums/:albumId")    // Get album by Id
	ApiDeleteAlbumsById = apiDelete("albums/:albumId") // Delete an albums by Id
)

type Operation struct {
	Method string
	Path   string
}

type OperationIds map[Operation]func(ctx *gin.Context)

func InstallWith(r gin.IRouter, ids OperationIds) {
	r = r.Group(ApiVersion)

	for opt, handler := range ids {
		r.Handle(opt.Method, opt.Path, handler)
	}
}

func apiGet(path string) Operation {
	return Operation{Method: "GET", Path: path}
}

func apiPut(path string) Operation {
	return Operation{Method: "PUT", Path: path}
}

func apiPost(path string) Operation {
	return Operation{Method: "POST", Path: path}
}

func apiDelete(path string) Operation {
	return Operation{Method: "DELETE", Path: path}
}
