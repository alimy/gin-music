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
		var r gin.IRouter
		if group == "" {
			r = router
		} else {
			r = router.Group(group)
		}
		for _, operation := range OperationIds {
			if operation.Group == group && operation.Handler != nil {
				r.Handle(operation.Method, operation.Path, operation.Handler)
			}
		}
	}
}

// build GET operation
func apiGet(path string, groups ...string) *Operation {
	if len(groups) == 1 {
		return &Operation{Group: groups[0], Path: path, Method: http.MethodGet}
	}
	return &Operation{Path: path, Method: http.MethodGet}
}

// build GET operation
func apiPut(path string, groups ...string) *Operation {
	if len(groups) == 1 {
		return &Operation{Group: groups[0], Path: path, Method: http.MethodPut}
	}
	return &Operation{Path: path, Method: http.MethodPut}
}

// build POST operation
func apiPost(path string, groups ...string) *Operation {
	if len(groups) == 1 {
		return &Operation{Group: groups[0], Path: path, Method: http.MethodPost}
	}
	return &Operation{Path: path, Method: http.MethodPost}
}

// build DELETE operation
func apiDelete(path string, groups ...string) *Operation {
	if len(groups) == 1 {
		return &Operation{Group: groups[0], Path: path, Method: http.MethodDelete}
	}
	return &Operation{Path: path, Method: http.MethodDelete}
}

// build HEAD operation
func apiHead(path string, groups ...string) *Operation {
	if len(groups) == 1 {
		return &Operation{Group: groups[0], Path: path, Method: http.MethodHead}
	}
	return &Operation{Path: path, Method: http.MethodHead}
}
