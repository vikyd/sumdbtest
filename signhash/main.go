package main

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
)

// This example shows:
//   - parse hash from a sign
//   - the hash is the first 32bit of the sign
//   - ref: https://github.com/golang/exp/blob/master/sumdb/internal/note/note.go#L576
//   - the hash is the same as public key hash
//   - for Golang the hash in hex is always: 033de0ae
func main() {
	// this base64 string can be found from sumdb `/lookup` or `/latest` :
	//   - may be the sign is not the same, but this hash is not changed
	//   - https://sum.golang.org/lookup/github.com/google/uuid@v1.1.1
	//   - https://sum.golang.org/latest
	b64 := "Az3grrJsLRs6sNa2gQWy6G6jb/FLI7opFZErrJT1PWmmP4iUdRxoJhMgfmSkirJgj3zj7n3N61yL16+9521wNu12Sgo="
	sig, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		panic(err)
	}
	checkHash(sig[0:4])
	checkSignLen(sig[4:])

}

// check the length of the real sign
func checkSignLen(bs []byte) {
	msg := fmt.Sprintf("length of sign: %d bits", len(bs))
	fmt.Println(msg)
}

// check what is the real hash
func checkHash(hashBs []byte) {
	// from: https://github.com/golang/go/blob/master/src/cmd/go/internal/modfetch/key.go#L8
	expectdHash := "033de0ae"

	hash := binary.BigEndian.Uint32(hashBs)

	hex := fmt.Sprintf("%08x", hash)
	fmt.Println("hash: " + hex)

	isOk := "√"
	if hex != expectdHash {
		isOk = "×"
	}

	fmt.Println("is equals: " + isOk)

}
