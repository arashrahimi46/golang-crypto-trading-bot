package helpers

import (
	"github.com/arashrahimi46/golang-crypto-trading-bot/environment"
	"github.com/arashrahimi46/golang-crypto-trading-bot/exchanges"
	"github.com/shopspring/decimal"
)

//InitExchange initialize a new ExchangeWrapper binded to the specified exchange provided.
func InitExchange(exchangeConfig environment.ExchangeConfig, simulatedMode bool, fakeBalances map[string]decimal.Decimal, depositAddresses map[string]string) exchanges.ExchangeWrapper {
	if depositAddresses == nil && !simulatedMode {
		return nil
	}

	var exch exchanges.ExchangeWrapper
	switch exchangeConfig.ExchangeName {
	case "bittrexV2":
		exch = exchanges.NewBittrexV2Wrapper(exchangeConfig.PublicKey, exchangeConfig.SecretKey, depositAddresses)
	case "poloniex":
		exch = exchanges.NewPoloniexWrapper(exchangeConfig.PublicKey, exchangeConfig.SecretKey, depositAddresses)
	case "binance":
		exch = exchanges.NewBinanceWrapper(exchangeConfig.PublicKey, exchangeConfig.SecretKey, depositAddresses)
	case "bitfinex":
		exch = exchanges.NewBitfinexWrapper(exchangeConfig.PublicKey, exchangeConfig.SecretKey, depositAddresses)
	case "hitbtc":
		exch = exchanges.NewHitBtcV2Wrapper(exchangeConfig.PublicKey, exchangeConfig.SecretKey, depositAddresses)
	case "kucoin":
		exch = exchanges.NewKucoinWrapper(exchangeConfig.PublicKey, exchangeConfig.SecretKey, depositAddresses)
	default:
		return nil
	}

	if simulatedMode {
		if fakeBalances == nil {
			return nil
		}
		exch = exchanges.NewExchangeWrapperSimulator(exch, fakeBalances)
	}

	return exch
}
