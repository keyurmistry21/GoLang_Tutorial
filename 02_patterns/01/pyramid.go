package main

import "fmt"

func pyramid() {
	rows := 5
	for i := 1; i <= rows; i++ {
		for k := rows; k > i; k-- {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
