package stats

import "github.com/ilhom9045/bank/v2/pkg/types"

func Avg(payments []types.Payment) (money types.Money) {
	count := types.Money(0)
	for _, value := range payments {
		if value.Status != types.StatusFail {
			count++
			money += value.Amount
		}
	}
	return money / count
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	money := types.Money(0)

	for _, value := range payments {

		if value.Category == category && value.Status != types.StatusFail {

			money += value.Amount

		}
	}

	return money
}
