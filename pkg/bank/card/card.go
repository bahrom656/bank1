package card

import "bank/pkg/bank/types"

func Total(cards []types.Card) types.Money{
	sum := types.Money(0)

	for _, card := range cards{
		if !card.Active{
			continue
		}

		if card.Balance <= 0 {
			continue
		}

		sum += card.Balance
	}
	return sum
}

func PaymentSources(cards []types.Card) []types.PaymentSource{
	var source []types.PaymentSource
	var indexSource int
	source = make([]types.PaymentSource, len(cards))
	for index, card := range cards {
		if card.Active == true && card.Balance > 0 {
			source[indexSource].Number = string(cards[index].PAN)
			source[indexSource].Balance = cards[index].Balance
			source[indexSource].Type = "card"
			indexSource++
		}
	}
	return source
}