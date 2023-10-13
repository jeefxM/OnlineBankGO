package main

import (
	"errors"
	"fmt"
)

type BinanceAccount struct {
	BTC int
	ETH int
	XRP int
}

func NewBinanceAccount() *BinanceAccount {
	return &BinanceAccount{
		BTC: 0,
		ETH: 0,
		XRP: 0,
	}
}

func (b *BinanceAccount) getBalance() string {
	return fmt.Sprintf("BTC balance: %d, ETH balance: %d, XRP balance: %d", b.BTC, b.ETH, b.XRP)
}

type Currency string

const (
	BTC Currency = "BTC"
	ETH Currency = "ETH"
	XRP Currency = "XRP"
)

func (b *BinanceAccount) deposit(amount int, currency Currency ) {
	switch currency {
	case BTC:
		b.BTC += amount
	case ETH:
		b.ETH += amount
	case XRP:
		b.XRP += amount
	default:
		fmt.Println("Unknown currency type.")
	}
}

func(b *BinanceAccount) withdraw(amount int, currency Currency) error{
	insufficentFunds := errors.New("Insufficent funds")
	switch currency{
	case BTC:
		if b.BTC < amount{
			return insufficentFunds
		}
		b.BTC = b.BTC - amount
	case ETH:
		if b.ETH < amount{
			return insufficentFunds
		}
		b.ETH = b.ETH - amount
	case XRP:
		if b.XRP < amount{
			return insufficentFunds
		}
		b.XRP = b.XRP - amount
	}
	return nil
}

