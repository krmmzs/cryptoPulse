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

// --- 具体实现 ---
// BinanceProvider 实现了 CryptoPriceProvider
type BinanceProvider struct{}

func (p *BinanceProvider) FetchCryptoPrice(client *http.Client, symbol string) (*CryptoPrice, error) {
	// 这里调用上面我们写的 FetchCryptoPrice，或者直接内联实现
	// 为了简单，我们假设 FetchCryptoPrice 是一个通用的实现，由 Provider 决定调用的具体 URL
	// 或者 FetchCryptoPrice 本身就可以作为 BinanceProvider 的一个方法。（取决于你如何组织）
	// 这里我们假设 FetchCryptoPrice 是一个独立的函数，可以被不同 provider 包装
	// ... (实际的 HTTP GET 和解析逻辑，类似之前的 FetchCryptoPrice)
	// 为简化，我们暂时还是调用之前独立的函数，只是演示接口的用法
	// 理想情况下，这个方法内部会直接实现针对币安的逻辑
	return FetchCryptoPrice(client, symbol) // 但注意，FetchCryptoPrice 内部已写死 Binance
}

// ExchangeRateHostProvider 实现了 FiatRateProvider
type ExchangeRateHostProvider struct{}

func (p *ExchangeRateHostProvider) FetchFiatRate(client *http.Client, base, quote string) (*FiatRate, error) {
	return FetchFiatRate(client, base, quote) // 同上，理想情况下逻辑在此方法内部
}
