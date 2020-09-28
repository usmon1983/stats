package stats

import (
	"github.com/usmon1983/bank/v2/pkg/types"
	"fmt"
)

func ExampleAvg()  {
	payments := []types.Payment {
		{
			ID: 1,
			Amount: 10_000_00,
			Category: "Рестораны",
			Status: types.StatusOk,
		},
		{
			ID: 2,
			Amount: 20_000_00,
			Category: "Рестораны",
			Status: types.StatusOk,
		},
		{
			ID: 3,
			Amount: 30_000_00,
			Category: "Рестораны",
			Status: types.StatusFail,
		},
	}
	PaymentAvg := Avg(payments)

	fmt.Println(PaymentAvg)

	//Output: 1500000
}

func ExampleTotalInCategory(){
	payments := []types.Payment{
	  {
		ID:			1,
		Amount:		93_000,
		Category: 	"Самбусахона",
		Status: 	types.StatusFail,
	  },
	  {
		ID:			2,
		Amount:		87_000,
		Category: 	"Самбусахона",
		Status: 	types.StatusFail,
	  },
	  {
		ID:			3,
		Amount:		12_000,
		Category: 	"Мантухона",
		Status: 	types.StatusOk,
	  },
	}
	
	fmt.Println(TotalInCategory(payments, "Мантухона"))
  
	//Output: 12000
  }