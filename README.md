# cryptoPulse

<div align="center">

[![Go Version](https://img.shields.io/badge/Go-1.24+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/krmmzs/cryptoPulse)](https://goreportcard.com/report/github.com/krmmzs/cryptoPulse)

A fast and lightweight Go library and CLI tool for fetching cryptocurrency prices and fiat currency exchange rates.

</div>

## ✨ Features

- 🚀 **Fast & Lightweight**: Efficient concurrent data fetching
- 💰 **Multi-Exchange Support**: Binance (with extensible provider pattern for other exchanges)
- 💱 **Fiat Currency Rates**: USD/CNY and other currency pairs via ExchangeRate-API
- 🔧 **Flexible Usage**: Use as CLI tool or Go library
- 🏗️ **Extensible Architecture**: Easy to add new exchanges and data sources
- ⚡ **Provider Pattern**: Clean interface-based design for maximum flexibility

## 🚀 Quick Start

### Installation

#### Install from source
```bash
# Install directly to your GOPATH/bin
go install github.com/krmmzs/cryptoPulse/cmd/btcq@latest

# Or clone and build locally
git clone https://github.com/krmmzs/cryptoPulse.git
cd cryptoPulse
make install
```

#### Build from source
```bash
git clone https://github.com/krmmzs/cryptoPulse.git
cd cryptoPulse
make build
./bin/btcq
```

### Basic Usage

```bash
# Fetch both crypto and fiat data (default: BTC/USDT + USD/CNY)
btcq

# Fetch only cryptocurrency prices
btcq -crypto

# Fetch only fiat exchange rates
btcq -fiat

# Fetch specific trading pair
btcq -pair ETHUSDT

# Fetch only specific crypto pair
btcq -crypto -pair ADAUSDT
```

## 📖 Detailed Usage

### CLI Commands

| Command | Description |
|---------|-------------|
| `btcq` | Fetch both crypto and fiat data (default: BTCUSDT + USD/CNY) |
| `btcq -crypto` | Fetch cryptocurrency prices only |
| `btcq -fiat` | Fetch fiat exchange rates only |
| `btcq -pair SYMBOL` | Specify trading pair (e.g., ETHUSDT, ADAUSDT) |
| `btcq -help` | Show help message |

### Environment Variables

| Variable | Required | Description |
|----------|----------|-------------|
| `EXCHANGERATE_API_KEY` | Yes (for fiat) | API key for ExchangeRate-API service |

### Example Output

```bash
$ btcq
Fetched BTCUSDT price from Binance: {Symbol:BTCUSDT Price:43250.50000000 Source:binance FetchedAt:2024-01-15 10:30:00}
Fetched USD/CNY rate from ExchangeRate-API: {BaseCurrency:USD QuoteCurrency:CNY Rate:7.2145 Source:exchangerate-api FetchedAt:2024-01-15 10:30:01}
```

## 🛠️ Development

### Project Structure

```
cryptoPulse/
├── cmd/btcq/           # CLI application
├── bin/                # Built binaries
├── interfaces.go       # Provider interfaces
├── providers.go        # Provider implementations
├── fetchers.go         # Core fetching logic
├── types.go           # Data structures
├── constants.go       # Configuration constants
├── Makefile          # Build automation
└── ARCHITECTURE.md   # Detailed architecture docs
```

### Building

```bash
# Build binary
make build

# Install to GOPATH/bin
make install

# Remove installed binary
make uninstall

# Clean build artifacts
make clean

# Show all available commands
make help
```

### Running Tests

```bash
go test ./...
```

### Using as a Library

```go
package main

import (
    "fmt"
    "net/http"
    "time"
    
    cryptopulse "github.com/krmmzs/cryptoPulse"
)

func main() {
    httpClient := &http.Client{Timeout: 10 * time.Second}
    
    // Create Binance provider
    binanceProvider := cryptopulse.NewBinanceProvider()
    
    // Fetch crypto price
    price, err := binanceProvider.FetchCryptoPrice(httpClient, "BTCUSDT", nil)
    if err != nil {
        panic(err)
    }
    fmt.Printf("BTC Price: %+v\n", *price)
    
    // Create fiat rate provider
    fiatProvider := cryptopulse.NewExchangeRateAPIProvider()
    
    // Fetch fiat rate
    rate, err := fiatProvider.FetchFiatRate(httpClient, "USD", "CNY")
    if err != nil {
        panic(err)
    }
    fmt.Printf("USD/CNY Rate: %+v\n", *rate)
}
```

## 🏗️ Architecture

cryptoPulse uses a clean provider pattern that makes it easy to add new exchanges and data sources:

- **CryptoPriceProvider**: Interface for cryptocurrency price providers
- **FiatRateProvider**: Interface for fiat currency rate providers  
- **Extensible Design**: Add new exchanges by implementing the provider interfaces

For detailed architecture documentation, see [ARCHITECTURE.md](./ARCHITECTURE.md).

### Adding New Exchanges

```go
// Example: Adding OKX support
type OKXProvider struct {
    config CryptoExchangeConfig
}

func (p *OKXProvider) FetchCryptoPrice(client *http.Client, symbol string, config *CryptoExchangeConfig) (*CryptoPrice, error) {
    // Implementation for OKX API
}

func (p *OKXProvider) GetName() string {
    return "OKX"
}

// Use the new provider
okxProvider := &OKXProvider{config: okxConfig}
price, err := okxProvider.FetchCryptoPrice(httpClient, "BTC-USDT", nil)
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go best practices and idioms
- Add tests for new functionality
- Update documentation as needed
- Use the existing provider pattern for new exchanges
- Maintain backward compatibility

## 📝 API Reference

### Core Types

```go
type CryptoPrice struct {
    Symbol    string    `json:"symbol"`
    Price     string    `json:"price"`
    Source    string    `json:"source"`
    FetchedAt time.Time `json:"fetched_at"`
}

type FiatRate struct {
    BaseCurrency  string
    QuoteCurrency string
    Rate          float64
    Source        string
    FetchedAt     time.Time
}
```

### Provider Interfaces

```go
type CryptoPriceProvider interface {
    FetchCryptoPrice(client *http.Client, symbol string, config *CryptoExchangeConfig) (*CryptoPrice, error)
    GetName() string
    ValidateSymbol(symbol string) bool
    GetSupportedSymbols() []string
}

type FiatRateProvider interface {
    FetchFiatRate(client *http.Client, baseCurrency, quoteCurrency string) (*FiatRate, error)
    GetName() string
    GetSupportedCurrencies() []string
}
```

## 📊 Supported Exchanges & Currencies

### Cryptocurrency Exchanges
- ✅ **Binance**: Full support with symbol validation
- 🔄 **OKX**: Framework ready, implementation pending
- 🔄 **Coinbase**: Framework ready, implementation pending

### Fiat Currency Sources
- ✅ **ExchangeRate-API**: 170+ currencies supported

## 🔗 Links

- 📋 [Architecture Documentation](./ARCHITECTURE.md)
- 🐛 [Issue Tracker](https://github.com/krmmzs/cryptoPulse/issues)
- 📖 [Go Documentation](https://pkg.go.dev/github.com/krmmzs/cryptoPulse)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Binance API](https://binance-docs.github.io/apidocs/) for cryptocurrency data
- [ExchangeRate-API](https://exchangerate-api.com/) for fiat currency rates
- Go community for excellent tooling and libraries

---

<div align="center">

**[⬆ Back to Top](#cryptopulse)**

Made with ❤️ by the cryptoPulse team

</div>