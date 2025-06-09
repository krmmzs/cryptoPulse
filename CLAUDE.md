# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

cryptoPulse is a Go library and CLI tool for fetching cryptocurrency prices and fiat currency exchange rates. The main purpose is to quickly get BTC/USDT prices from crypto exchanges (primarily Binance) and USD/CNY exchange rates for display.

## Common Commands

### Build and Run
```bash
go build ./cmd/btcq
./btcq
```

### Run directly
```bash
go run ./cmd/btcq
```

### Module management
```bash
go mod tidy
go mod download
```

## Architecture

### Core Components

- **Fetchers** (`fetchers.go`): Main data fetching logic for crypto prices and fiat rates
- **Types** (`types.go`): Data structures for CryptoPrice, FiatRate, and exchange configurations
- **Calculator** (`calculator.go`): Price conversion logic between crypto and fiat currencies
- **Interfaces** (`interfaces.go`): Provider interfaces for extensibility

### Key Functions

- `FetchCryptoPrice()`: Fetches crypto prices from exchanges (configurable, defaults to Binance)
- `FetchFiatRate()`: Fetches fiat exchange rates from exchangerate-api.com (requires EXCHANGERATE_API_KEY env var)
- `ConvertCryptoToFiat()`: Converts crypto prices to target fiat currency (available in library but not used in main CLI)

### Configuration

The system uses a flexible configuration approach via `CryptoExchangeConfig` struct that allows fetching from different exchanges by specifying:
- BaseURL
- URLPath 
- Source identifier
- Symbol parameter name

Default configuration uses Binance API, but can be overridden for other exchanges like OKX.

### Environment Variables

- `EXCHANGERATE_API_KEY`: Required for fiat currency exchange rate fetching

### Error Handling

Functions return detailed error messages with context about which API failed and why, including HTTP status codes and response bodies for debugging.
