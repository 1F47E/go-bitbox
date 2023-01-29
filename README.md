# AESpass 🔒
```
 _______  _______  _______  _______  _______  _______  _______ 
|   _   ||       ||       ||       ||   _   ||       ||       |
|  |_|  ||    ___||  _____||    _  ||  |_|  ||  _____||  _____|
|       ||   |___ | |_____ |   |_| ||       || |_____ | |_____ 
|       ||    ___||_____  ||    ___||       ||_____  ||_____  |
|   _   ||   |___  _____| ||   |    |   _   | _____| | _____| |
|__| |__||_______||_______||___|    |__| |__||_______||_______|


A command line tool for encrypting and decrypting text using a password. 

The output is base64 encoded.
```

## Why use a password instead of a key?

AESpass uses the [scrypt](https://godoc.org/golang.org/x/crypto/scrypt) key derivation function to transform the user-provided password into a 32-byte key suitable for use with AES.

## How it works

AESpass uses [Advanced Encryption Standard (AES)](https://en.wikipedia.org/wiki/Advanced_Encryption_Standard) in [Galois/Counter Mode (GCM)](https://en.wikipedia.org/wiki/Galois/Counter_Mode) for encryption and decryption. GCM is a mode of operation for symmetric key cryptographic block ciphers that provides authenticity and confidentiality.

## Usage

```bash
NAME:
   AESpass - Encrypt and Decrypt text with AES using password. Encrypted result is base64 encoded

USAGE:
   aespass [command] text password

COMMANDS:
   encrypt, e  Encrypt a string
   decrypt, d  Decrypt a string
```

## Examples

Encrypt a string with the password 12345678:
```
aespass e secret 123456789
```
Output:
```
WOx69r3Y1a0qNqZGvwG44eJ5FaU+xyHjPdFDt8klThlrZ5m1o7rG3tMztr8KibXMfN4UW3m/dfLyCmimnyQPX3f5
```

Decrypt a string with the password 12345678:

```
aespass d WOx69r3Y1a0qNqZGvwG44eJ5FaU+xyHjPdFDt8klThlrZ5m1o7rG3tMztr8KibXMfN4UW3m/dfLyCmimnyQPX3f5 12345678
```
Output:
```
secret
```


## Todo 🚀
- [x] Encode string with a string
- [ ] Encode file
- [ ] Use any file as a password
- [ ]  Allow to choose output format
- [ ]  Allow to choose delay for key derivation function


## Contributions

Feel free to open pull requests or issues to improve AESpass.


## License

AESpass is open-source software licensed under the MIT License.