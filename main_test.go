package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	go Start()

	t.Run("it should return 200 when health is ok", func(t *testing.T) {
		resp, err := http.Get("http://localhost:8080/v1/health")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("it should return 404 for any other endpoint", func(t *testing.T) {
		resp, err := http.Get("http://localhost:8080/v1/nonexistent")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("it should return error on GET", func(t *testing.T) {
		_, err := http.Get("http://localhost:8081")
		assert.NoError(t, err)
		assert.ErrorContains(t, assert.AnError, "connection refused")
	})

	t.Run("it should return unauthorized when API-key not presentGET", func(t *testing.T) {
		t.Fatal("todo")
	})
}
