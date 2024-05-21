package main

import "fmt"

func floydsTriangle() {
	rows := 5

	k := 1
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(k, " ")
			k++
		}
		fmt.Println()
	}

}
