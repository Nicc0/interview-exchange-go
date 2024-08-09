package main

import (
	"exchang-go/pkg/crypto"
	"exchang-go/route/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setUpExchangeRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/exchange", v1.DoExchange)

	crypto.InitCryptoCurrencies()

	return router
}

func TestExchangeWBTCHandler(t *testing.T) {
	mockResponse := "{\"From\":\"WBTC\",\"To\":\"USDT\",\"Amount\":\"57613.367072\"}"

	r := setUpExchangeRouter()
	req, _ := http.NewRequest("GET", "/exchange?from=WBTC&to=USDT&amount=1.0", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestExchangeUSDTHandler(t *testing.T) {
	mockResponse := "{\"From\":\"USDT\",\"To\":\"BEER\",\"Amount\":\"40227.540476154573809708\"}"

	r := setUpExchangeRouter()
	req, _ := http.NewRequest("GET", "/exchange?from=USDT&to=BEER&amount=1.0", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestExchangeWithInvalidCurrencyHandler(t *testing.T) {
	mockResponse := "{}"

	r := setUpExchangeRouter()
	req, _ := http.NewRequest("GET", "/exchange?from=MATIC&to=GATE&amount=0.999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestExchangeWithoutAmountHandler(t *testing.T) {
	mockResponse := "{}"

	r := setUpExchangeRouter()
	req, _ := http.NewRequest("GET", "/exchange?from=USDT&to=GATE", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
