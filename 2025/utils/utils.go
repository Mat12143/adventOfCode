package utils

import (
	"log"
	"os"
)

func ReadFile(filename string) string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	data := string(bytes)
	return data
}
