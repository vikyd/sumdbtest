package main

import (
	"crypto/rand"
	"fmt"

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

	// Golang text like: https://sum.golang.org/lookup/github.com/google/uuid@v1.1.1
	text := "This is a message example\n" +
		"\n" +
		publicNote + "\n"

	// ---- ↓ sign ↓ ----
	signer, err := note.NewSigner(privateNote)
	if err != nil {
		panic(err)
	}

	msg, err := note.Sign(&note.Note{Text: text}, signer)
	if err != nil {
		panic(err)
	}

	// ---- ↓ verify ↓ ----
	verifier, err := note.NewVerifier(publicNote)
	if err != nil {
		panic(err)
	}
	verifiers := note.VerifierList(verifier)

	// do verify
	n, err := note.Open(msg, verifiers)
	if err != nil {
		panic(err)
	}
	fmt.Println("signer name: " + n.Sigs[0].Name)
	fmt.Println("text to be signed: " + n.Text)
	fmt.Println("√ verify success")

}
