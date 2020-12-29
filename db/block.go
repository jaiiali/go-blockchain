package db

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type HashType [32]byte

type BlockWithHash struct {
	Hash HashType `json:"hash"`
	Data Block    `json:"block"`
}

type Block struct {
	Header BlockHeader `json:"header"`
	Txs    []*Tx       `json:"tx"`
}

type BlockHeader struct {
	ParentHash HashType `json:"parent"`
	Height     uint64   `json:"number"`
	Nonce      uint32   `json:"nonce"`
	Time       int64    `json:"timestamp"`
}

func NewBlock(txs []*Tx) *Block {
	return &Block{
		Header: BlockHeader{
			ParentHash: HashType{},
			Height:     0,
			Nonce:      0,
			Time:       time.Now().UnixNano(),
		},
		Txs: txs,
	}
}

func (b *Block) String() string {
	return fmt.Sprintf("%#v\n\n%#v", *b, b.Txs)
}

func (b *Block) Hash() (HashType, error) {
	blockJson, err := json.Marshal(b)
	if err != nil {
		return HashType{}, err
	}

	h := sha256.Sum256(blockJson)

	return h, nil
}

func (b *Block) IsBlockHashValid() bool {
	return true
}
