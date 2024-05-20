package main

import "fmt"

func reverseLeftAngleTriangle() {
	rows := 5
	for i := rows; i >= 1; i-- {
		for k := rows; k > i; k-- {
			fmt.Print("  ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
