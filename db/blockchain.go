package db

import "errors"

type Blockchain struct {
	Blocks     []*Block
	LastHeight int64
}

func NewBlockChain() *Blockchain {
	bc := &Blockchain{
		Blocks:     []*Block{genesisBlock()},
		LastHeight: 1,
	}

	return bc
}

func (bc *Blockchain) AddBlock(b *Block) error {
	// ParentHash
	lastBlock := bc.Blocks[bc.LastHeight-1]

	if lastBlock.Height+1 != b.Height {
		return errors.New("height is now correct")
	}

	if lastBlock.Hash != b.ParentHash {
		return errors.New("parentHash is now correct")
	}

	if !b.IsBlockHashValid() {
		return errors.New("block hash is now correct")
	}

	bc.Blocks = append(bc.Blocks, b)
	bc.LastHeight++
	return nil
}
