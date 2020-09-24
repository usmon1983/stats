package stats

import (
	"github.com/usmon1983/myDZ/pkg/bank/types"
)

//Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	sum := 0
	next := 0
	for _, payment := range payments {
		sum += int(payment.Amount)
		next += 1
	}
	AVG := sum/next
	return types.Money(AVG)
}

//TotalInCategory
func TotalInCategory(payments []types.Payment, category types.Category) types.Money  {
	sum := 0
	next := 0
	for _, payment := range payments {
		if payment.Category == category {
			sum += int(payment.Amount)
			next += 1
		}
	}
	//AVG := sum/next
	return types.Money(sum)
}