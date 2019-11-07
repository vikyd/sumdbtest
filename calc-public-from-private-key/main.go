package main

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/vikyd/sumdbtest/common/note"
)

var sep = "+"

// This example shows:
func main() {
	// new key pair
	privateNote, publicNote, err := note.GenerateKey(rand.Reader, "example.com")
	if err != nil {
		panic(err)
	}

	// the private key and public key will change after each generation
	fmt.Println("golang private note: " + privateNote)
	fmt.Println("golang public  note: " + publicNote)

	priKey := parsePrivateKeyBytes(privateNote)
	pubKey := parsePublicKeyBytes(publicNote)

	newKey := ed25519.NewKeyFromSeed(priKey)

	newKeyPrefix := newKey[:32]
	newKeySubfix := newKey[32:]
	if bytes.Equal(newKeyPrefix, priKey) {
		fmt.Println("√ private key is equal from calc prefix")
	} else {
		fmt.Println("× private key is not equal from calc prefix")
	}

	if bytes.Equal(newKeySubfix, pubKey) {
		fmt.Println("√ private key can create publick key")
	} else {
		fmt.Println("× private key can not create public key")
	}

}

func parsePrivateKeyBytes(privateNote string) []byte {
	_, after := chop(privateNote, sep)
	_, after = chop(after, sep)
	_, after = chop(after, sep)
	_, algAndKey := chop(after, sep)

	// ---- ↓ split algorithm and key ↓ ----
	bs, _ := base64.StdEncoding.DecodeString(algAndKey)
	keyBs := bs[1:]
	return keyBs
}

func parsePublicKeyBytes(publicNote string) []byte {
	_, after := chop(publicNote, sep)
	_, algAndKey := chop(after, sep)

	// ---- ↓ split algorithm and key ↓ ----
	bs, _ := base64.StdEncoding.DecodeString(algAndKey)
	keyBs := bs[1:]
	return keyBs
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
