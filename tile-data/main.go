package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// this url can be captured when `go get github.com/google/uuid v1.1.1`
	tileURL := "https://sum.golang.org/tile/8/0/003"

	hList := getTile(tileURL)
	fmt.Println("SHA-256 hashes in response: ")
	fmt.Println()
	for i, hash := range hList {
		fmt.Printf("[%03d]%s\n", i+1, hash)
	}
}

func getTile(url string) []string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("url: %s\n", url)
	fmt.Printf("response size: %d bits\n", len(body))

	// convert to base64
	b64List := []string{}
	sha256Len := 32
	for i := 0; i < len(body); i = i + sha256Len {
		b64 := base64.StdEncoding.EncodeToString(body[i : i+sha256Len])
		b64List = append(b64List, b64)
	}

	return b64List
}
