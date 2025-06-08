package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	cryptopulse "github.com/krmmzs/cryptoPulse"
)

func main() {
	// Clients and Transports are safe for concurrent use by multiple goroutines and for efficiency should only be created once and re-used.
	// So create Client in main to concurrent
	httpClient := &http.Client{Timeout: 10 * time.Second}

	// Get BTC/USDT prices from Binance (using default config)
	btcPriceInfo, err := cryptopulse.FetchCryptoPrice(httpClient, "BTCUSDT", nil)
	if err != nil {
		log.Fatalf("Error fetching BTC price: %v", err)
	}
	fmt.Printf("Fetched price: %+v\n", *btcPriceInfo)

	// 获取 USD/CNY 汇率
	usdCnyRateInfo, err := cryptopulse.FetchFiatRate(httpClient, "USD", "CNY")
	if err != nil {
		log.Fatalf("Error fetching USD/CNY rate: %v", err)
	}
	fmt.Printf("Fetched rate: %+v\n", *usdCnyRateInfo)

	// 计算 BTC 的 CNY 价格
	btcPriceInCny, err := cryptopulse.ConvertCryptoToFiat(btcPriceInfo, usdCnyRateInfo)
	if err != nil {
		log.Fatalf("Error converting BTC to CNY: %v", err)
	}
	fmt.Printf("BTC Price in CNY (from %s and %s): %.2f %s\n",
		btcPriceInfo.Source,
		usdCnyRateInfo.Source,
		btcPriceInCny,
		usdCnyRateInfo.QuoteCurrency,
	)

	// 示例：如果要从其他交易所获取价格
	// okxConfig := &cryptopulse.CryptoExchangeConfig{
	//     BaseURL:     "https://www.okx.com",
	//     URLPath:     "/api/v5/market/ticker",
	//     Source:      "okx",
	//     SymbolParam: "instId",
	// }
	// btcPriceFromOKX, err := cryptopulse.FetchCryptoPrice(httpClient, "BTC-USDT", okxConfig)
}
