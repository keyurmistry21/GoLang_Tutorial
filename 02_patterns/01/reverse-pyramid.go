package main

func reversePyramid() {
	rows := 5
	for i := 0; i <= rows; i++ {
		for k := 1; k <= i; k++ {
			print(" ")
		}
		for j := 1; j <= rows-i; j++ {
			print(j, " ")
		}
		println()
	}
}
