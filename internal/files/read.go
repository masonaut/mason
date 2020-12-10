package files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ReadFile(path string) *bufio.Reader {
	buffer, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewReader(buffer)
}

func ReadFileToByte(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
	}

	return data
}
