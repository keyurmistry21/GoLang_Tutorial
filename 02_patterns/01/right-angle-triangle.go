package main

import "fmt"

func rightAngleTriangle() {
	rows := 5
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			//add space to align the pyramid
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
