package httphandler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewAppHandler(t *testing.T) {
	dep := NewAppHandler()
	assert.NotNil(t, dep)
}

func TestAppHandler_RouteNotFound(t *testing.T) {
	dep := NewAppHandler()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	dep.RouteNotFound(c)

	assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	assert.Equal(t, "{\"message\":\"Route not found\"}", w.Body.String())
}

func TestAppHandler_Index(t *testing.T) {
	dep := NewAppHandler()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	dep.Index(c)

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, `{"message":"this is the job application exercise!"}`, w.Body.String())
}
