package tools

import (
	"log"
	"os"
)

func errHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Load(file string) {
	dataset, err := os.Open(file)
	errHandler(err)
	defer dataset.Close()
}
