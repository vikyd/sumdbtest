package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"fmt"
)

// This example shows that how to calculate the hash from sumdb url, algorithm number, public key
func main() {
	// Example: sum.golang.org+033de0ae+Ac4zctda0e5eza+HJyk9SxEdh+s3Ux18htTTAD8OuAn8
	// From: https://github.com/golang/go/blob/master/src/cmd/go/internal/modfetch/key.go#L8
	name := "sum.golang.org"
	expectedHash := "033de0ae"
	algAndPublicKey := "Ac4zctda0e5eza+HJyk9SxEdh+s3Ux18htTTAD8OuAn8"

	// ---- ↓ check hash ↓ ----
	hash := keyHashStr(name, algAndPublicKey)

	fmt.Println("calculated hash: " + hash)
	fmt.Println("expected   hash: " + expectedHash)
	isEqual := "√"
	if hash != expectedHash {
		isEqual = "×"
	}
	fmt.Println("is equal: " + isEqual)

	// ---- ↓ check public key length ↓ ----
	// keyBytes, _ := base64.StdEncoding.DecodeString(key)

}

// keyHashStr computes hash and return string
func keyHashStr(name string, key string) string {
	keyBytes, _ := base64.StdEncoding.DecodeString(key)
	uintVal := keyHash(name, keyBytes)
	return fmt.Sprintf("%08x", uintVal)
}

// keyHash computes the key hash for the given server name and encoded public key.
func keyHash(name string, key []byte) uint32 {
	h := sha256.New()
	h.Write([]byte(name))
	h.Write([]byte("\n"))
	h.Write(key)
	sum := h.Sum(nil)
	return binary.BigEndian.Uint32(sum)
}
