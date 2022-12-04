package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello there!")
	Login()
}

func Login() {
	var login string
	var checker int = 0
	login = "Deneme"
	for checker < 5 {
		fmt.Println("Please enter your login name: ")
		fmt.Scanf("%s", &login)
		fmt.Scanln()
		if login == "brucewayne" {
			fmt.Printf("Welcome Master Bruce, what do you wish to do today?\n")
			operator_incognito()
			return
		}
		if 4-checker == 0 {
			return
		}
		fmt.Printf("Unrecognized login. Acces is denied. You have %d attempt(s) remaining.\n", 4-checker)
		checker++
	}
}

func operator_incognito() {
	var operation int
	fmt.Println("1. List all even numbers until 100.")
	fmt.Println("2. List all odd numbers until 100.")
	fmt.Println("3. Take square of the entered number.")
	fmt.Scanf("%d", &operation)
	fmt.Scanln()
	if operation == 1 {
		begin_even()
	}
	if operation == 2 {
		begin_odd()
	}
	if operation == 3 {
		begin_square()
	}
}

func begin_even() {
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("%d is an even number\n", i)
		}
	}
	operator_incognito()
}

func begin_odd() {
	for i := 0; i <= 100; i++ {
		if i%2 == 1 {
			fmt.Printf("%d is an odd number\n", i)
		}
	}
	operator_incognito()
}
func begin_square() {
	var x int
	fmt.Println("Please enter the number to take square: ")
	fmt.Scanf("%d", &x)
	fmt.Scanln()
	if x == 253 {
		time.Sleep(2 * time.Second)

		batcave()
		return
	}
	fmt.Printf("Square of %d = %d\n", x, x*x)
	operator_incognito()
}
func batcave() {
	fmt.Println("Welcome to batcave Master Bruce.")
	fmt.Println("What do you wish to do today?")
	operator_batcave()
}
func operator_batcave() {
	var a int
	fmt.Println("1. Access to archives.")
	fmt.Scanf("%d", &a)
	fmt.Scanln()
	if a == 1 {
		archives()
		return
	}
}
func archives() {
	var a int
	fmt.Println("Please choose to topic you want to explore.")
	fmt.Println("Topics:")
	fmt.Println("1. Arkham Asylum")
	fmt.Scanf("%d", &a)
	fmt.Scanln()
	if a == 1 {
		fmt.Println("There is no data to show here.")
		archives()
	}
	fmt.Println("Unrecognized input please try again.")
	archives()
}
