package main

import (
	"fmt"
	"os"

	"github.com/brothergiez/assessment/anagram"
	"github.com/brothergiez/assessment/bookingsystem"
	"github.com/brothergiez/assessment/calculate"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [anagram|booking]")
		return
	}

	command := os.Args[1]

	switch command {
	case "anagram":
		words := []string{"listen", "silent", "enlist", "google", "gogole", "cat", "act"}
		result := anagram.GroupAnagram(words)
		fmt.Println(result)

	case "booking":
		system := bookingsystem.NewBookingSystem()

		err := system.AddRoom("room1", 5)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Room 'room1' with capacity 5 added successfully.")
		}

		message, err := system.Book("room1", "2023-01-08 10:00", 2, 2)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(message)
		}

		message, err = system.Book("room1", "2023-01-08 11:00", 1, 1)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(message)
		}

	case "calculate":
		items := []map[string]interface{}{
			{"name": "Apple", "price": "10", "quantity": 2},
			{"name": "Orange", "price": "5", "quantity": 3},
		}

		total, err := calculate.CalculateTotal(items)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Total Price:", total)

	default:
		fmt.Println("Unknown command. Use 'anagram' or 'booking'.")
	}
}
