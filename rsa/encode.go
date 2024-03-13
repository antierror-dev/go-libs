package rsa

import (
    "crypto/rand"
    RSA "crypto/rsa"
    "crypto/x509"
    "encoding/base64"
    "encoding/pem"
    "fmt"
)

func Encode(publicKeyStr, plaintext string) (encodedText string, err error) {
    publicKeyBlock, _ := pem.Decode([]byte(publicKeyStr))
    if publicKeyBlock == nil {
        return "", fmt.Errorf("public key decode error")
    }
    publicKeyObj, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
    if err != nil {
        return "", fmt.Errorf("public key parsing error: %v", err)
    }
    rsaPub, ok := publicKeyObj.(*RSA.PublicKey)
    if !ok {
        return "", fmt.Errorf("public key casting error")
    }

    ciphertext, err := RSA.EncryptPKCS1v15(rand.Reader, rsaPub, []byte(plaintext))
    if err != nil {
        return "", err
    }

    encodedText = base64.StdEncoding.EncodeToString(ciphertext)
    return encodedText, nil
}
