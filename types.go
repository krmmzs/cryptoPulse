package cryptopulse

import "time"

// CryptoExchangeConfig 定义了加密货币交易所的配置
type CryptoExchangeConfig struct {
	// BaseURL 是API的基础URL，例如 "https://api.binance.com"
	BaseURL string
	// URLPath 是API的路径模板，例如 "/api/v3/ticker/price"
	URLPath string
	// Source 是数据来源的标识，例如 "binance"
	Source string
	// QueryParam 是交易对参数的名称，例如 "symbol"(Binance), "instId"(OKX), "product_id"(Coinbase)
	QueryParam string
}

// CryptoPrice represents price information for binance.
type CryptoPrice struct {
	Symbol    string    `json:"symbol"`
	Price     string    `json:"price"`      // Holding strings to ensure precision
	Source    string    `json:"source"`     // Data sources, e.g. "binance"
	FetchedAt time.Time `json:"fetched_at"` // Time to get data
}

// FiatRate stands for Fiat Currency Exchange Rate Information
type exchangeRateAPIResponse struct {
	Result          string             `json:"result"`
	Documentation   string             `json:"documentation"`
	TermsOfUse      string             `json:"terms_of_use"`
	TimeLastUpdate  int64              `json:"time_last_update_unix"`
	BaseCode        string             `json:"base_code"`        // 注意：base_code
	ConversionRates map[string]float64 `json:"conversion_rates"` // 注意：conversion_rates
}

type FiatRate struct {
	BaseCurrency  string
	QuoteCurrency string
	Rate          float64
	Source        string
	FetchedAt     time.Time
}
