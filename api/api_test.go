package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	testIndex ApiId = "/index/"
	testGroup       = "test"
	testV1          = "v1"
	testV2          = "v2"
)

func testHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ping")
}

func TestApiId(t *testing.T) {
	operationGet := testIndex.Get(testHandler)
	assert.Equal(t, string(testIndex), operationGet.Path)
	assert.Equal(t, http.MethodGet, operationGet.Method)
	assert.NotNil(t, operationGet.Handler)

	operationPut := testIndex.Put(testHandler)
	assert.Equal(t, string(testIndex), operationPut.Path)
	assert.Equal(t, http.MethodPut, operationPut.Method)
	assert.NotNil(t, operationPut.Handler)

	operationPost := testIndex.Post(testHandler)
	assert.Equal(t, string(testIndex), operationPost.Path)
	assert.Equal(t, http.MethodPost, operationPost.Method)
	assert.NotNil(t, operationPost.Handler)

	operationDelete := testIndex.Delete(testHandler)
	assert.Equal(t, string(testIndex), operationDelete.Path)
	assert.Equal(t, http.MethodDelete, operationDelete.Method)
	assert.NotNil(t, operationDelete.Handler)

	operationHead := testIndex.Head(testHandler)
	assert.Equal(t, string(testIndex), operationHead.Path)
	assert.Equal(t, http.MethodHead, operationHead.Method)
	assert.NotNil(t, operationHead.Handler)
}

func TestGroup(t *testing.T) {
	operations := Group(testV1, testGroup)
	assert.NotNil(t, operations)
	RecycleAll()
}

func TestRecycle(t *testing.T) {
	Group(testV1, testGroup)
	assert.NotNil(t, operationGroupVs[testV1])
	Recycle(testV1)
	assert.Nil(t, operationGroupVs[testV1])
}

func TestRecycleAll(t *testing.T) {
	operations1 := Group(testV1, testGroup)
	operations2 := Group(testV2, testGroup)
	assert.Equal(t, operations1, operationGroupVs[testV1].Group(testGroup))
	assert.Equal(t, operations2, operationGroupVs[testV2].Group(testGroup))
	RecycleAll()
	assert.Equal(t, 0, len(operationGroupVs))
}

func TestOperationSlice_Register(t *testing.T) {
	operation := testIndex.Get(testHandler)
	Group(testV1, testGroup).Register(operation)
	operations := operationGroupVs[testV1].Group(testGroup)
	assert.Equal(t, operation, (*operations)[0])
	RecycleAll()
}

func TestInstallDefault(t *testing.T) {
	// TODO
}

func TestInstallWith(t *testing.T) {
	// TODO
}
