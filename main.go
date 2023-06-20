package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/zardan4/currencybeacon-sdk-go/client"
)

func main() {
	// user api key input
	var usersApiKey string
	fmt.Println("Write your API key for currencybeacon:")
	fmt.Scanln(&usersApiKey)
	usersApiKey = strings.Trim(usersApiKey, " ")

	client, err := client.NewClient(time.Second*2, usersApiKey)
	if err != nil {
		log.Fatal(err)
	}

	// check for valid API key
	if !client.CheckValidAPIKey() {
		log.Fatal("invalid API key")
	}

	// get all currencies test
	fmt.Println("========================== Get all currencies test ==========================")
	res1, err := client.GetAllCurrencies()
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range res1 {
		fmt.Print(c.Info())
	}

	// get one currency test
	fmt.Println("========================== Get one currency test ==========================")
	res2, err := client.GetCurrency("UAH")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(res2.Info())

	// get exchange value test. отримуємо на даний момент
	fmt.Println("========================== Get exchange value test test ==========================")
	res3, err := client.GetExchangeRate("UAH", []string{"EUR", "USD"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(res3.Info())

	// get exchange value by date test. отримуємо по конкретній даті
	fmt.Println("========================== Get exchange value by date test ==========================")
	res4, err := client.GetExchangeRateByDate("UAH", []string{"EUR", "USD"}, "2008-12-07")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(res4.Info())

	// convert currencies test
	fmt.Println("========================== Convert currencies test ==========================")
	res5, err := client.Convert("UAH", "USD", 20)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(res5.Info())
}
