package db

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type HashType [32]byte

type Block struct {
	Hash       HashType `json:"hash"`
	ParentHash HashType `json:"parent_hash"`
	Height     int64    `json:"height"`
	Time       int64    `json:"timestamp"`
	Txs        []*Tx    `json:"tx"`
}

func NewBlock(parentBlock *Block, txs []*Tx) (*Block, error) {
	b := &Block{
		ParentHash: parentBlock.Hash,
		Height:     parentBlock.Height + 1,
		Time:       time.Now().UnixNano(),
		Txs:        txs,
	}

	b.Hash = b.CalcHash()

	return b, nil
}

func (b *Block) String() string {
	return fmt.Sprintf("%#v\n\n%#v", *b, b.Txs)
}

func (b *Block) MarshalJSON() ([]byte, error) {
	type Alias Block
	block := struct {
		Hash       string `json:"hash"`
		ParentHash string `json:"parent_hash"`
		*Alias
	}{
		Hash:       fmt.Sprintf("%#x", b.Hash),
		ParentHash: fmt.Sprintf("%#x", b.ParentHash),
		Alias:      (*Alias)(b),
	}

	return json.Marshal(block)
}

func (b *Block) CalcHash() HashType {
	record := fmt.Sprintf("%#x %#x %#x %p", b.ParentHash, b.Height, b.Time, b.Txs)
	spew.Dump(record)
	h := sha256.Sum256([]byte(record))

	return h
}

func (b *Block) IsBlockHashValid() bool {
	h := b.CalcHash()
	if h != b.Hash {
		return false
	}

	return true
}
