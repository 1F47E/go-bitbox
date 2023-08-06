package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/1F47E/go-bitbox/pkg/myaes"
	"github.com/1F47E/go-bitbox/utils"
	"github.com/urfave/cli"
)

// auth
type Symmetric interface {
	Encrypt([]byte, string) ([]byte, error)
	Decrypt([]byte, string) ([]byte, error)
}

func main() {
	app := cli.NewApp()
	app.Name = "Bitbox"
	app.Usage = "Encrypt and Decrypt anything with AES using password"
	app.UsageText = "bitbox [command] key text"
	app.HideHelp = true
	app.HideVersion = true
	app.ArgsUsage = ""
	app.Commands = []cli.Command{
		{
			Name:    "encrypt",
			Aliases: []string{"e"},
			Usage:   "Encrypt a string",
			Action:  encrypt,
		},
		{
			Name:    "decrypt",
			Aliases: []string{"d"},
			Usage:   "Decrypt a string",
			Action:  decrypt,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		utils.PrintError(err.Error())
		os.Exit(1)
	}
}

func encrypt(c *cli.Context) error {
	var crypter = Symmetric(&myaes.AEScrypter{})
	data := []byte(c.Args().Get(0))
	pass := c.Args().Get(1)
	encText, err := crypter.Encrypt(data, pass)
	if err != nil {
		return fmt.Errorf("error encrypting: %s", err.Error())
	}
	b64 := base64.StdEncoding.EncodeToString(encText)
	utils.PrintSuccess(b64)
	return nil
}

func decrypt(c *cli.Context) error {
	var crypter = Symmetric(&myaes.AEScrypter{})
	text := []byte(c.Args().Get(0))
	pass := c.Args().Get(1)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return fmt.Errorf("error decoding: %s", err.Error())
	}
	decText, err := crypter.Decrypt(data, pass)
	if err != nil {
		return fmt.Errorf("error decrypting: %s", err.Error())
	}
	utils.PrintSuccess(string(decText))
	return nil
}
