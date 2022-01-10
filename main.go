package main

import "fmt"

func getInput() {
	var x, y int
	var operator string

	fmt.Print("输入x: ")
	fmt.Scanf("%d\n", &x)

	fmt.Print("输入[ + - * / ]号: ")
	fmt.Scanf("%s\n", &operator)

	fmt.Print("输入y: ")
	fmt.Scanf("%d\n", &y)

	switch operator {
	case "+":
		fmt.Println(x + y)
	case "-":
		fmt.Println(x - y)
	case "*":
		fmt.Println(x * y)
	case "/":
		fmt.Println(x / y)

	}

}

func main() {
	for {
		getInput()
	}
}
