package main

import (
"net/http"
"net/http/httptest"
	"strings"
	"testing"
"github.com/stretchr/testify/assert"
)

func TestWhenValidRequestEvalReturns200(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/evaluate", strings.NewReader("my request"))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "my request", w.Body.String())
}

func TestInvalidEndpointReturns404(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/other", strings.NewReader("my request"))
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}