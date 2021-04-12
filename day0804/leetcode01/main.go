package main

import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "crypto/sha256"
    "errors"
    "golang.org/x/crypto/pbkdf2"
    "log"
)

func PKCS7Padding(cipherText []byte, blockSize int) []byte{
    paddingLength := blockSize - len(cipherText) % blockSize
    paddingCtx := bytes.Repeat([]byte{byte(paddingLength)}, paddingLength)
    return append(cipherText, paddingCtx...)
}
func PKCS7UnPadding(cipherText []byte) ([]byte, error) {
    length := len(cipherText)
    if length <= 0 {
        return nil ,errors.New("unpadding context is invalid")
    }
    unPaddingLength := int(cipherText[length - 1])
    if unPaddingLength > length {
        return nil, errors.New("unPadding length over origin data")
    }
    return cipherText[:(length - unPaddingLength)], nil
    
}

func AESCBCEncrypt(plainText []byte, key, iv []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    blockSize := block.BlockSize()
    paddingText := PKCS7Padding(plainText, blockSize)
    
    blockMode := cipher.NewCBCEncrypter(block, iv)
    encryptText := make([]byte, len(paddingText))
    blockMode.CryptBlocks(encryptText, paddingText)
    return encryptText, nil
}

func AesCBCDecrypt(encryptText []byte, key, iv []byte)([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    blockMode := cipher.NewCBCDecrypter(block, iv)
    decryptLen := len(encryptText)
    decryptText := make([]byte, decryptLen)
    blockMode.CryptBlocks(decryptText, encryptText)
    unPaddingText, err := PKCS7UnPadding(decryptText)
    return unPaddingText, err
}


func getPbkdf2Key(securityKey, salt []byte) []byte {
    key := pbkdf2.Key(securityKey, salt, 1, len(securityKey), sha256.New)
    return key
}

func main() {
    salt := []byte("ABCDEF0123456789")
    iv := []byte("0123456789ABCDEF")
    key := getPbkdf2Key([]byte("MYPasswordxfaewfaewfaewfaeafeawf"),salt)
    plainText := "Hello World"
    encrypt, err := AESCBCEncrypt([]byte(plainText), key, iv)
    if err != nil {
        log.Println(err)
        return
    }
    decrypt, err := AesCBCDecrypt(encrypt, key, iv)
    
    log.Printf("encrypt  value : %q", encrypt)
    log.Printf("plainText value %q: ", plainText)
    log.Printf("decrypt value %q: ", decrypt)
    
    if err == nil && string(decrypt) == plainText  {
        log.Println("TEST PASS")
    }else{
        log.Fatal("TEST FAILED")
    }
}
