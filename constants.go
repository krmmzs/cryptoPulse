package cryptopulse

const (
	// see https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information
	binanceAPIURL = "https://api.binance.com/api/v3/ticker/price?symbol=%s"
	// see https://www.exchangerate-api.com/docs/overview
	exchangeRateAPIURL = "https://v6.exchangerate-api.com/v6/%s/latest/%s"
)

const exchangeRateAPIKeyEnvVar = "EXCHANGERATE_API_KEY"
