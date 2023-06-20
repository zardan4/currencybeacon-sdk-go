package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/zardan4/currencybeacon-sdk-go/client/responses"
)

type Client struct {
	apiKey string
	client *http.Client
}

func NewClient(timeout time.Duration, apiKey string) (*Client, error) {
	if timeout <= 0 {
		return nil, errors.New("timeout must be positive")
	}

	return &Client{
		apiKey: apiKey,
		client: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

func (c *Client) StringifyEndpoint(endpoint string, params ...string) string {
	var paramsString string
	for _, p := range params {
		if p != "" {
			paramsString += "&" + p
		}
	}

	return fmt.Sprintf("https://api.currencybeacon.com/v1%s?api_key=%s%s", endpoint, c.apiKey, paramsString)
}

// check for valid api_key
func (c *Client) CheckValidAPIKey() bool {
	client := c.client

	res, _ := client.Get(c.StringifyEndpoint("/currencies"))
	return res.StatusCode != 401
}

// get all currencies
func (c *Client) GetAllCurrencies() (map[string]responses.Currency, error) {
	client := c.client

	res, err := client.Get(c.StringifyEndpoint("/currencies"))
	if err != nil {
		return nil, errors.New("error while fetching")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("error while reading response body")
	}

	var result responses.CurrenciesResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.New("error while transforming response from json")
	}

	return result.Response.Fiats, nil
}

// get one currency
func (c *Client) GetCurrency(currencyCode string) (responses.Currency, error) {
	allCurrencies, err := c.GetAllCurrencies()
	if err != nil {
		return responses.Currency{}, err
	}
	result, ok := allCurrencies[currencyCode]
	if !ok {
		return responses.Currency{}, errors.New("not a valid currency code")
	}

	return result, nil
}

// get exchange rate by date
// base - currency code that will be used as base converting other currencies
// date - date at which exchange rate is searching for YYYY-MM-DD
// currencies - list of currencies that you want to get exchange value. if want to get all of them just pass empty array
func (c *Client) GetExchangeRateByDate(base string, currencies []string, date string) (responses.ExchangeRateResponse, error) {
	client := c.client

	// params for request
	baseParam := fmt.Sprintf("base=%s", base) // base param
	// currencies param
	currenciesParam := "symbols="
	if len(currencies) == 0 {
		currenciesParam = ""
	}
	for i, c := range currencies {
		if i != len(currencies)-1 { // of it is not the last currency
			c += ","
		}
		currenciesParam += c
	}
	// date param
	dateParam := "date=" + date

	res, err := client.Get(c.StringifyEndpoint("/historical", baseParam, currenciesParam, dateParam))
	if err != nil {
		return responses.ExchangeRateResponse{}, errors.New("error while fetching")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return responses.ExchangeRateResponse{}, errors.New("error while reading response body")
	}

	var result responses.ExchangeRateResponseWrapper
	err = json.Unmarshal(body, &result)
	if err != nil {
		return responses.ExchangeRateResponse{}, err
	}

	return result.Response, nil
}

// get current exchange rate
// base - currency code that will be used as base converting other currencies
// currencies - list of currencies that you want to get exchange value. if want to get all of them just pass empty array
func (c *Client) GetExchangeRate(base string, currencies []string) (responses.ExchangeRateResponse, error) {
	currDate := time.Now().Format("2006-01-02")

	// just searching by current date
	return c.GetExchangeRateByDate(base, currencies, currDate)
}

// convert currencies
func (c *Client) Convert(from string, to string, amount float64) (responses.ConvertResponse, error) {
	client := c.client

	// params
	fromParam := "from=" + from
	toParam := "to=" + to
	amountParam := fmt.Sprintf("amount=%f", amount)

	res, err := client.Get(c.StringifyEndpoint("/convert", fromParam, toParam, amountParam))
	if err != nil {
		return responses.ConvertResponse{}, errors.New("error while fetching")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return responses.ConvertResponse{}, errors.New("error while reading response body")
	}

	var result responses.ConvertResponseWrapper
	if err = json.Unmarshal(body, &result); err != nil {
		return responses.ConvertResponse{}, err
	}

	return result.Response, nil
}
