package main

import (
	"exchang-go/pkg/clients"
	"exchang-go/pkg/setting"
	"exchang-go/route/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func setUpRatesRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/rates", v1.GetRates)

	setting.Setup()
	clients.InitClients()

	return router
}

func TestRatesWithCurrenciesHandler(t *testing.T) {
	patternResponse := `\[{"From":"USD","To":"GBP","Rate":(\d+)\.(\d+)},{"From":"GBP","To":"USD","Rate":(\d+)\.(\d+)},{"From":"USD","To":"EUR","Rate":(\d+)\.(\d+)},{"From":"EUR","To":"USD","Rate":(\d+)\.(\d+)},{"From":"GBP","To":"EUR","Rate":(\d+)\.(\d+)},{"From":"EUR","To":"GBP","Rate":(\d+)\.(\d+)}\]`
	regex := regexp.MustCompile(patternResponse)

	r := setUpRatesRouter()
	req, _ := http.NewRequest("GET", "/rates?currencies=USD,GBP,EUR", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.MatchRegex(t, string(responseData), regex)
}

func TestRatesWithTwoCurrenciesHandler(t *testing.T) {
	patternResponse := `\[{"From":"GBP","To":"EUR","Rate":(\d+).(\d+)},{"From":"EUR","To":"GBP","Rate":(\d+).(\d+)}\]`
	regex := regexp.MustCompile(patternResponse)

	r := setUpRatesRouter()
	req, _ := http.NewRequest("GET", "/rates?currencies=GBP,EUR", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.MatchRegex(t, string(responseData), regex)
}

func TestRatesWithoutCurrencyHandler(t *testing.T) {
	mockResponse := "{}"

	r := setUpRatesRouter()
	req, _ := http.NewRequest("GET", "/rates", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, mockResponse, string(responseData))
}
