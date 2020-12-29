package main

import (
	"github.com/davecgh/go-spew/spew"

	"github.com/jaiiali/go-simple-blockchain/db"
)

func main() {
	//gen, _ := db.LoadGenesis()
	//fmt.Printf("%#v\n\n", gen)

	var from, to db.Account

	from = db.NewAccount("a")
	to = db.NewAccount("b")
	tx1 := db.NewTx(from, to, 100.0, 1, "")
	//fmt.Println(tx1)

	from = db.NewAccount("c")
	to = db.NewAccount("d")
	tx2 := db.NewTx(from, to, 200.0, 2, "")
	//fmt.Println(tx2)

	//
	//state, err := db.NewState()
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = state.Add(tx)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Printf("%#v\n\n", state)
	//
	//err = state.Store()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Printf("%#v\n\n", state)

	var txs []*db.Tx
	var b *db.Block

	bc := db.NewBlockChain()

	txs = []*db.Tx{tx1}
	b = db.NewBlock(txs)
	_ = bc.AddBlock(b)

	txs = []*db.Tx{tx2}
	b = db.NewBlock(txs)
	_ = bc.AddBlock(b)

	spew.Dump(bc)
}
