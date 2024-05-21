package main

import "fmt"

func pascalsTriangle() {
	rows := 5
	for i := 0; i < rows; i++ {
		val := 1
		for k := 0; k <= rows-i; k++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print(val, " ")
			val = val * (i - j) / (j + 1)
		}

		fmt.Println()
	}

}
