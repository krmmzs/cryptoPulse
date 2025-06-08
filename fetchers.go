package cryptopulse

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// FetchCryptoPrice fetches the latest price for a given cryptocurrency symbol from the Binance API.
//   - client: The *http.Client to use for the request.
//   - symbol: The crypto symbol, e.g., "BTCUSDT".
//
// It returns a *CryptoPrice with the price information on success, or an error on failure.
func FetchCryptoPrice(client *http.Client, symbol string) (*CryptoPrice, error) {
	if client == nil {
		return nil, fmt.Errorf("http client cannot be nil")
	}

	url := fmt.Sprintf(binanceAPIURL, symbol)
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch price for %s from Binance: %w", symbol, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Binance API returned non-200 status for %s: %d", symbol, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body) // 或 io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body for %s: %w", symbol, err)
	}

	var binanceResp struct { // 临时的内部结构体
		Symbol string `json:"symbol"`
		Price  string `json:"price"`
	}
	if err := json.Unmarshal(body, &binanceResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Binance response for %s: %w", symbol, err)
	}

	return &CryptoPrice{
		Symbol:    binanceResp.Symbol,
		Price:     binanceResp.Price,
		Source:    "binance",
		FetchedAt: time.Now().UTC(),
	}, nil
}

// FetchFiatRate fetches the latest exchange rate between two fiat currencies from the exchangerate.host API.
//   - client: The *http.Client to use for the request.
//   - baseCurrency: The base currency code, e.g., "USD".
//   - quoteCurrency: The quote currency code, e.g., "EUR".
//
// It returns a *FiatRate with the rate information on success, or an error on failure.
func FetchFiatRate(client *http.Client, baseCurrency, quoteCurrency string) (*FiatRate, error) {
	if client == nil {
		return nil, fmt.Errorf("http client cannot be nil")
	}

	// Read API key from an environment variable
	apiKey := os.Getenv(exchangeRateAPIKeyEnvVar)
	if apiKey == "" {
		return nil, fmt.Errorf("API key not found: environment variable %s is not set", exchangeRateAPIKeyEnvVar)
	}

	url := fmt.Sprintf(exchangeRateAPIURL, apiKey, baseCurrency)
	log.Printf("DEBUG: 正在从URL %s 获取法定货币汇率", url)
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("无法获取 %s/%s 的汇率: %w", baseCurrency, quoteCurrency, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("无法读取 %s/%s 的响应体: %w", baseCurrency, quoteCurrency, err)
	}
	log.Printf("DEBUG: %s/%s 的原始API响应: %s", baseCurrency, quoteCurrency, string(body))
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ExchangeRate API (v6) 返回非200状态码给 %s/%s: %d. 响应体: %s", baseCurrency, quoteCurrency, resp.StatusCode, string(body))
	}
	var apiResp exchangeRateAPIResponse // 使用新结构
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("无法将ExchangeRate API (v6) 响应的 %s/%s 解码: %w. 响应体是: %s", baseCurrency, quoteCurrency, err, string(body))
	}
	if apiResp.Result != "success" {
		// API可能会返回200 OK，但在JSON体中表示错误
		return nil, fmt.Errorf("ExchangeRate API (v6) 对 %s/%s 报告了错误: %s. 完整响应: %s", baseCurrency, quoteCurrency, apiResp.Result, string(body))
	}
	// 从ConversionRates映射中提取汇率
	rate, ok := apiResp.ConversionRates[quoteCurrency]
	if !ok {
		// 记录可用的汇率，以便调试。
		availableCurrencies := make([]string, 0, len(apiResp.ConversionRates))
		for k := range apiResp.ConversionRates {
			availableCurrencies = append(availableCurrencies, k)
		}
		log.Printf("DEBUG: 基础为 %s 的可转换汇率: %v", apiResp.BaseCode, availableCurrencies)
		return nil, fmt.Errorf("目标货币 %s 未在ExchangeRate API (v6) 响应中找到基础货币 %s. 在调试日志中检查可用货币。", quoteCurrency, apiResp.BaseCode)
	}
	return &FiatRate{
		BaseCurrency:  apiResp.BaseCode, // 使用API响应中的基础代码
		QuoteCurrency: quoteCurrency,
		Rate:          rate,
		Source:        "v6.exchangerate-api.com", // 更新源
		FetchedAt:     time.Now().UTC(),
	}, nil
}
