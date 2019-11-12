package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

// This example shows:
//   - how to parse the response of a tile
//   - the length of a full size tile
func main() {
	// this url can be captured when `go get github.com/google/uuid v1.1.1`
	tileURL := "https://sum.golang.org/tile/8/0/003"

	resp, err := http.Get(tileURL)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// convert to base64
	b64List := []string{}
	sha256Len := 32
	for i := 0; i < len(body); i = i + sha256Len {
		b64 := base64.StdEncoding.EncodeToString(body[i : i+sha256Len])
		b64List = append(b64List, b64)
	}

	fmt.Println("SHA-256 hashes in response: ")
	fmt.Println()
	for i, hash := range b64List {
		fmt.Printf("[%03d]%s\n", i+1, hash)
	}

	fmt.Println()
	fmt.Printf("url: %s\n", tileURL)
	fmt.Printf("response length: %d bytes\n", len(body))
}
