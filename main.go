package main

import (
	"bank/pkg/bank/card"
	"bank/pkg/bank/types"
	"fmt"
)

func main() {


	cards := []types.Card{
		{
		PAN: "312312321",
		Balance: 10000,
		Active: true,
	},
	{
		PAN: "312312321",
		Balance: 10000,
		Active: true,
	},
	{
		PAN: "312312321",
		Balance: 10000,
		Active: true,
	},
	}

	sources := card.PaymentSources(cards)


	for _, value := range sources{
		fmt.Println(value.Number, "number:")
	}



}