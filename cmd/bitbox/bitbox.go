package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli"
	"golang.org/x/crypto/scrypt"
)

func main() {
	app := cli.NewApp()
	app.Name = "Bitbox"
	app.Usage = "Encrypt and Decrypt anything with AES using bytes. Encrypted result is base64 encoded"
	app.UsageText = "Bitbox [command] key text"
	app.HideHelp = true
	app.HideVersion = true
	app.ArgsUsage = ""
	app.Commands = []cli.Command{
		{
			Name:    "encrypt",
			Aliases: []string{"e"},
			Usage:   "Encrypt a string",
			Action: func(c *cli.Context) error {
				encText, err := encrypt(c.Args().Get(0), c.Args().Get(1))
				if err != nil {
					fmt.Println("error encrypting: ", err.Error())
				}
				fmt.Println(encText)
				return nil
			},
		},
		{
			Name:    "decrypt",
			Aliases: []string{"d"},
			Usage:   "Decrypt a string",
			Action: func(c *cli.Context) error {
				decText, err := decrypt(c.Args().Get(0), c.Args().Get(1))
				if err != nil {
					fmt.Println("error decrypting: ", err.Error())
				}
				fmt.Println(decText)
				return nil
			},
		},
	}
	app.Run(os.Args)
}

func encrypt(text, keyfile string) (string, error) {

	// check inputs
	if len(text) == 0 {
		return "", errors.New("text is empty")
	}
	// read contents of a key from file
	// read from file
	b, err := os.ReadFile(keyfile)
	if err != nil {
		return "", err
	}

	// transform text password into appropriate 32 byte key for AES
	// generate a new aes cipher using our 32 byte long key
	key, salt, err := DeriveKey(b, nil)
	if err != nil {
		return "", err
	}

	// generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	// if any error generating new GCM
	// handle them
	if err != nil {
		return "", err
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	encryptedTest := gcm.Seal(nonce, nonce, []byte(text), nil)
	// add salt at the end
	encryptedTest = append(encryptedTest, salt...)
	return base64.StdEncoding.EncodeToString(encryptedTest), nil
}

func decrypt(text, keyfile string) (string, error) {

	// check inputs
	if len(text) == 0 {
		return "", errors.New("text is empty")
	}
	b, err := os.ReadFile(keyfile)
	if err != nil {
		return "", err
	}

	// decode text from base64
	ciphertextWithSalt, err := base64.StdEncoding.DecodeString(text) // bytes
	if err != nil {
		return "", err
	}
	// check input text length
	if len(ciphertextWithSalt) < 32 {
		return "", errors.New("invalid input text")
	}

	// get salt from the end
	salt, ciphertext := ciphertextWithSalt[len(ciphertextWithSalt)-32:], ciphertextWithSalt[:len(ciphertextWithSalt)-32]

	key, _, err := DeriveKey(b, salt)
	if err != nil {
		return "", err
	}

	// generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	// get the nonce size
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println(err)
	}
	// extract our nonce from our encrypted text
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// here we decrypt our text using the Open function
	// Open decrypts and authenticates ciphertext, authenticates the
	// additional data and, if successful, appends the resulting plaintext
	// to dst, returning the updated slice. The nonce must be NonceSize()
	// bytes long and both it and the additional data must match the
	// value passed to Seal.
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func DeriveKey(password, salt []byte) ([]byte, []byte, error) {
	if salt == nil {
		salt = make([]byte, 32)
		if _, err := rand.Read(salt); err != nil {
			return nil, nil, err
		}
	}

	// minimum N is 16384
	// *32 will take about 2 sec
	n := 16384 * 32
	key, err := scrypt.Key(password, salt, n, 8, 1, 32)
	if err != nil {
		return nil, nil, err
	}

	return key, salt, nil
}
