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
		// Get cryptocurrency prices from Binance (using default config)
		cryptoPriceInfo, err := cryptopulse.FetchCryptoPrice(httpClient, *pairFlag, nil)
		if err != nil {
			log.Fatalf("Error fetching %s price: %v", *pairFlag, err)
		}
		fmt.Printf("Fetched price: %+v\n", *cryptoPriceInfo)
	}

	if fetchFiat {
		// 获取 USD/CNY 汇率
		usdCnyRateInfo, err := cryptopulse.FetchFiatRate(httpClient, "USD", "CNY")
		if err != nil {
			log.Fatalf("Error fetching USD/CNY rate: %v", err)
		}
		fmt.Printf("Fetched rate: %+v\n", *usdCnyRateInfo)
	}

	// 示例：如果要从其他交易所获取价格
	// okxConfig := &cryptopulse.CryptoExchangeConfig{
	//     BaseURL:     "https://www.okx.com",
	//     URLPath:     "/api/v5/market/ticker",
	//     Source:      "okx",
	//     QueryParam: "instId",
	// }
	// btcPriceFromOKX, err := cryptopulse.FetchCryptoPrice(httpClient, "BTC-USDT", okxConfig)
}
