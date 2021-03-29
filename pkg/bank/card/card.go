package card

import (
	"bank/pkg/bank/types"
)


func Total(cards []types.Card) types.Money {
	
	sum := types.Money(0)

	for _, card := range cards{

		if !card.Active{
			continue
		} 
		if card.Balance <= 0{
			continue
		}
		sum += card.Balance
		
	}

	return sum
	
}



func PaymentSources(cards []types.Card)[]types.PaymentSource {
		payments := []types.PaymentSource{}


		for _, card := range cards{
			if card.Balance <= 0{
				continue
			}
			if !card.Active{
				continue
			}

			payment := types.PaymentSource{
				Type: "card",
				Number: string(card.PAN),
				Balance: card.Balance,
			}	

			payments = append(payments, payment)
			

		}

		return payments
	
}