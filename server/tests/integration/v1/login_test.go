package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/asxcandrew/herd/server/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func LoginTest(t *testing.T) {
	r := gin.Default()
	routes.ApiRoutes(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/p/sign_in", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
