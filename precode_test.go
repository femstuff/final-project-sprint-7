package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCorrectRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecoder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecoder, req)

	assert.Equal(t, 200, responseRecoder.Code)
	assert.NotEmpty(t, responseRecoder.Body)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)

	expected := `count missing`
	assert.Equal(t, responseRecorder.Body.String(), expected)

	countStr := req.URL.Query().Get("count")
	count, err := strconv.Atoi(countStr)
	require.Error(t, err)
	assert.Greater(t, totalCount, count)
	assert.NotEqual(t, 200, responseRecorder.Code)
}

func TestMainHandlerWhenWrongCityValue(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, 400, responseRecorder.Code)
}
