package main

import "fmt"

func reverseLeftAnglePyramid() {
	for i := 5; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
