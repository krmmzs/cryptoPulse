package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	cryptopulse "github.com/krmmzs/cryptoPulse"
)

func main() {
	// Custom usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "cryptoPulse - Fetch cryptocurrency prices and fiat exchange rates\n\n")
		fmt.Fprintf(os.Stderr, "OPTIONS:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nEXAMPLES:\n")
		fmt.Fprintf(os.Stderr, "  %s                        # fetch both crypto and fiat data (default: BTCUSDT)\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -crypto                # fetch cryptocurrency prices only\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -fiat                  # fetch fiat exchange rates only\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -pair ETHUSDT          # fetch ETHUSDT price from Binance\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -crypto -pair ADAUSDT  # fetch only ADAUSDT price\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -help                  # show this help message\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "ENVIRONMENT:\n")
		fmt.Fprintf(os.Stderr, "  EXCHANGERATE_API_KEY    Required for fiat currency exchange rates\n\n")
	}

	// Define command-line flags
	cryptoFlag := flag.Bool("crypto", false, "fetch cryptocurrency prices only")
	fiatFlag := flag.Bool("fiat", false, "fetch fiat currency exchange rates only")
	pairFlag := flag.String("pair", "BTCUSDT", "trading pair to fetch (format depends on exchange, e.g., BTCUSDT for Binance, BTC-USDT for OKX)")
	flag.Parse()

	// Clients and Transports are safe for concurrent use by multiple goroutines and for efficiency should only be created once and re-used.
	// So create Client in main to concurrent
	httpClient := &http.Client{Timeout: 10 * time.Second}

	// Determine what to fetch based on flags
	// fetchCrypto: true or all of flag is false(default) to print
	fetchCrypto := *cryptoFlag || (!*cryptoFlag && !*fiatFlag) // default behavior includes crypto
	fetchFiat := *fiatFlag || (!*cryptoFlag && !*fiatFlag)     // default behavior includes fiat

	if fetchCrypto {
		// 创建Binance提供商实例
		binanceProvider := cryptopulse.NewBinanceProvider()
		
		// 验证交易对格式
		if !binanceProvider.ValidateSymbol(*pairFlag) {
			log.Printf("Warning: %s may not be a valid symbol for %s", *pairFlag, binanceProvider.GetName())
			log.Printf("Supported symbols include: %v", binanceProvider.GetSupportedSymbols()[:5]) // 显示前5个作为示例
		}
		
		// 使用接口方法获取价格
		cryptoPriceInfo, err := binanceProvider.FetchCryptoPrice(httpClient, *pairFlag, nil)
		if err != nil {
			log.Fatalf("Error fetching %s price from %s: %v", *pairFlag, binanceProvider.GetName(), err)
		}
		fmt.Printf("Fetched %s price from %s: %+v\n", *pairFlag, binanceProvider.GetName(), *cryptoPriceInfo)
	}

	if fetchFiat {
		// 创建汇率API提供商实例
		fiatProvider := cryptopulse.NewExchangeRateAPIProvider()
		
		// 使用接口方法获取汇率
		usdCnyRateInfo, err := fiatProvider.FetchFiatRate(httpClient, "USD", "CNY")
		if err != nil {
			log.Fatalf("Error fetching USD/CNY rate from %s: %v", fiatProvider.GetName(), err)
		}
		fmt.Printf("Fetched USD/CNY rate from %s: %+v\n", fiatProvider.GetName(), *usdCnyRateInfo)
	}

	// 示例：如果要从其他交易所获取价格，需要创建相应的Provider
	// okxProvider := cryptopulse.NewOKXProvider() // 需要先实现OKXProvider
	// btcPriceFromOKX, err := okxProvider.FetchCryptoPrice(httpClient, "BTC-USDT", nil)
}
