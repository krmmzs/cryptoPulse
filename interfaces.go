// Package cryptopulse provides interfaces and implementations for fetching cryptocurrency
// prices and fiat currency exchange rates from various sources.
//
// 📋 ARCHITECTURE OVERVIEW:
// See ARCHITECTURE.md for visual diagrams and detailed explanations of how
// interfaces, providers, and data flow work together.
//
// 🧩 INTERFACE DESIGN:
// This file defines the contracts (interfaces) that all data providers must implement.
// Think of interfaces as "plug standards" - any implementation that fits the standard can be used.
package cryptopulse

import (
	"net/http"
)

// CryptoPriceProvider 定义了获取加密货币价格的行为
type CryptoPriceProvider interface {
	// FetchCryptoPrice 获取指定交易对的价格
	// symbol: 交易对符号（格式取决于具体交易所）
	// config: 可选的交易所配置，为nil时使用提供商默认配置
	FetchCryptoPrice(client *http.Client, symbol string, config *CryptoExchangeConfig) (*CryptoPrice, error)
	
	// GetName 返回提供商名称
	GetName() string
	
	// GetSupportedSymbols 返回支持的交易对列表（可选实现）
	GetSupportedSymbols() []string
	
	// ValidateSymbol 验证交易对格式是否符合该交易所要求
	ValidateSymbol(symbol string) bool
}

// FiatRateProvider 定义了获取法币汇率的行为
type FiatRateProvider interface {
	// FetchFiatRate 获取两种法币之间的汇率
	FetchFiatRate(client *http.Client, baseCurrency, quoteCurrency string) (*FiatRate, error)
	
	// GetName 返回提供商名称
	GetName() string
	
	// GetSupportedCurrencies 返回支持的货币列表
	GetSupportedCurrencies() []string
}

