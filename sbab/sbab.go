package sbab

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Keyer(key string) string {
	str, err := hex.DecodeString(key)

	if err != nil {
		panic(err)
	}

	return string(str)
}

func defkey() string {
	return Keyer("6368616e67652074686973e070e17323")
}

type Te struct {
	opts   options
	sha    string
	bcrypt string
	aes    string
	base   string
}

func Check(orig string, cipher string, mask string) bool {
	var base, err = base64.StdEncoding.DecodeString(cipher)
	if err != nil {
		panic(err)
	}

	if mask == "" {
		mask = defkey()
	}

	var aes = aes_decryption(string(base), mask)
	orig = sha_encryption(orig)

	ok := bcrypt.CompareHashAndPassword([]byte(aes), []byte(orig))
	if ok != nil {
		return false
	}

	return true
}

func New(optfuncs ...OptFunc) *Te {
	var te = &Te{}
	var opts = options{}

	for _, f := range optfuncs {
		f(&opts)
	}

	te.opts = opts
	return te
}

func (t *Te) Encryption() string {
	t.sha = sha_encryption(t.opts.source)                     // sha512
	t.bcrypt = bcrypt_encryption(t.sha)                       // bcrypt 强度 10
	t.aes = aes_encryption(t.bcrypt, t.opts.mask)             // aes
	t.base = base64.StdEncoding.EncodeToString([]byte(t.aes)) // base
	return t.base
}

func sha_encryption(value string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(value)))
}

func bcrypt_encryption(value string) string {
	h, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	
	if err != nil {
		panic(err)
	}

	return string(h)
}

// aes cbc
func aes_encryption(value, key string) string {
	if key == "" {
		key = defkey()
	}

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		panic(err)
	}

	var block_size = block.BlockSize()
	var p_value = pkcs7padding([]byte(value), block.BlockSize())
	var block_mode = cipher.NewCBCEncrypter(block, []byte(key)[:block_size])

	encryption := make([]byte, len(p_value))
	block_mode.CryptBlocks(encryption, []byte(p_value))
	return string(encryption)
}

func aes_decryption(value, key string) string {
	block, err := aes.NewCipher([]byte(key)) // .(cipher.Block)
	if err != nil {
		panic(err)
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, []byte(key)[:blockSize])
	decryption := make([]byte, len(value))
	blockMode.CryptBlocks(decryption, []byte(value))
	decryption = pkcs7unpadding(decryption)
	return string(decryption)
}

func pkcs7padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs7unpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
