package db

import (
	"encoding/json"
	"fmt"
)

type State struct {
	Balances  map[Account]float32
	txMempool []Tx

	db string
}

func NewState() (*State, error) {
	gen, err := loadGenesis()
	if err != nil {
		return nil, err
	}

	balances := make(map[Account]float32)
	for account, balance := range gen.Balances {
		balances[account] = balance
	}

	state := &State{
		Balances:  balances,
		txMempool: make([]Tx, 0),
		db:        "",
	}

	return state, err
}

func (s *State) Add(tx *Tx) error {
	s.txMempool = append(s.txMempool, *tx)

	return nil
}

func (s *State) Store() error {
	for i, tx := range s.txMempool {
		txJson, err := json.Marshal(tx)
		fmt.Println(string(txJson))
		if err != nil {
			return err
		}

		s.db += string(txJson)
		//s.db = append(s.db, txJson...)
		s.txMempool = append(s.txMempool[:i], s.txMempool[i+1:]...)
	}

	return nil
}
