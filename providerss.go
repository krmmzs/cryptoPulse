package cryptopulse

/**/
/* import "net/http" */
/**/
/* // BinanceProvider 实现 CryptoPriceProvider 接口的Binance交易所提供商 */
/* type BinanceProvider struct { */
/* 	defaultConfig *CryptoExchangeConfig */
/* } */
/**/
/* // NewBinanceProvider 创建一个新的Binance提供商实例 */
/* func NewBinanceProvider() *BinanceProvider { */
/* 	return &BinanceProvider{ */
/* 		defaultConfig: defaultBinanceConfig(), */
/* 	} */
/* } */
/**/
/* // FetchCryptoPrice 实现 CryptoPriceProvider 接口 - 获取加密货币价格 */
/* func (b *BinanceProvider) FetchCryptoPrice(client *http.Client, symbol string, config *CryptoExchangeConfig) (*CryptoPrice, error) { */
/* 	// 如果没有提供config，使用默认的Binance配置 */
/* 	if config == nil { */
/* 		config = b.defaultConfig */
/* 	} */
/**/
/* 	return FetchCryptoPrice(client, symbol, config) */
/* } */
