package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/jaiiali/go-simple-blockchain/db"
)

func run() error {
	r := makeMuxRouter()
	httpAddr := "8080"
	log.Println("Listening on ", httpAddr)
	s := &http.Server{
		Addr:    ":" + httpAddr,
		Handler: r,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods(http.MethodGet)
	muxRouter.HandleFunc("/", handleWriteBlock).Methods(http.MethodPost)

	return muxRouter
}

func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(storage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	io.WriteString(w, string(bytes))
}

func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var txs []db.Tx

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&txs); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	spew.Dump(txs)

	var inputs []*db.Tx
	for i := range txs {
		v := db.NewTx(txs[i].From, txs[i].To, txs[i].Value, txs[i].AccountNonce, txs[i].Data)
		inputs = append(inputs, v)
	}

	b, err := db.NewBlock(storage.Blocks[storage.LastHeight-1], inputs)
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, err)
		return
	}

	err = storage.AddBlock(b)
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, err)
		return
	}

	respondWithJSON(w, r, http.StatusCreated, b)
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}

	w.WriteHeader(code)
	w.Write(response)
}
