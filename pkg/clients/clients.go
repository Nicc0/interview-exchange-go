package clients

import (
	"exchang-go/pkg/setting"
	"github.com/go-resty/resty/v2"
)

var (
	OpenExchangeApiClient = *resty.New()
)

func InitClients() {
	OpenExchangeApiClient.
		SetBaseURL(setting.OpenExchangeApiSetting.Url).
		SetHeader("Accept", "application/json").
		SetQueryParam("app_id", setting.OpenExchangeApiSetting.Key).
		SetQueryParam("base", setting.OpenExchangeApiSetting.BaseCurrency)
}
