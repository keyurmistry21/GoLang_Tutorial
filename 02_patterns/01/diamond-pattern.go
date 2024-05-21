package main

func diamondPattern() {
	rows := 5
	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {
			print(" ")
		}
		for j := 1; j <= i; j++ {
			print(j, " ")
		}
		println()
	}
	for i := 1; i <= rows; i++ {
		for j := i; j >= 1; j-- {
			print(" ")
		}
		for j := 1; j <= rows-i; j++ {
			print(j, " ")
		}
		println()
	}
}
