package main

import (
	"fmt"
)

type bankAccount interface{
	getBalance() int
	deposit(amount int)
	withdraw(amount int)
}

func main() {
	BinanceAccount := NewBinanceAccount()
	bogAccount := newBogAccount()

	BinanceAccount.deposit(10,BTC)
	BinanceAccount.withdraw(5, BTC)
	BinanceAccount.deposit(3, ETH)
	BinanceBalance := BinanceAccount.getBalance()
	fmt.Print(BinanceBalance)

	bogAccount.deposit(10)
	BogBalance := bogAccount.getBalance()
	fmt.Printf("\nBog account balance is: %d\n", BogBalance)
	bogAccount.withdraw(5)
	BogBalance = bogAccount.getBalance()
	fmt.Printf("\n Bog balance after withdraw is %d\n", BogBalance)
}