package stats

import "github.com/ilhom9045/bank/v2/pkg/types"

func Avg(payments []types.Payment) types.Money {
	money := types.Money(0)
	count := types.Money(len(payments))

	for _, value := range payments {
		money += value.Amount
	}

	money = money / count
	return money
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
