package main

import (
	"fmt"
)

func calculateExpenseSharing(expenses map[string]float64) map[string]float64 {
	totalExpenses := 0.0
	for _, expense := range expenses {
		totalExpenses += expense
	}

	averageExpense := totalExpenses / float64(len(expenses))

	expenseSharing := make(map[string]float64)
	for person, expense := range expenses {
		expenseSharing[person] = averageExpense - expense
	}

	return expenseSharing
}

func main() {
	expenses := map[string]float64{
		"John":  50.0,
		"Mary":  75.0,
		"Peter": 30.0,
		"Alice": 60.0,
	}

	expenseSharing := calculateExpenseSharing(expenses)

	for person, amount := range expenseSharing {
		fmt.Printf("%s: %.2f\n", person, amount)
	}
}
