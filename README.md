```
 _           _       _               
 | |         | |     | |              
 | |__  _   _| |_ ___| |__   _____  __
 | '_ \| | | | __/ _ \ '_ \ / _ \ \/ /
 | |_) | |_| | ||  __/ |_) | (_) >  < 
 |_.__/ \__, |\__\___|_.__/ \___/_/\_\
         __/ |                        
        |___/            

A command line tool written in Go for encrypting and decrypting text using a password. 

The output is base64 encoded.
```

## Why use a password instead of a key?

Bytebox uses the [scrypt](https://godoc.org/golang.org/x/crypto/scrypt) key derivation function to transform the user-provided password into a 32-byte key suitable for use with AES.

## How it works

Bytebox uses [Advanced Encryption Standard (AES)](https://en.wikipedia.org/wiki/Advanced_Encryption_Standard) in [Galois/Counter Mode (GCM)](https://en.wikipedia.org/wiki/Galois/Counter_Mode) for encryption and decryption. GCM is a mode of operation for symmetric key cryptographic block ciphers that provides authenticity and confidentiality.

## Usage

```bash
NAME:
   Bytebox - Encrypt and Decrypt text with AES using password. Encrypted result is base64 encoded

USAGE:
   Bytebox [command] text password

COMMANDS:
   encrypt, e  Encrypt a string
   decrypt, d  Decrypt a string
```

## Examples

Encrypt a string "secret_string" with the password 12345678:
```
Bytebox e secret_string 123456789
```
Output:
```
WOx69r3Y1a0qNqZGvwG44eJ5FaU+xyHjPdFDt8klThlrZ5m1o7rG3tMztr8KibXMfN4UW3m/dfLyCmimnyQPX3f5
```

Decrypt a base64 string with the password 12345678:

```
Bytebox d WOx69r3Y1a0qNqZGvwG44eJ5FaU+xyHjPdFDt8klThlrZ5m1o7rG3tMztr8KibXMfN4UW3m/dfLyCmimnyQPX3f5 12345678
```
Output:
```
secret_string
```


## Todo 🚀
- [x] Encode a string with a string
- [ ] Encode a string with a file
- [ ] Encode file with a file
- [ ]  Allow to choose output format
- [ ]  Allow to choose delay for key derivation function


## Contributions

Feel free to open pull requests or issues to improve Bytebox.


## License

Bytebox is open-source software licensed under the MIT License.