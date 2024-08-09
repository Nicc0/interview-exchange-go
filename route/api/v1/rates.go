package v1

import (
	"errors"
	"exchang-go/pkg/clients"
	"exchang-go/pkg/setting"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/currency"
	"net/http"
	"strings"
)

const CurrenciesUrl = "latest.json"

type RatesPayload struct {
	Currencies string `form:"currencies" json:"currencies"`
}

type OEARatesResponse struct {
	Timestamp int64              `json:"timestamp"`
	Base      string             `json:"base"`
	Rates     map[string]float32 `json:"rates"`
}

type CurrencyRate struct {
	From string
	To   string
	Rate float32
}

func GetRates(c *gin.Context) {
	var err error
	var currencies []string
	var rates map[string]float32

	payload, err := validateRatesPayload(c)

	if err == nil {
		currencies = strings.Split(payload.Currencies, ",")

		rates, err = fetchExchangeRates(currencies)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})

		return
	}

	baseCurrency := setting.OpenExchangeApiSetting.BaseCurrency
	response := getCombinationRates(currencies)

	for i, pair := range response {
		rateKey := pair.From + pair.To

		if rate, ok := rates[rateKey]; ok {
			response[i].Rate = rate
		} else if rateFromBase, okFrom := rates[baseCurrency+pair.From]; okFrom {
			if rateToBase, okTo := rates[baseCurrency+pair.To]; okTo {
				response[i].Rate = rateToBase / rateFromBase
			}
		}
	}

	c.JSON(http.StatusOK, response)
}

func validateRatesPayload(c *gin.Context) (*RatesPayload, error) {
	var payload RatesPayload

	if c.ShouldBindQuery(&payload) != nil {
		return nil, errors.New("no payload given")
	}

	currencies := strings.Split(payload.Currencies, ",")

	for _, isoCurrency := range currencies {
		_, err := currency.ParseISO(isoCurrency)

		if err != nil {
			return nil, errors.New("invalid currency")
		}
	}

	if len(currencies) < 2 {
		return nil, errors.New("not enough currencies given")
	}

	return &payload, nil
}

func getCombinationRates(currencies []string) []CurrencyRate {
	var rates []CurrencyRate

	for i := 0; i < len(currencies); i++ {
		for j := i + 1; j < len(currencies); j++ {
			rates = append(rates, CurrencyRate{
				From: currencies[i],
				To:   currencies[j],
			})
			rates = append(rates, CurrencyRate{
				From: currencies[j],
				To:   currencies[i],
			})
		}
	}

	return rates
}

func fetchExchangeRates(currencies []string) (map[string]float32, error) {
	resp, err := clients.OpenExchangeApiClient.R().
		SetQueryParam("symbols", strings.Join(currencies, ",")).
		SetResult(&OEARatesResponse{}).
		Get(CurrenciesUrl)

	if err != nil {
		return nil, err
	}

	exchangeRates := make(map[string]float32)
	ratesResponse := resp.Result().(*OEARatesResponse)

	for to, rate := range ratesResponse.Rates {
		exchangeRates[ratesResponse.Base+to] = rate
		exchangeRates[to+ratesResponse.Base] = 1 / rate
	}

	return exchangeRates, nil
}
