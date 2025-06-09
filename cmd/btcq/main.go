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
		fmt.Fprintf(os.Stderr, "  %s              # fetch both crypto and fiat data (default)\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -crypto      # fetch cryptocurrency prices only\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -fiat        # fetch fiat exchange rates only\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -help        # show this help message\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "ENVIRONMENT:\n")
		fmt.Fprintf(os.Stderr, "  EXCHANGERATE_API_KEY    Required for fiat currency exchange rates\n\n")
	}

	// Define command-line flags
	cryptoFlag := flag.Bool("crypto", false, "fetch cryptocurrency prices only")
	fiatFlag := flag.Bool("fiat", false, "fetch fiat currency exchange rates only")
	flag.Parse()

	// Clients and Transports are safe for concurrent use by multiple goroutines and for efficiency should only be created once and re-used.
	// So create Client in main to concurrent
	httpClient := &http.Client{Timeout: 10 * time.Second}

	// Determine what to fetch based on flags
	fetchCrypto := *cryptoFlag || (!*cryptoFlag && !*fiatFlag) // default behavior includes crypto
	fetchFiat := *fiatFlag || (!*cryptoFlag && !*fiatFlag)     // default behavior includes fiat

	if fetchCrypto {
		// Get BTC/USDT prices from Binance (using default config)
		btcPriceInfo, err := cryptopulse.FetchCryptoPrice(httpClient, "BTCUSDT", nil)
		if err != nil {
			log.Fatalf("Error fetching BTC price: %v", err)
		}
		fmt.Printf("Fetched price: %+v\n", *btcPriceInfo)
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
