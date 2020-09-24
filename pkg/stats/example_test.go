package stats

import (
	"github.com/usmon1983/myDZ/pkg/bank/types"
	"fmt"
)

func ExampleAvg()  {
	payments := []types.Payment {
		{
			ID: 1,
			Amount: 10_000_00,
			Category: "Рестораны",
		},
		{
			ID: 2,
			Amount: 20_000_00,
			Category: "Рестораны",
		},
		{
			ID: 3,
			Amount: 30_000_00,
			Category: "Рестораны",
		},
	}
	PaymentAvg := Avg(payments)

	fmt.Println(PaymentAvg)

	//Output: 2000000
}

func ExampleTotalInCategory(){
	payments := []types.Payment{
	  {
		ID:			1,
		Amount:		93_000,
		Category: 	"Самбусахона",
	  },
	  {
		ID:			2,
		Amount:		87_000,
		Category: 	"Самбусахона",
	  },
	  {
		ID:			3,
		Amount:		12_000,
		Category: 	"Мантухона",
	  },
	}
  
	fmt.Println(TotalInCategory(payments, "Мантухона"))
  
	//Output: 12000
  }