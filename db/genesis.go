package db

import (
	"encoding/json"
)

var genesisData = []byte(`
{
  "genesis_time": "2020-12-28T03:30:00.000000000Z",
  "chain_id": "simple-blockchain",
  "balances": {
    "a": 1000000.0
  }
}`)

type genesis struct {
	Balances map[Account]float32 `json:"balances"`
}

func genesisBlock() *Block {
	b, _ := NewBlock(nil, nil)

	return b
}

func loadGenesis() (genesis, error) {
	var gen genesis
	err := json.Unmarshal(genesisData, &gen)
	if err != nil {
		return genesis{}, err
	}

	return gen, nil
}
