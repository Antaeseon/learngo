package main

import (
	"fmt"

	"github.com/antaeseon/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	err := dictionary.Add("hello", "Greeting")
	if err != nil {
		fmt.Println(err)
	}
	definition, err := dictionary.Search("hello")
	fmt.Println(definition, err)

}
