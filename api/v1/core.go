package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Operation is the Api handler info
type Operation struct {
	Group   string          // api's url prefix
	Path    string          // api's url relative path
	Method  string          // api's http method
	Handler gin.HandlerFunc // api's handler
}

// Register add id-handler map for OperationIds
func Register(apiId int, handlerFunc gin.HandlerFunc) {
	if operation, ok := OperationIds[apiId]; ok {
		operation.Handler = handlerFunc
	}
}

// InstallDefault install router to all operation in OperationIds
func InstallDefault(router gin.IRouter) {
	groups := make(map[string]bool)
	for _, operation := range OperationIds {
		if !groups[operation.Group] {
			groups[operation.Group] = true
		}
	}
	groupSlice := make([]string, 0, len(groups))
	for g := range groups {
		groupSlice = append(groupSlice, g)
	}
	InstallWith(router, groupSlice...)
}

// InstallWith install router to give groups's operation in OperationIds
func InstallWith(router gin.IRouter, groups ...string) {
	for _, group := range groups {
		r := router.Group(group)
		for _, operation := range OperationIds {
			if operation.Group == group && operation.Handler != nil {
				r.Handle(operation.Method, operation.Path, operation.Handler)
			}
		}
	}
}

func apiGet(group, path string) *Operation {
	return &Operation{Group: group, Path: path, Method: http.MethodGet}
}

func apiPut(group, path string) *Operation {
	return &Operation{Group: group, Path: path, Method: http.MethodPut}
}

func apiPost(group, path string) *Operation {
	return &Operation{Group: group, Path: path, Method: http.MethodPost}
}

func apiDelete(group, path string) *Operation {
	return &Operation{Group: group, Path: path, Method: http.MethodDelete}
}

func apiHead(group, path string) *Operation {
	return &Operation{Group: group, Path: path, Method: http.MethodHead}
}