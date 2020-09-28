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
func CategoriesAvg(payments []types.Payment)  map[types.Category]types.Money {
	n := make(map[types.Category]types.Money)
	// categories := map[types.Category]types.Money{}
	for _, payment := range payments {
		if _, er := n[payment.Category]; er {
			continue
		}
		filtered := FilterByCategory(payments, payment.Category)
		n[payment.Category]=Avg(filtered)
	}
	return n 
}
