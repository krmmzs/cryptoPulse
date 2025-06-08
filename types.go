package cryptopulse

import "time"

// CryptoPrice represents price information for cryptocurrencies.
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
