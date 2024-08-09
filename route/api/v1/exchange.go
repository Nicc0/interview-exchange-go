package v1

import (
	"errors"
	"exchang-go/pkg/crypto"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
)

type ExchangePayload struct {
	From   string `form:"from" json:"from"`
	To     string `form:"to" json:"to"`
	Amount string `form:"amount" json:"amount"`
}

type CryptoCurrencyRate struct {
	From   string
	To     string
	Amount string
}

func DoExchange(c *gin.Context) {
	payload, err := validateExchangePayload(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	from, _ := crypto.GetCryptoCurrencyRate(payload.From)
	to, _ := crypto.GetCryptoCurrencyRate(payload.To)

	amount, _, err := big.ParseFloat(payload.Amount, 10, 60, big.ToNearestEven)
	amount.Mul(amount, from.Rate)
	amount.Quo(amount, to.Rate)

	response := CryptoCurrencyRate{
		From:   from.Currency,
		To:     to.Currency,
		Amount: amount.Text('f', int(to.DecimalPlaces)),
	}

	c.JSON(http.StatusOK, response)
}

func validateExchangePayload(c *gin.Context) (*ExchangePayload, error) {
	var payload ExchangePayload

	if c.ShouldBindQuery(&payload) != nil {
		return nil, errors.New("no payload given")
	}

	if !crypto.IsAvailableCryptoCurrency(payload.From) {
		return nil, errors.New("invalid field \"from\" for crypto currency")
	}

	if !crypto.IsAvailableCryptoCurrency(payload.To) {
		return nil, errors.New("invalid field \"to\" for crypto currency")
	}

	if _, _, err := big.ParseFloat(payload.Amount, 10, 60, big.ToNearestEven); err != nil {
		return nil, errors.New("invalid field \"amount\", cannot parse float")
	}

	return &payload, nil
}
