package hash_test

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"testing"
)

func calcHash(s string) string {
	bytes := sha256.Sum256([]byte(s))
	str := hex.EncodeToString(bytes[:])
	log.Printf("%s, %s", s, str)
	return str
}

func TestHash(t *testing.T) {
	str := "Oho X"
	hashStr := calcHash(str)
	log.Printf("%s", hashStr)
}
