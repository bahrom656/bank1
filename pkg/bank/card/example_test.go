package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 100000,
			Active: true,
		},
	}
	fmt.Println(Total(cards))
	//Output: 100000
}
