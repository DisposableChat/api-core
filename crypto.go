package core

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

var LocalHashPassword string

func Encrypt(password, message string) (string, error) {
    salt := make([]byte, 8)
    _, err := io.ReadFull(rand.Reader, salt)
    if err != nil {
        return "", err
    }

    // Derive a key from the password using PBKDF2
    key := pbkdf2.Key([]byte(password), salt, 4096, 32, sha256.New)

    // Use the key to encrypt the message
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }
    nonce := make([]byte, gcm.NonceSize())
    _, err = io.ReadFull(rand.Reader, nonce)
    if err != nil {
        return "", err
    }
    encrypted := gcm.Seal(nil, nonce, []byte(message), nil)

    buf := bytes.NewBuffer(nil)
    buf.Write(salt)
    buf.Write(nonce)
    buf.Write(encrypted)
    return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func Decrypt(password, encrypted string) (string, error) {
    decoded, err := base64.StdEncoding.DecodeString(encrypted)
    if err != nil {
        return "", err
    }
    salt, decoded := decoded[:8], decoded[8:]
    nonce, decoded := decoded[:12], decoded[12:]

    key := pbkdf2.Key([]byte(password), salt, 4096, 32, sha256.New)

    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }
    decrypted, err := gcm.Open(nil, nonce, decoded, nil)
    if err != nil {
        return "", err
    }
    return string(decrypted), nil
}

func GenerateLocalHashPassword() error {
	var err error
	LocalHashPassword, err = RandomToken(os.Getpid())
	if err != nil {
		return err
	}

	return err
}