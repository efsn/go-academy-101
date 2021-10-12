package test

import (
	"io"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestCURL(t *testing.T) {
	r, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)

	dest := io.MultiWriter(os.Stdout, file)

	_, err1 := io.Copy(dest, r.Body)
	if err1 != nil {
		log.Fatalln(err1)
	}

	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}
}
