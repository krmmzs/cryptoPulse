// This file contains concrete implementations of the interfaces defined in interfaces.go
//
// 🏪 PROVIDER IMPLEMENTATIONS:
// Each provider implements one or more interfaces, providing actual data fetching logic
// for specific exchanges or APIs. See ARCHITECTURE.md for visual explanation.
//
// 🔌 HOW TO ADD NEW PROVIDERS:
// 1. Create a struct (e.g., OKXProvider)
// 2. Implement all required interface methods
// 3. See ARCHITECTURE.md "How to Add New Exchange" section for step-by-step guide
package cryptopulse

import (
	"net/http"
	"regexp"
	"strings"
)

// BinanceProvider 实现 CryptoPriceProvider 接口的Binance交易所提供商
type BinanceProvider struct {
	defaultConfig *CryptoExchangeConfig
}

// NewBinanceProvider 创建一个新的Binance提供商实例
func NewBinanceProvider() *BinanceProvider {
	return &BinanceProvider{
		defaultConfig: defaultBinanceConfig(),
	}
}

// FetchCryptoPrice 实现 CryptoPriceProvider 接口 - 获取加密货币价格
func (b *BinanceProvider) FetchCryptoPrice(client *http.Client, symbol string, config *CryptoExchangeConfig) (*CryptoPrice, error) {
	// 如果没有提供config，使用默认的Binance配置
	if config == nil {
		config = b.defaultConfig
	}

	// 复用现有的fetchCryptoPrice函数
	return fetchCryptoPrice(client, symbol, config)
}

// GetName 返回提供商名称
func (b *BinanceProvider) GetName() string {
	return "binance"
}

// ValidateSymbol 验证交易对格式是否符合Binance要求
// Binance格式：BTCUSDT, ETHUSDT (大写字母，无分隔符)
func (b *BinanceProvider) ValidateSymbol(symbol string) bool {
	if symbol == "" {
		return false
	}

	// Binance交易对格式：大写字母，长度通常6-12位，无分隔符
	// 例如：BTCUSDT, ETHUSDT, ADAUSDT
	// ^ - 字符串开始
	// [A-Z] - 只允许大写英文字母
	// {6,12} - 长度限制在6到12个字符之间
	// $ - 字符串结束
	matched, _ := regexp.MatchString(`^[A-Z]{6,12}$`, symbol)
	return matched && (strings.HasSuffix(symbol, "USDT") ||
		strings.HasSuffix(symbol, "BTC") ||
		strings.HasSuffix(symbol, "ETH") ||
		strings.HasSuffix(symbol, "BNB"))
}

// GetSupportedSymbols 返回支持的主要交易对列表
func (b *BinanceProvider) GetSupportedSymbols() []string {
	return []string{
		"BTCUSDT", "ETHUSDT", "BNBUSDT", "ADAUSDT", "DOTUSDT",
		"LINKUSDT", "LTCUSDT", "BCHUSDT", "XLMUSDT", "EOSUSDT",
		"TRXUSDT", "XRPUSDT", "ATOMUSDT", "VETUSDT", "NEOUSDT",
		// 对BTC的交易对
		"ETHBTC", "BNBBTC", "ADABTC", "DOTBTC", "LINKBTC",
	}
}

// ExchangeRateAPIProvider 实现 FiatRateProvider 接口的汇率API提供商
type ExchangeRateAPIProvider struct{}

// NewExchangeRateAPIProvider 创建一个新的汇率API提供商实例
func NewExchangeRateAPIProvider() *ExchangeRateAPIProvider {
	return &ExchangeRateAPIProvider{}
}

// FetchFiatRate 实现 FiatRateProvider 接口 - 获取法币汇率
func (e *ExchangeRateAPIProvider) FetchFiatRate(client *http.Client, baseCurrency, quoteCurrency string) (*FiatRate, error) {
	// 复用现有的fetchFiatRate函数
	return fetchFiatRate(client, baseCurrency, quoteCurrency)
}

// GetName 返回提供商名称
func (e *ExchangeRateAPIProvider) GetName() string {
	return "exchangerate-api"
}

// GetSupportedCurrencies 返回支持的主要货币列表
func (e *ExchangeRateAPIProvider) GetSupportedCurrencies() []string {
	return []string{
		"USD", "EUR", "GBP", "JPY", "AUD", "CAD", "CHF", "CNY",
		"SEK", "NZD", "MXN", "SGD", "HKD", "NOK", "TRY", "RUB",
		"INR", "BRL", "ZAR", "KRW", "THB", "PLN", "CZK", "HUF",
	}
}
