package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	excepted := len(strings.Split(responseRecorder.Body.String(), ","))
	assert.Greater(t, totalCount, excepted)
	assert.NotEqual(t, 200, responseRecorder.Code)
}

func TestMainHandlerWhenWrongCityValue(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, 400, responseRecorder.Code)
}
