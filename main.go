package main

import (
	"fmt"

	"github.com/antaeseon/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("ts")
	fmt.Println(account)
}
