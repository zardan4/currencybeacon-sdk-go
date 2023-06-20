package responses

import "fmt"

// currency represents
type currenciesResponseRes struct {
	Fiats  map[string]Currency `json:"fiats"`
	Crypto map[string]Currency `json:"crypto"`
}
type CurrenciesResponse struct {
	Response currenciesResponseRes `json:"response"`
}

type Currency struct {
	Name      string   `json:"currency_name"`
	Code      string   `json:"currency_code"`
	Countries []string `json:"countries"`
}

func (c *Currency) Info() string {
	return fmt.Sprintf("[%s] %s is a currency of %d countries\n", c.Code, c.Name, len(c.Countries))
}
