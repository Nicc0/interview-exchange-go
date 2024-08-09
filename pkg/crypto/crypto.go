package crypto

import (
	"errors"
	"golang.org/x/exp/maps"
	"math/big"
	"slices"
	"strings"
)

type CryptoCurrencyRate struct {
	Currency      string
	DecimalPlaces uint8
	RateText      string
	Rate          *big.Float
}

var (
	cryptoCurrencies = map[string]CryptoCurrencyRate{
		"BEER": {
			Currency:      "BEER",
			DecimalPlaces: 18,
			RateText:      "0.00002461",
			Rate:          new(big.Float).SetPrec(60).SetMode(big.ToNearestEven),
		},
		"FLOKI": {
			Currency:      "FLOKI",
			DecimalPlaces: 18,
			RateText:      "0.0001428",
			Rate:          new(big.Float).SetPrec(60).SetMode(big.ToNearestEven),
		},
		"GATE": {
			Currency:      "GATE",
			DecimalPlaces: 18,
			RateText:      "6.87",
			Rate:          new(big.Float).SetPrec(60).SetMode(big.ToNearestEven),
		},
		"USDT": {
			Currency:      "USDT",
			DecimalPlaces: 6,
			RateText:      "0.99",
			Rate:          new(big.Float).SetPrec(20).SetMode(big.ToNearestEven),
		},
		"WBTC": {
			Currency:      "WBTC",
			DecimalPlaces: 8,
			RateText:      "57037.22",
			Rate:          new(big.Float).SetPrec(27).SetMode(big.ToNearestEven),
		},
	}
)

func InitCryptoCurrencies() {
	for _, rate := range cryptoCurrencies {
		rate.Rate.Parse(rate.RateText, 10)
	}
}

func GetCryptoCurrencyRate(currency string) (*CryptoCurrencyRate, error) {
	val, ok := cryptoCurrencies[strings.ToUpper(currency)]

	if !ok {
		return nil, errors.New("currency not found")
	}

	return &val, nil
}

func IsAvailableCryptoCurrency(currency string) bool {
	return slices.Contains(GetAvailableCryptoCurrencies(), strings.ToUpper(currency))
}

func GetAvailableCryptoCurrencies() []string {
	return maps.Keys(cryptoCurrencies)
}
