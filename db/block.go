package db

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type HashType [32]byte

var difficulty = 2

type Block struct {
	Hash       HashType `json:"hash"`
	ParentHash HashType `json:"parent_hash"`
	Height     int64    `json:"height"`
	Difficulty int      `json:"difficulty"`
	Nonce      int64    `json:"nonce"`
	Time       int64    `json:"timestamp"`
	Txs        []*Tx    `json:"tx"`
}

func NewBlock(parentBlock *Block, txs []*Tx) (*Block, error) {
	var parentHash HashType
	var height int64

	if parentBlock == nil {
		parentHash = HashType{}
		height = 1
	} else {
		parentHash = parentBlock.Hash
		height = parentBlock.Height + 1
	}

	b := &Block{
		ParentHash: parentHash,
		Height:     height,
		Difficulty: difficulty,
		Nonce:      0,
		Time:       time.Now().UnixNano(),
		Txs:        txs,
	}

	// Calc nonce
	var elaspse time.Time = time.Now()
	var attempt int64 = 0
	for {
		b.Nonce = attempt
		b.Hash = b.CalcHash()
		if isHashValid(b.Hash, difficulty) {
			break
		}

		attempt++
	}
	defer fmt.Printf("\nElaspled: %s\n", time.Since(elaspse))

	fmt.Printf("\nAttempts: %d", attempt)

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
	record := fmt.Sprintf("%#x %#x %#x %#x %p", b.ParentHash, b.Height, b.Nonce, b.Time, b.Txs)
	//spew.Dump(record)
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

func isHashValid(hash HashType, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(fmt.Sprintf("%x", hash), prefix)
}
