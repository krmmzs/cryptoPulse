package cryptopulse

import (
	"net/http"
)

// CryptoPriceProvider 定义了获取加密货币价格的行为
type CryptoPriceProvider interface {
	FetchCryptoPrice(client *http.Client, symbol string) (*CryptoPrice, error)
}

// FiatRateProvider 定义了获取法币汇率的行为
type FiatRateProvider interface {
	FetchFiatRate(client *http.Client, baseCurrency, quoteCurrency string) (*FiatRate, error)
}
