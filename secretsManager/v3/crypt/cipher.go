package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func entryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, err := newCipherBlock(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewCFBEncrypter(block, iv), nil

}
func decryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, err := newCipherBlock(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewCFBDecrypter(block, iv), nil

}

func Encrypt(plaintext, key string) (string, error) {
	// copde below copied straight from https://pkg.go.dev/crypto/cipher#NewCFBEncrypter
	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream, err := entryptStream(key, iv)
	if err != nil {
		return "", err
	}

	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))
	return fmt.Sprintf("%x", ciphertext), nil
}

func EncryptWriter(key string, w io.Writer) (*cipher.StreamWriter, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream, err := entryptStream(key, iv)
	if err != nil {
		return nil, err
	}
	n, err := w.Write(iv)
	if n != len(iv) || err != nil {
		return nil, fmt.Errorf("unable to write iv to writer (%v)", err)
	}
	return &cipher.StreamWriter{S: stream, W: w}, nil
}

func Decrypt(cipherHex, key string) (string, error) {
	ciphertext, err := hex.DecodeString(cipherHex)
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream, err := decryptStream(key, iv)
	if err != nil {
		return "", err
	}
	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}

func DecryptReader(key string, r io.Reader) (*cipher.StreamReader, error) {
	iv := make([]byte, aes.BlockSize)
	n, err := r.Read(iv)
	if n != len(iv) || err != nil {
		return nil, fmt.Errorf("unable to read iv to reader (%v)", err)
	}
	stream, err := decryptStream(key, iv)
	if err != nil {
		return nil, err
	}
	return &cipher.StreamReader{S: stream, R: r}, nil
}

func newCipherBlock(key string) (cipher.Block, error) {
	// NewCipher creates and returns a new cipher.Block. The key argument should be the AES key,
	// either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
	h := md5.New()
	io.WriteString(h, key)

	cipherKey := h.Sum(nil)
	return aes.NewCipher(cipherKey)
}
