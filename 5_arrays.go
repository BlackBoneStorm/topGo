package main

import "fmt"

func main()  {

	var a [5]int
	var x [10]string
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	x[9] = "lol"
	fmt.Println(x[9])
	fmt.Println("len x:", len(x))
	fmt.Println(x)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2nd: ", twoD)
}
