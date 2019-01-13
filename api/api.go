package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	// DefaultGroup indicate default group
	DefaultGroup = "/"
)

// ApiId API ID type
type ApiId string

// Operation is the Api handler info
type Operation struct {
	Path    string          // api's url relative path
	Method  string          // api's http method
	Handler gin.HandlerFunc // api's handler
}

// OperationSlice contains ApiId-Operation items
type OperationSlice []*Operation

// OperationGroup contains Group-OperationMap items
type OperationGroup map[string]*OperationSlice

// OperationGroupVs contains Version-OperationGroup items
type OperationGroupVs map[string]OperationGroup

// HandlerMap contains ApiId-HandlerFunc items
type HandlerMap map[ApiId]gin.HandlerFunc

// operationGroupVs contains version-operationGroupV map used for install and
// register handler info
var operationGroupVs = make(OperationGroupVs)

// Group return a OperationSlice that mapped for group
func Group(version string, groups ...string) *OperationSlice {
	var (
		group          string
		operationGroup OperationGroup
	)
	if len(groups) == 0 {
		group = DefaultGroup
	} else {
		group = groups[0]
	}
	if operationGroup = operationGroupVs[version]; operationGroup == nil {
		operationGroup = make(OperationGroup)
		operationGroupVs[version] = operationGroup
	}
	return operationGroup.Group(group)
}

// InstallDefault install router to all operation in operationIds
func InstallDefault(version string, router gin.IRouter) {
	if operationGroup, ok := operationGroupVs[version]; ok {
		operationGroup.InstallDefault(router)
	}
}

// InstallWith install router to give groups's operation in operationIds
func InstallWith(version string, router gin.IRouter, groups ...string) {
	if operationGroup, ok := operationGroupVs[version]; ok {
		operationGroup.InstallWith(router, groups...)
	}
}

// Recycle delete data that map the of version's item in operationGroupVs.
// Used to gc resource after InstallDefault(...)/InstallWith(...) invoked.
func Recycle(version string) {
	if _, ok := operationGroupVs[version]; ok {
		delete(operationGroupVs, version)
	}
}

// RecycleAll delete all data in operationGroupVs.
// Used to gc resource after InstallDefault(...)/InstallWith(...) invoked.
func RecycleAll() {
	for version := range operationGroupVs {
		delete(operationGroupVs, version)
	}
}

// Register add id-handler map for OperationIds
func (m *OperationSlice) Register(operations ...*Operation) {
	*m = append(*m, operations...)
}

// Group return a OperationMap instance that mapped for group
func (g OperationGroup) Group(group string) (operations *OperationSlice) {
	if operations = g[group]; operations == nil {
		operationSlice := make(OperationSlice, 0)
		g[group] = &operationSlice
		operations = &operationSlice
	}
	return
}

// InstallDefault install router to all operation in OperationIds
func (g OperationGroup) InstallDefault(router gin.IRouter) {
	for group, operations := range g {
		g.installWith(router, group, operations)
	}
}

// InstallWith install router to give groups's operation in OperationIds
func (g OperationGroup) InstallWith(router gin.IRouter, groups ...string) {
	for _, group := range groups {
		if operations, ok := g[group]; ok {
			g.installWith(router, group, operations)
		}
	}
}

// installWith install router to give group's operation in OperationIds
func (g OperationGroup) installWith(router gin.IRouter, group string, operations *OperationSlice) {
	var r gin.IRouter
	if group == "" || group == "/" {
		r = router
	} else {
		r = router.Group(group)
	}
	for _, operation := range *operations {
		if operation.Handler != nil {
			r.Handle(operation.Method, operation.Path, operation.Handler)
		}
	}
}

// Get build GET operation
func (id ApiId) Get(handler gin.HandlerFunc) *Operation {
	return &Operation{Path: string(id), Method: http.MethodGet, Handler: handler}
}

// Put build PUT operation
func (id ApiId) Put(handler gin.HandlerFunc) *Operation {
	return &Operation{Path: string(id), Method: http.MethodPut, Handler: handler}
}

// Post build POST operation
func (id ApiId) Post(handler gin.HandlerFunc) *Operation {
	return &Operation{Path: string(id), Method: http.MethodPost, Handler: handler}
}

// Delete build DELETE operation
func (id ApiId) Delete(handler gin.HandlerFunc) *Operation {
	return &Operation{Path: string(id), Method: http.MethodDelete, Handler: handler}
}

// Head build HEAD operation
func (id ApiId) Head(handler gin.HandlerFunc) *Operation {
	return &Operation{Path: string(id), Method: http.MethodHead, Handler: handler}
}
