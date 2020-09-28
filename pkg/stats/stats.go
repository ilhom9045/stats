package stats

import (
	"github.com/ilhom9045/bank/v2/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	money := types.Money(0)
	count := types.Money(len(payments))

	for _, value := range payments {
		money += value.Amount
	}

	money = money / count
	return money
}

//FilterByCategory
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}
	return filtered
}

//CategoriesAvg, которая считает среднюю сумму платежа по каждой категории
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	n := make(map[types.Category]types.Money)
	// categories := map[types.Category]types.Money{}
	for _, payment := range payments {
		if _, er := n[payment.Category]; er {
			continue
		}
		filtered := FilterByCategory(payments, payment.Category)
		n[payment.Category] = Avg(filtered)
	}
	return n
}

func PeriodsDynamic(
	first map[types.Category]types.Money,
	second map[types.Category]types.Money,
) map[types.Category]types.Money {

	periods := map[types.Category]types.Money{}

	if len(first) > len(second) {
		for value := range first {
			periods[value] = second[value] - first[value]
		}
		return periods
	}
	for value := range second {
		periods[value] = second[value] - first[value]
	}
	return periods

}
