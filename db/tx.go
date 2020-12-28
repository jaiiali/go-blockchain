package db

import (
	"fmt"
	"time"
)

//
//type SingedTx struct {
//	Tx
//	Sig string `json:"signature"`
//}

type Tx struct {
	From         Account   `json:"from"`
	To           Account   `json:"to"`
	Value        float32   `json:"value"`
	AccountNonce uint64    `json:"nonce"`
	Data         string    `json:"data"`
	Timestamp    time.Time `json:"timestamp"`
}

func NewTx(from Account, to Account, value float32, nonce uint64, data string) *Tx {
	return &Tx{
		From:         from,
		To:           to,
		Value:        value,
		AccountNonce: nonce,
		Data:         data,
		Timestamp:    time.Now(),
	}
}

func (tx *Tx) String() string {
	return fmt.Sprintf("%+v", *tx)
}
