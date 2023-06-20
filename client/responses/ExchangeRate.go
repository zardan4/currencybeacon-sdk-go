package responses

import "fmt"

type ExchangeRateResponse struct {
	Date  string             `json:"date"`
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

type ExchangeRateResponseWrapper struct {
	Response ExchangeRateResponse `json:"response"`
}

func (e ExchangeRateResponse) Info() string {
	var res string

	for k, v := range e.Rates {
		res += fmt.Sprintf("%s costs %.2f %s\n", e.Base, v, k)
	}

	return res
}
