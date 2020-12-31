package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/jaiiali/go-simple-blockchain/db"
)

var storage *db.BlockChain

func init() {
	storage = db.NewBlockChain()
}

func main() {
	var from, to db.Account

	// first transaction
	from = db.NewAccount("a")
	to = db.NewAccount("b")
	tx1 := db.NewTx(from, to, 100.0, 1, "")

	// second transaction
	from = db.NewAccount("c")
	to = db.NewAccount("d")
	tx2 := db.NewTx(from, to, 200.0, 2, "")

	// third transaction
	from = db.NewAccount("c")
	to = db.NewAccount("d")
	tx3 := db.NewTx(from, to, 200.0, 2, "")

	var txs []*db.Tx
	var b *db.Block

	// first block
	txs = []*db.Tx{tx1}
	b, _ = db.NewBlock(storage.Blocks[storage.LastHeight-1], txs)
	_ = storage.AddBlock(b)

	// second block
	txs = []*db.Tx{tx2, tx3}
	b, _ = db.NewBlock(storage.Blocks[storage.LastHeight-1], txs)
	_ = storage.AddBlock(b)

	spew.Dump(storage)
	log.Fatal(run())
}
