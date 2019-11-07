package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Lookup struct {
	ModuleNum int
	LeafCount int
}

type Tile struct {
	H int   // height of tile (1 ≤ H ≤ 30)
	L int   // level in tiling (-1 ≤ L ≤ 63)
	N int64 // number within level (0 ≤ N, unbounded)
	W int   // width of tile (1 ≤ W ≤ 2**H; 2**H is complete tile)
}

var SumdbURL = "https://sum.golang.org"

var TileHeight = 8

// 2^tileHeight
var TileWidth = 1 << TileHeight

var LeafLevel = 0

func main() {
	lookURL := "https://sum.golang.org/lookup/github.com/google/uuid@v1.1.1"

	lookup := getLookup(lookURL)
	fmt.Printf("module number: %d\n", lookup.ModuleNum)
	fmt.Printf("leaf count   : %d\n", lookup.LeafCount)

	calcTileContainsModHash(lookup.ModuleNum, lookup.LeafCount)
	calcTileLeafEnd(lookup.LeafCount)

}

// receive a response from lookup endpoint,
// and return the module number in leafs of the tree
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
		ModuleNum: modNum,
		LeafCount: leafCount,
	}
	return l
}

func calcTileContainsModHash(modNum int, leafCount int) {
	// example: 5 / 2 = 2
	tileCount := modNum / TileWidth
	remain := modNum % TileWidth

	// index begin from 0
	tileIndex := tileCount - 1
	tileWidth := TileWidth
	if remain > 0 {
		tileIndex = tileIndex + 1
		if tileIndex*TileWidth > leafCount {
			tileWidth = remain
		}
	}

	tile := Tile{
		H: TileHeight,
		L: LeafLevel,
		N: int64(tileIndex),
		W: tileWidth,
	}
	path := tile.Path()
	fmt.Println("path of tile contains the module: " + SumdbURL + path)
}

func calcTileLeafEnd(leafCount int) {
	// example: 5 / 2 = 2
	tileCount := leafCount / TileWidth
	remain := leafCount % TileWidth

	// index begin from 0
	tileIndex := tileCount - 1
	tileWidth := TileWidth
	if remain > 0 {
		tileIndex = tileIndex + 1
		tileWidth = remain
	}

	tile := Tile{
		H: TileHeight,
		L: LeafLevel,
		N: int64(tileIndex),
		W: tileWidth,
	}
	path := tile.Path()
	fmt.Println("path of tile at the end of leafs: " + SumdbURL + path)
}

// To limit the size of any particular directory listing,
// we encode the (possibly very large) number N
// by encoding three digits at a time.
// For example, 123456789 encodes as x123/x456/789.
// Each directory has at most 1000 each xNNN, NNN, and NNN.p children,
// so there are at most 3000 entries in any one directory.
const pathBase = 1000

// Path returns a tile coordinate path describing t.
func (t Tile) Path() string {
	n := t.N
	nStr := fmt.Sprintf("%03d", n%pathBase)
	for n >= pathBase {
		n /= pathBase
		nStr = fmt.Sprintf("x%03d/%s", n%pathBase, nStr)
	}
	pStr := ""
	if t.W != 1<<uint(t.H) {
		pStr = fmt.Sprintf(".p/%d", t.W)
	}
	var L string
	if t.L == -1 {
		L = "data"
	} else {
		L = fmt.Sprintf("%d", t.L)
	}
	return fmt.Sprintf("tile/%d/%s/%s%s", t.H, L, nStr, pStr)
}
