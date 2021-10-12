package chain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Index        int64
	Timestamp    int64
	PreBlockHash string
	Hash         string
	Data         string
}

func (b *Block) Println() {
	fmt.Printf("Index: %d\nPreBlockHash: %s\nHash: %s\nData: %s\nTimestamp: %d", b.Index, b.PreBlockHash, b.Hash, b.Data, b.Timestamp)
}

func calcHash(b Block) string {
	blockData := strconv.FormatInt(b.Index, 10) + strconv.FormatInt(b.Timestamp, 10) + b.PreBlockHash + b.Data
	bytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(bytes[:])
}

func GenerateNext(pre Block, data string) Block {
	b := Block{
		Index:        pre.Index + 1,
		PreBlockHash: pre.Hash,
		Timestamp:    time.Now().Unix(),
		Data:         data,
	}
	b.Hash = calcHash(b)
	return b
}

func GenerateRoot() Block {
	pre := Block{
		Index: -1,
	}
	return GenerateNext(pre, "")
}
