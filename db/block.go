package db

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Hash [32]byte

type Block struct {
	//Header    Header    `json:"header"`
	Timestamp time.Time `json:"timestamp"`
	Parent    Hash      `json:parent`
	Tx        *[]Tx     `json:"tx"`
}

type Header struct {
	Parent    Hash      `json:"parent"`
	Height    uint64    `json:"number"`
	Nonce     uint32    `json:"nonce"`
	Timestamp time.Time `json:"timestamp"`
}

func NewBlock(txs *[]Tx) *Block {
	return &Block{
		Timestamp: time.Now(),
		Tx:        txs,
	}
}

func (b *Block) String() string {
	return fmt.Sprintf("%#v\n\n%#v", *b, b.Tx)
}

func (b *Block) Hash() error {
	blockJson, err := json.Marshal(b)
	if err != nil {
		return err
	}

	b.Parent = sha256.Sum256(blockJson)
	return nil
}

func (b *Block) IsBlockHashValid() bool {
	return true
}
