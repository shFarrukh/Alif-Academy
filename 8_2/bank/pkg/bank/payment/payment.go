package payment

import "bank/pkg/bank/types"

func Max(cardfs []types.Payment) types.Payment{
	mx:= cardfs[0].Amount
	pymnt := types.Payment{}
	for _, i := range cardfs{
		if (i.Amount>mx) {
			mx=i.Amount
			pymnt = i
		}
	}
	return pymnt
}