package ReadySteadyGo

import "fmt"

func main() {

	if 5%2 == 0 {
		fmt.Println("it is pair")
	} else {
		fmt.Println("it is not pair")
	}

	if 7%4 == 0 {
		fmt.Println("8 % 4")
	} else {
		fmt.Println("!% 4")
	}

	if num := -9; num < 0 {
		fmt.Println("negative int")
	} else if num < 10 {
		fmt.Println("it has 1 int")
	}
}