package main

import (
	"errors"
)

type BOG struct{
	balance int
}

func newBogAccount() *BOG{
	return &BOG{
		balance: 0,
	}
}

func (b *BOG) getBalance() int{
	return b.balance
}

func (b *BOG) deposit(amount int){
	b.balance += amount
}

func (b *BOG) withdraw(amount int) error{
	newBalance := b.balance - amount
	if newBalance < 0{
		return errors.New("insufficient fundss")
	}
	b.balance = newBalance
	return nil
}