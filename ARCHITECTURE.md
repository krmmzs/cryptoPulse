# cryptoPulse Architecture Guide

> 📋 **Purpose**: Help developers (including future me) understand the codebase structure and extend functionality easily.

## 🏗️ Project Structure Overview

```
cryptoPulse/
├── 📋 interfaces.go     # Contract definitions (what functionality looks like)
├── 🏪 providers.go      # Concrete implementations (actual data fetchers)
├── 🔧 fetchers.go       # Core fetching logic (reusable functions)
├── 📊 types.go          # Data structures (what data looks like)
├── ⚙️  constants.go     # Configuration constants
└── 🚀 cmd/btcq/main.go  # CLI application entry point
```

## 🧩 Interface Architecture (interfaces.go)

```
                            cryptoPulse/interfaces.go
                                   📋 CONTRACT LAYER
    
    ┌─────────────────────────────────────────────────────────────────────────┐
    │                          🏪 PROVIDER INTERFACES                         │
    └─────────────────────────────────────────────────────────────────────────┘
    
    
         💰 CryptoPriceProvider                    💱 FiatRateProvider
    ┌─────────────────────────────┐         ┌─────────────────────────────┐
    │        📊 CRYPTO PRICES     │         │        💲 CURRENCY RATES    │
    │                             │         │                             │
    │ 🔧 Methods:                 │         │ 🔧 Methods:                 │
    │  ├─ FetchCryptoPrice()      │         │  ├─ FetchFiatRate()         │
    │  │   📈 Get trading pair     │         │  │   📊 Get USD→CNY rate     │
    │  │   Input: BTCUSDT         │         │  │   Input: USD, CNY         │
    │  │   Output: *CryptoPrice   │         │  │   Output: *FiatRate       │
    │  │                          │         │  │                          │
    │  ├─ GetName()               │         │  ├─ GetName()               │
    │  │   🏷️  "binance"           │         │  │   🏷️  "exchangerate-api"  │
    │  │                          │         │  │                          │
    │  ├─ GetSupportedSymbols()   │         │  ├─ GetSupportedCurrencies()│
    │  │   📝 ["BTCUSDT","ETHUSDT"]│         │  │   📝 ["USD","EUR","CNY"]   │
    │  │                          │         │  │                          │
    │  └─ ValidateSymbol()        │         │  └─ (no validation method)  │
    │      ✅ Check format valid   │         │                             │
    └─────────────────────────────┘         └─────────────────────────────┘
                    │                                         │
                    │                                         │
                    ▼                                         ▼
    
    ┌─────────────────────────────┐         ┌─────────────────────────────┐
    │     🏢 BinanceProvider       │         │  🏦 ExchangeRateAPIProvider │
    │                             │         │                             │
    │ 🎯 Implements:              │         │ 🎯 Implements:              │
    │ CryptoPriceProvider         │         │ FiatRateProvider            │
    │                             │         │                             │
    │ 🌐 Data Source:             │         │ 🌐 Data Source:             │
    │ api.binance.com             │         │ exchangerate-api.com        │
    │                             │         │                             │
    │ 💼 Handles:                 │         │ 💼 Handles:                 │
    │ BTCUSDT, ETHUSDT...         │         │ USD/CNY, EUR/USD...         │
    └─────────────────────────────┘         └─────────────────────────────┘
```

## 🔄 Data Flow Diagram

```
    User Command          Interface Contract       Implementation        External API
    ─────────────         ─────────────────       ─────────────         ──────────────
    
    ./btcq -pair ETHUSDT
         │
         ▼
    main.go creates
    BinanceProvider
         │
         ▼
    provider.FetchCryptoPrice()  ──────►  BinanceProvider  ──────►  api.binance.com
         │                               implements                       │
         │                               interface                        │
         ▼                               contract                         ▼
    Display result ◄─────────────────────────────────────────────  JSON Response
```

## 🧠 Mental Models & Key Concepts

### 🔌 Interface = Plug Standard
```
Interface defines:     │  Implementation provides:
"什么形状的插头"        │  "实际的设备"
                      │
- Method signatures   │  - Actual logic
- Input/Output types  │  - API calls
- Behavior contract   │  - Error handling
```

### 🏪 Provider Pattern
```
🏪 Any "store" that sells crypto prices must:
   ✅ Have a FetchCryptoPrice() method
   ✅ Have a GetName() method  
   ✅ Be able to ValidateSymbol()
   ✅ Know GetSupportedSymbols()

Currently implemented stores:
🏢 BinanceProvider    → Sells data from Binance
🏦 ExchangeRateAPI    → Sells fiat rates
```

## 🚀 How to Add New Exchange (e.g., OKX)

### Step 1: Create Provider
```go
// In providers.go
type OKXProvider struct {
    defaultConfig *CryptoExchangeConfig
}

func NewOKXProvider() *OKXProvider {
    return &OKXProvider{
        defaultConfig: &CryptoExchangeConfig{
            BaseURL:    "https://www.okx.com",
            URLPath:    "/api/v5/market/ticker",
            Source:     "okx",
            QueryParam: "instId",
        },
    }
}
```

### Step 2: Implement Interface Methods
```go
func (o *OKXProvider) FetchCryptoPrice(client *http.Client, symbol string, config *CryptoExchangeConfig) (*CryptoPrice, error) {
    // Use existing FetchCryptoPrice function with OKX config
}

func (o *OKXProvider) GetName() string { return "okx" }

func (o *OKXProvider) ValidateSymbol(symbol string) bool {
    // OKX uses "BTC-USDT" format (with dash)
    return strings.Contains(symbol, "-")
}

func (o *OKXProvider) GetSupportedSymbols() []string {
    return []string{"BTC-USDT", "ETH-USDT", "ADA-USDT"}
}
```

### Step 3: Use in main.go
```go
// Choose provider based on flag or config
var provider CryptoPriceProvider
if *exchangeFlag == "okx" {
    provider = NewOKXProvider()
} else {
    provider = NewBinanceProvider()  // default
}

price, err := provider.FetchCryptoPrice(httpClient, *pairFlag, nil)
```

## 📚 File Responsibilities

| File | Purpose | Key Concepts |
|------|---------|--------------|
| `interfaces.go` | 📋 **Contracts** | Defines what methods providers must have |
| `providers.go` | 🏪 **Implementations** | Actual code that fetches data |
| `fetchers.go` | 🔧 **Core Logic** | Reusable HTTP fetching functions |
| `types.go` | 📊 **Data Models** | Structures for CryptoPrice, FiatRate |
| `main.go` | 🚀 **Application** | CLI logic and user interaction |

## 🎯 Design Principles

### 1. **Separation of Concerns**
- Interfaces define behavior contracts
- Providers implement specific exchange logic  
- Fetchers handle HTTP communication
- Types define data structures

### 2. **Open/Closed Principle**
- ✅ Open for extension (add new exchanges)
- ✅ Closed for modification (don't change interfaces)

### 3. **Dependency Inversion**
- High-level code depends on interfaces
- Low-level code implements interfaces
- Easy to swap implementations

## 🔧 Extension Points

### Adding New Data Sources
1. **New Crypto Exchange**: Implement `CryptoPriceProvider`
2. **New Fiat Rate API**: Implement `FiatRateProvider`  
3. **New Data Type**: Add to `types.go` + create new interface

### Adding New Features
1. **Price Caching**: Add caching layer to providers
2. **Rate Limiting**: Add rate limiting to HTTP clients
3. **Historical Data**: Extend interfaces with historical methods
4. **WebSocket Support**: Add real-time data interfaces

## 💡 Memory Hooks

```
🧩 Interface = 拼图的形状规格
🏪 Provider = 能装进这个形状的具体商店  
🔌 Plug & Play = 换个Provider就像换插头一样简单
🎯 Contract = 说好的规则，大家都要遵守
🏗️ Architecture = 房子的蓝图，指导怎么盖房子
```

---

> 💡 **Quick Reference**: When you forget how this works, start here:
> 1. Look at `interfaces.go` - understand the contracts
> 2. Look at `providers.go` - see how contracts are implemented  
> 3. Look at `main.go` - see how everything connects
> 4. This document explains the "why" behind the "what"