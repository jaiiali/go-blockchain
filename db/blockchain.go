package db

import "time"

type BlockChain struct {
	Blocks []*BlockWithHash
}

func NewBlockChain() *BlockChain {
	bh := &BlockWithHash{
		Hash: HashType{},
		Data: Block{
			Header: BlockHeader{
				Time: time.Now().UnixNano(),
			},
		},
	}

	bc := &BlockChain{
		Blocks: []*BlockWithHash{bh},
	}

	return bc
}

func (bc *BlockChain) AddBlock(b *Block) error {
	// ParentHash
	height := len(bc.Blocks)
	b.Header.ParentHash = bc.Blocks[height-1].Hash

	// Hash
	h, err := b.Hash()
	if err != nil {
		return err
	}

	// BlockWithHash
	bh := &BlockWithHash{
		Hash: h,
		Data: *b,
	}

	bc.Blocks = append(bc.Blocks, bh)
	return nil
}
