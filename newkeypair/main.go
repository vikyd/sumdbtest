package main

import (
	"crypto/rand"
	"fmt"
	"strings"

	"github.com/vikyd/sumdbtest/common/note"
)

// This example shows:
//   - how to generate Private Key, Public Key of ed25519 algorithm
//   - the hash in both key is the same: calculated only from name and public key
func main() {
	privateNote, publicNote, err := note.GenerateKey(rand.Reader, "example.com")
	if err != nil {
		panic(err)
	}

	// the private key and public key will change after each generation
	fmt.Println("golang private note: " + privateNote)
	fmt.Println("golang public  note: " + publicNote)

	// find the hash -----

	sep := "+"

	_, after := chop(privateNote, sep)
	privateKeyHash, after := chop(after, sep)
	privateKeyHash, after = chop(after, sep)
	privateKeyHash, _ = chop(after, sep)

	_, after = chop(publicNote, sep)
	publicKeyHash, _ := chop(after, sep)

	isOk := "√"
	if privateKeyHash != publicKeyHash {
		isOk = "×"
	}

	fmt.Println("hash in private key: " + privateKeyHash)
	fmt.Println("hash in public  key: " + publicKeyHash)
	fmt.Println("is equal ? " + isOk)
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
