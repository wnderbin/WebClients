package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <url>\n", filepath.Base(os.Args[0]))
		return
	}
	url := os.Args[1]
	data, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(os.Stdout, data.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	data.Body.Close()
}
