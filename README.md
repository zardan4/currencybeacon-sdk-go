# currencybeacon-sdk-go
sdk for fetching https://currencybeacon.com/

## How to install?
`go get github.com/zardan4/currencybeacon-sdk-go`

`import "github.com/zardan4/currencybeacon-sdk-go/client"`
#### USE!

## Methods
### Check valid api key
```
client.CheckValidAPIKey()
// returns bool
```
### Get all currencies 
```
client.GetAllCurrencies()
// returns map[string]responses.Currency
```
### Get one currency
```
client.GetCurrency("UAH")
// arg1: code of currency to return info about. more: https://currencybeacon.com/supported-currencies
// returns responses.Currency
```
### Get exchange rate by date
```
client.GetExchangeRateByDate("UAH", []string{"EUR", "USD"}, "2008-12-07")
// arg1: code of currency to be as convertation base
// arg2: slice of currencies' codes to convert in. if you want all just leave empty slice
// arg3: date in format YYYY-MM-DD
// returns responses.ExchangeRateResponse
```
### Get current exchange rate
```
client.GetExchangeRate("UAH", []string{"EUR", "USD"})
// arg1: code of currency to be as convertation base
// arg2: slice of currencies' codes to convert in. if you want all just leave empty slice
// returns responses.ExchangeRateResponse
```
### Convert currencies
```
client.Convert("UAH", "USD", 20)
// arg1: code of currency to convert from
// arg2: currency to convert in
// arg3: amount of {arg1} to convert in {arg2}
// returns responses.ConvertResponse
```
