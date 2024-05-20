package main

import "fmt"

func leftAngleTriangle() {
	rows := 5
	for i := 1; i <= rows; i++ {
		for k := rows; k > i; k-- {
			fmt.Print("  ")
		}

		for j := 1; j <= i; j++ {
			//add space to align the pyramid
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
