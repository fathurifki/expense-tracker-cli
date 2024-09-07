package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	// "strconv"
	// "text/scanner"

	// "strconv"
	// "strings"
	"time"
)

type Expense struct {
	ID          int
	Date        string
	description string
	amount      float64
}

func NewExpense(id int, description string, amount float64) Expense {
	now := time.Now()
	formattedDate := now.Format("2006-01-02")
	return Expense{
		ID:          id,
		Date:        formattedDate,
		description: description,
		amount:      amount,
	}
}

func main() {
	terminal_reader := bufio.NewScanner(os.Stdin)
	inputID := 1
	expense_list := make(map[int]Expense)

	for {
		fmt.Println("expense-tracker> ")
		terminal_reader.Scan()
		input := terminal_reader.Text()
		parts := strings.Fields(input)
		command := parts[0]

		switch command {
		case "add":
			if len(parts) > 2 && strings.HasPrefix(parts[1], "-d") || strings.HasPrefix(parts[1], "--description") {
				description := parts[2]
				amount, err := strconv.ParseFloat(parts[4], 64)
				if err != nil {
					fmt.Println("invalid amount")
				}

				expense := NewExpense(inputID, description, amount)
				expense_list[expense.ID] = expense
				inputID++
				fmt.Printf("Added expense: %s, $%.2f\n", description, amount)
			} else {
				fmt.Println("Invalid Command")
			}
		case "list":
			fmt.Println("\n# ID  Date       Description  Amount")
			fmt.Println("------------------------------------")
			for _, expense := range expense_list {
				fmt.Printf("%3d  %s  %-12s $%8.2f\n",
					expense.ID,
					expense.Date,
					expense.description,
					expense.amount)
			}
			fmt.Println()
		case "summary":
			var summary_expense float64

			if len(parts) > 2 && strings.HasPrefix(parts[1], "--month") {
				selectedDate, err := strconv.Atoi(parts[2])
				if err != nil {
					fmt.Printf("Summary with ID %d not found.\n", selectedDate)
				}
				for _, expense := range expense_list {
					expenseDate, _ := time.Parse("2006-01-02", expense.Date)
					if err != nil {
						fmt.Printf("Error parsing date for expense ID %d: %v\n", expense.ID, err)
						continue
					}

					if int(expenseDate.Month()) == selectedDate {
						summary_expense += expense.amount
					}

					fmt.Printf("There is no expense at %d month \n", selectedDate)
				}
			} else {
				for _, expense := range expense_list {
					summary_expense += expense.amount
				}
				fmt.Println(summary_expense)
			}

		case "delete":
			if len(parts) > 2 && strings.HasPrefix(parts[1], "--id") {
				deletedId, err := strconv.Atoi(parts[2])
				if err != nil {
					fmt.Printf("Expense with ID %d not found.\n", deletedId)
					continue
				}
				delete(expense_list, deletedId)
				fmt.Printf("Expense Id %d has been delete \n", deletedId)
			}
		}

	}
}
