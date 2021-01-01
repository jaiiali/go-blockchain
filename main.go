package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/jaiiali/go-simple-blockchain/db"
	"github.com/joho/godotenv"
)

var storage *db.Blockchain

// bcServer handles incoming concurrent Blocks
var bcServer chan []db.Block
var mutex = &sync.Mutex{}

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

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	bcServer = make(chan []db.Block)

	tcpPort := os.Getenv("PORT")

	// start TCP and serve TCP server
	server, err := net.Listen("tcp", ":"+tcpPort)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("TCP  Server Listening on port :", tcpPort)
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()

	msg := "Enter a new transaction (from,to,value): "

	io.WriteString(conn, msg)
	scanner := bufio.NewScanner(conn)

	// take in Tx from stdin and add it to blockchain after conducting necessary validation
	go func() {
		for scanner.Scan() {

			rawTx := strings.Split(scanner.Text(), ",")
			if len(rawTx) < 3 {
				log.Printf("input is not correct")
				continue
			}

			from := db.NewAccount(strings.TrimSpace(rawTx[0]))
			to := db.NewAccount(strings.TrimSpace(rawTx[1]))
			valueStr := strings.TrimSpace(rawTx[2])
			value, err := strconv.ParseFloat(valueStr, 64)
			if err != nil {
				log.Printf("%v not a number: %v", scanner.Text(), err)
				continue
			}

			rand.Seed(time.Now().UnixNano())
			nonce := uint64(rand.Intn(100))

			tx := db.NewTx(from, to, value, nonce, "")
			txs := []*db.Tx{tx}
			b, err := db.NewBlock(storage.Blocks[storage.LastHeight-1], txs)
			if err != nil {
				log.Print(fmt.Errorf("block could not be generate: %w", err).Error())
				continue
			}

			err = storage.AddBlock(b)
			if err != nil {
				log.Print(fmt.Errorf("block could not be add: %w", err).Error())
				continue
			}

			spew.Dump(b)
			io.WriteString(conn, "\n"+msg)
		}
	}()

	// simulate receiving broadcast
	go func() {
		for {
			time.Sleep(30 * time.Second)
			output, err := json.Marshal(storage)
			if err != nil {
				log.Fatal(err)
			}
			io.WriteString(conn, string(output))
		}
	}()

	for _ = range bcServer {
		spew.Dump(storage)
	}

}
