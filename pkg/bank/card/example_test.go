package card

import (
	"bank/pkg/bank/types"
	"fmt"
)


func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 1_000_00,
			Active: true,

		},
	}

	fmt.Println(Total(cards))

	// Output: 
	// 100000
}




func ExamplePaymentSources()  {
	cards := []types.Card{
		{
			PAN: "5058 xxxx xxxx 8888",
			Balance: 10000,
			Active: true,
		},
		{
			PAN: "5058 xxxx xxxx 8889",
			Balance: 10000,
			Active: true,
		},
		{
			PAN: "5058 xxxx xxxx 8890",
			Balance: 10000,
			Active: true,
		},

	}
	payments := PaymentSources(cards)

	for _, payment := range payments{
		fmt.Println(payment.Number)
	}

	// Output:
	// 5058 xxxx xxxx 8888
	// 5058 xxxx xxxx 8889
	// 5058 xxxx xxxx 8890
	
}