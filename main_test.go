package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code) //проверяем статус код ответа, если не проходит останавливаем тест
	assert.NotEmpty(t, responseRecorder.Body)              //проверяем, что тело ответа не пустое
}

func TestMainHandlerWhenMissingCount(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expected := `count missing`

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code) //проверяем статус код ответа, если не проходит останавливаем тест
	assert.Equal(t, expected, responseRecorder.Body.String())      //проверяем тело ответа на соответствие ожиданиям
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	require.Equal(t, http.StatusOK, responseRecorder.Code) //проверяем статус код ответа, если не проходит останавливаем тест
	assert.Len(t, list, totalCount)                        //проверяем длину списка кафе из ответа с ожидаемой длинной
}
