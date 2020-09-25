package stats

import (
	"github.com/usmon1983/bank/v2/pkg/types"

)

//Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) (avg types.Money) {
	n := types.Money(0)
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			n++
			avg += (payment.Amount)
		}
	}
	avg = avg/n
	return avg
}

//TotalInCategory
func TotalInCategory(payments []types.Payment, category types.Category) types.Money  {
	var sum types.Money
	for _, payment := range payments{
		if payment.Status != types.StatusFail {
			if payment.Category == category {
				sum += payment.Amount
			}
		}
	}
	return sum
}