package main

import (
	"fmt"

	"github.com/jaiiali/go-simple-blockchain/db"
)

func main() {
	//gen, _ := db.LoadGenesis()
	//fmt.Printf("%#v\n\n", gen)

	from := db.NewAccount("a")
	to := db.NewAccount("b")
	tx := db.NewTx(from, to, 100.0, 0, "")
	fmt.Println(tx)

	state, err := db.NewState()
	if err != nil {
		panic(err)
	}

	err = state.Add(tx)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%#v\n\n", state)

	err = state.Store()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%#v\n\n", state)
}
