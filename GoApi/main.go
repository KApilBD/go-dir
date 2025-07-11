package main

import (
	"fmt"

	"mf.com/go/crypto/api"
)

func main() {
	rate, err := api.GetRate("BTC")
	if err == nil {
		fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)
	}
}
