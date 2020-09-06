package main

import (
	"fmt"

	"github.com/antaeseon/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("ts")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
}
