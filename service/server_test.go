package service

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWhenValidRequestEvalReturns200(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/evaluate", strings.NewReader("{\"expression\":\"What is 5?\"}"))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"result\":\"5.00\"}", w.Body.String())
}

func TestWhenValidRequestValidateReturns200(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/validate", strings.NewReader("{\"expression\":\"What is 5?\"}"))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"valid\":\"true\"}", w.Body.String())
}

func TestInvalidEndpointReturns404(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/other", strings.NewReader("my request"))
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}
