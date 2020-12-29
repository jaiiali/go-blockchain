package db

import (
	"fmt"
	"time"
)

type SingedTx struct {
	Tx
	Sig string `json:"signature"`
}

type Tx struct {
	From         Account `json:"from"`
	To           Account `json:"to"`
	Value        float64 `json:"value"`
	AccountNonce uint64  `json:"nonce"`
	Data         string  `json:"data"`
	Time         int64   `json:"timestamp"`
}

func NewTx(from Account, to Account, value float64, nonce uint64, data string) *Tx {
	return &Tx{
		From:         from,
		To:           to,
		Value:        value,
		AccountNonce: nonce,
		Data:         data,
		Time:         time.Now().UnixNano(),
	}
}

func (tx *Tx) String() string {
	return fmt.Sprintf("%+v", *tx)
}
