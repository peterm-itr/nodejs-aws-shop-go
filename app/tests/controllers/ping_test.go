package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/peterm-itr/nodejs-aws-shop-go/server"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := server.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
