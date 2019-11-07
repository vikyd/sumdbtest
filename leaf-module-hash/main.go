package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// this data is from: https://sum.golang.org/lookup/github.com/google/uuid@v1.1.1
	modNum := 842
	modHashNote := "github.com/google/uuid v1.1.1 h1:Gkbcsh/GbpXz7lPftLA3P6TYMwjCLYm83jiFQZF/3gY="
	gomodHashNote := "github.com/google/uuid v1.1.1/go.mod h1:TIyPZe4MgqvfeYDBFedMoGGpEw/LqOeaOT+nhxU+yHo="
	txt := modHashNote + "\n" + gomodHashNote + "\n"

	// this url can be captured when `go get github.com/google/uuid v1.1.1`
	tileURL := "https://sum.golang.org/tile/8/0/003"
	tileOffset := 3
	tileHeight := 8
	// power of 2, like 2^tileHeight
	tileWidth := 1 << tileHeight

	hashComposite := caclLeafHash(txt)

	fmt.Println("module hash: " + modHashNote)
	fmt.Println("module go.mod hash: " + gomodHashNote)
	fmt.Println("")
	fmt.Println("composite text: " + txt)
	fmt.Println("")
	fmt.Println("composite hash: " + hashComposite)

	hList := getTile(tileURL)
	for i, h := range hList {
		if hashComposite == h {
			fmt.Println("√ the tile contains the composite hash, the leaf hash is a composite of module hash and go.mod hash")
			fmt.Printf("hash offset in the tile: %d\n", i)

			leafNum := tileWidth*tileOffset + i
			if leafNum == modNum {
				fmt.Printf("√ leaf number matched, offset inside tile: %d\n", leafNum)
			} else {
				fmt.Printf("× leaf number not matched, offset inside tile: %d\n", leafNum)
			}
		}
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

	// convert to base64
	b64List := []string{}
	sha256Len := 32
	for i := 0; i < len(body); i = i + sha256Len {
		b64 := base64.StdEncoding.EncodeToString(body[i : i+sha256Len])
		b64List = append(b64List, b64)
	}

	return b64List
}

func caclLeafHash(txt string) string {
	var zeroPrefix = []byte{0x00}
	h := sha256.New()
	h.Write(zeroPrefix)
	h.Write([]byte(txt))
	sum := h.Sum(nil)

	b64 := base64.StdEncoding.EncodeToString(sum)
	return b64
}
