package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/vikyd/sumdbtest/common/note"
)

var sep = "+"

// This example shows:
//   - create a new sign for a message by private key
//   - show the length of private key, public key, sign
func main() {
	// new key pair
	privateNote, publicNote, err := note.GenerateKey(rand.Reader, "example.com")
	if err != nil {
		panic(err)
	}

	// the private key and public key will change after each generation
	fmt.Println("golang private note: " + privateNote)
	fmt.Println("golang public  note: " + publicNote)

	checkPrivateKeyLen(privateNote)
	checkPublicKeyLen(publicNote)

}

func checkPrivateKeyLen(privateNote string) {
	_, after := chop(privateNote, sep)
	_, after = chop(after, sep)
	_, after = chop(after, sep)
	_, algAndKey := chop(after, sep)

	// ---- ↓ split algorithm and key ↓ ----
	bs, _ := base64.StdEncoding.DecodeString(algAndKey)
	algB := bs[0]
	keyBs := bs[1:]

	// ---- ↓ check sign algorithm ↓ ----
	fmt.Printf("private note algorithm number: %b\n", algB)

	// ---- ↓ check key length ↓ ----
	keyLen := len(keyBs)
	fmt.Printf("private key length: %d bits\n", keyLen)
}

func checkPublicKeyLen(publicNote string) {
	_, after := chop(publicNote, sep)
	_, algAndKey := chop(after, sep)

	// ---- ↓ split algorithm and key ↓ ----
	bs, _ := base64.StdEncoding.DecodeString(algAndKey)
	algB := bs[0]
	keyBs := bs[1:]

	// ---- ↓ check sign algorithm ↓ ----
	fmt.Printf("public note algorithm number: %b\n", algB)

	// ---- ↓ check key length ↓ ----
	keyLen := len(keyBs)
	fmt.Printf("public key length: %d bits\n", keyLen)
}

// chop chops s at the first instance of sep, if any,
// and returns the text before and after sep.
// If sep is not present, chop returns before is s and after is empty.
func chop(s, sep string) (before, after string) {
	i := strings.Index(s, sep)
	if i < 0 {
		return s, ""
	}
	return s[:i], s[i+len(sep):]
}
