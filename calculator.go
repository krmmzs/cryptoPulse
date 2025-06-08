package cryptopulse

import (
	"fmt"
	"strconv"
)

// ConvertCryptoToFiat 将加密货币价格转换为指定法币价格
// 假设 cryptoPrice 的价格是相对于 fiatRate.BaseCurrency 的 (例如 BTC/USD, USD/CNY)
// 或者 cryptoPrice 的 Symbol (如 BTCUSDT) 的 quote 部分与 fiatRate.BaseCurrency 匹配
func ConvertCryptoToFiat(cryptoPrice *CryptoPrice, fiatRate *FiatRate) (float64, error) {
	if cryptoPrice == nil {
		return 0, fmt.Errorf("crypto price data cannot be nil")
	}
	if fiatRate == nil {
		return 0, fmt.Errorf("fiat rate data cannot be nil")
	}

	// 简单校验：确保加密价格的标价货币与汇率的基础货币一致
	// 例如, cryptoPrice.Symbol = "BTCUSDT", fiatRate.BaseCurrency = "USD"
	// 这是一个简化的假设，实际场景中你可能需要更复杂的逻辑来匹配交易对的报价货币
	// 例如，如果 cryptoPrice.Symbol 是 "BTCUSDT"，我们期望 fiatRate.BaseCurrency 是 "USD"
	// 这里我们假设 symbol 的后三位是报价货币
	var cryptoQuoteCurrency string
	if len(cryptoPrice.Symbol) > 3 {
		cryptoQuoteCurrency = cryptoPrice.Symbol[len(cryptoPrice.Symbol)-3:] // TODO: 这很粗糙，需要更健壮的解析
		// 实际项目中，CryptoPrice 结构体可能直接包含 BaseAsset 和 QuoteAsset 字段
	} else {
		return 0, fmt.Errorf("invalid crypto symbol format: %s", cryptoPrice.Symbol)
	}

	if cryptoQuoteCurrency != fiatRate.BaseCurrency {
		return 0, fmt.Errorf(
			"mismatch between crypto quote currency (%s) and fiat rate base currency (%s)",
			cryptoQuoteCurrency,
			fiatRate.BaseCurrency,
		)
	}

	priceFloat, err := strconv.ParseFloat(cryptoPrice.Price, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse crypto price '%s': %w", cryptoPrice.Price, err)
	}

	convertedPrice := priceFloat * fiatRate.Rate
	return convertedPrice, nil
}
