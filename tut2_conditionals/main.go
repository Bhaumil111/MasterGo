package main

import "fmt" 



func Greetings(){
	fmt.Println("Use Profit Calculator")
}

func main() {
	Greetings() 
	var revenue, expenses, taxRate float64
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses

	profit := (revenue - expenses) / 100

	fmt.Println("Ebt " , ebt)

	fmt.Println("Profit ", profit)
	fmt.Printf("Profit %.1f", profit)

}
