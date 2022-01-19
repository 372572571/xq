package sbab

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"xq/internal"

	"golang.org/x/crypto/bcrypt"
)

var defkey = string(internal.Util.MustTake(hex.DecodeString("6368616e67652074686973e070e17323")).([]byte))

type Te struct {
	opts   options
	sha    string
	bcrypt string
	aes    string
	base   string
}

func Check(orig string, cipher string, mask string) bool {
	var base = internal.Util.MustTake(base64.StdEncoding.DecodeString(cipher)).([]byte)
	
	if mask == "" {
		mask = defkey
	}

	var aes  = aes_decryption(string(base), mask)
	    orig = sha_encryption(orig)

	   ok := bcrypt.CompareHashAndPassword([]byte(aes), []byte(orig))
	if ok != nil {
		return false
	}

	return true
}

func New(optfuncs ...OptFunc) *Te {
	var te   = &Te{}
	var opts = options{}
	for _,    f := range optfuncs {
		f(&opts)
	}
	te.opts = opts
	return te
}

func (t *Te) Encryption() string {
	t.sha    = sha_encryption(t.opts.source)                     // sha512
	t.bcrypt = bcrypt_encryption(t.sha)                          // bcrypt 强度 10
	t.aes    = aes_encryption(t.bcrypt, t.opts.mask)             // aes
	t.base   = base64.StdEncoding.EncodeToString([]byte(t.aes))  // base
	return t.base
}

func sha_encryption(value string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(value)))
}

func bcrypt_encryption(value string) string {
	var h = internal.Util.MustTake(bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)).([]byte)
	return string(h)
}

// aes cbc
func aes_encryption(value, key string) string {
	if key == "" {
		key = defkey
	}
	var block      = internal.Util.MustTake(aes.NewCipher([]byte(defkey))).(cipher.Block)
	var block_size = block.BlockSize()
	var p_value    = pkcs7padding([]byte(value), block.BlockSize())
	var block_mode = cipher.NewCBCEncrypter(block, []byte(key)[:block_size])

	encryption := make([]byte, len(p_value))
	block_mode.CryptBlocks(encryption, []byte(p_value))
	return string(encryption)
}

func aes_decryption(value, key string) string {
	block      := internal.Util.MustTake(aes.NewCipher([]byte(key))).(cipher.Block)
	blockSize  := block.BlockSize()
	blockMode  := cipher.NewCBCDecrypter(block, []byte(key)[:blockSize])
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
	length    := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
