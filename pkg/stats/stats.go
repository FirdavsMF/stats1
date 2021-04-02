package stats

import (
	"github.com/FirdavsMF/bank1/v2/pkg/types"
	
)

//Avg расчитовает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	// countPayments := types.Money(len(payments))
	sumPaymenys := types.Money(0)
	indexPayments := types.Money(0)
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		moneyPayments := payment.Amount
		sumPaymenys += moneyPayments
		indexPayments++
	}
	return sumPaymenys / indexPayments
}

// TotalInCategory находит 	сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sumPaymenys := types.Money(0)
	for _, payment := range payments {
		if payment.Category != category {
			continue
		}
		if payment.Status == types.StatusFail {
			continue
		}
		moneyPayments := payment.Amount
		sumPaymenys += moneyPayments
	}
	return sumPaymenys
}
//CategoriesAvg расчитовает среднюю сумму платежа
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	count := map[types.Category]types.Money{}
	result := map[types.Category]types.Money{}

	for _, payment := range payments {
		result[payment.Category] += payment.Amount
		count[payment.Category]++
	}

	for key := range result {
		result[key] /= count[key]
	}

	return result
}