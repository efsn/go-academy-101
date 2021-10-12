package rpc

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/efsn/go-workshop-101/internal/chain"
)

var blockChain *chain.BlockChain

func Start(b *chain.BlockChain) {
	blockChain = b
	http.HandleFunc("/chain/read", ReadHandler)
	http.HandleFunc("/chain/write", WriteHandler)
	_ = http.ListenAndServe(":8000", nil)
}

// ReadHandler modify with point
func ReadHandler(w http.ResponseWriter, _ *http.Request) {
	bytes, err := json.Marshal(blockChain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = io.WriteString(w, string(bytes))
}

func WriteHandler(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("data")
	blockChain.SendData(data)
	ReadHandler(w, r)
}
