package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"
	"os"
)

const secretKey = "supersecretkeyyy"
const filePath = "./filewithencryptedtext"

func decrypt(ciphered []byte) string {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		log.Fatal(err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalln(err)
	}

	nonceSize := aesgcm.NonceSize()
	nonce, cipheredText := ciphered[0:nonceSize], ciphered[nonceSize:]

	originalText, err := aesgcm.Open(nil, nonce, cipheredText, nil)
	if err != nil {
		log.Fatalln(err)
	}
	return string(originalText)
}

func main() {
	// open s3 object and read as bytes
	encryptedText, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(decrypt(encryptedText))
}
