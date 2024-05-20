package main

import "fmt"

func reverseRightAngleTriangle() {
	rows := 5
	for i := rows; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
