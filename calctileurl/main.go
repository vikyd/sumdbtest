package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Lookup struct {
	ModNum    int
	LeafCount int
}

var tileHeight = 8

// 2^tileHeight
var tileWidth = 1 << tileHeight

func main() {
	lookURL := "https://sum.golang.org/lookup/github.com/google/uuid@v1.1.1"

	lookup := getLookup(lookURL)
	fmt.Println(lookup)

	calcTileContainsModHash(lookup.ModNum)
	calcTileLeafEnd(lookup.LeafCount)

}

func getLookup(url string) Lookup {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	bodyStr := string(body)

	fmt.Println("-------- lookup response ↓ --------")
	fmt.Println(bodyStr)
	fmt.Println("-------- lookup response ↑ --------")

	lines := strings.Split(bodyStr, "\n")

	modNum, _ := strconv.Atoi(lines[0])
	leafCount, _ := strconv.Atoi(lines[5])
	l := Lookup{
		ModNum:    modNum,
		LeafCount: leafCount,
	}
	return l
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

func calcTileContainsModHash(modNum int) {
	leafLevel := 0

	// example: 5 / 2 = 2
	tileCount := modNum / tileWidth
	remain := modNum % tileWidth

	// index begin from 0
	tileIndex := tileCount - 1
	if remain > 0 {
		tileIndex = tileIndex + 1
	}

	tileURL := fmt.Sprintf("http://sum.golang.org/tile/%d/%d/%03d", tileHeight, leafLevel, tileIndex)
	fmt.Println(tileURL)
}

func calcTileLeafEnd(leafCount int) {
	leafLevel := 0

	// example: 5 / 2 = 2
	tileCount := leafCount / tileWidth
	remain := leafCount % tileWidth

	// index begin from 0
	tileIndex := tileCount - 1
	if remain > 0 {
		tileIndex = tileIndex + 1
	}

	fmt.Println(tileCount)

	tileURL := fmt.Sprintf("http://sum.golang.org/tile/%d/%d/%03d", tileHeight, leafLevel, tileIndex)
	fmt.Println(tileURL)
}
