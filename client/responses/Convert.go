package responses

import "fmt"

type ConvertResponseWrapper struct {
	Response ConvertResponse `json:"response"`
}

type ConvertResponse struct {
	Time   int     `json:"timestamp"`
	Date   string  `json:"date"`
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
	Value  float64 `json:"value"`
}

func (c ConvertResponse) Info() string {
	return fmt.Sprintf("%.2f %s equals to %.2f %s at %s", c.Amount, c.From, c.Value, c.To, c.Date)
}
