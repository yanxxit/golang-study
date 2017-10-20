package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"os"
)

func encrypt(aeskey string, filename string) {
	plaintext, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}

	// Byte array of the string
	key := []byte(aeskey)

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// create a new file for saving the encrypted data.
	f, err := os.Create(filename + ".aes")
	if err != nil {
		panic(err.Error())
	}
	_, err = io.Copy(f, bytes.NewReader(ciphertext))
	if err != nil {
		panic(err.Error())
	}
}

func decrypt(aesKey string, inputFile string) {

	ciphertext, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err.Error())
	}

	// Key
	key := []byte(aesKey)

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Before even testing the decryption,
	// if the text is too small, then it is incorrect
	if len(ciphertext) < aes.BlockSize {
		panic("Text is too short")
	}

	// Get the 16 byte IV
	iv := ciphertext[:aes.BlockSize]

	// Remove the IV from the ciphertext
	ciphertext = ciphertext[aes.BlockSize:]

	// Return a decrypted stream
	stream := cipher.NewCFBDecrypter(block, iv)
	// Decrypt bytes from ciphertext
	stream.XORKeyStream(ciphertext, ciphertext)
	// create a new file for saving the encrypted data.
	f, err := os.Create(inputFile + ".ts")
	if err != nil {
		panic(err.Error())
	}
	_, err = io.Copy(f, bytes.NewReader(ciphertext))
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	key := "0123456789123456"
	encrypt(key, "1.ts")
	decrypt(key, "1.ts.aes")
}
