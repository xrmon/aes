# AES Implementation in Go

An AES implementation I wrote in Go, used to create the CryptoFunHouse challenge on TryHackMe (which you can try out [here](https://tryhackme.com/room/cryptofunhouse)).

I wrote this to learn about the internals of the AES cipher, also known as Rijndael, and Go itself. It's here in case it's useful in CTF challenges, but is not intended for real-world crypto. Use a tried and tested implementation if you need that.

Import by placing the following at the top of your Go program:

```go
import "github.com/xrmon/aes"
```

Then run the command:

```
go get github.com/xrmon/aes
```

As always, any function starting in a capital letter is exported. For example, the function:

```go
func BlockEncrypt128(input [16]byte, key [16]byte) [16]byte
```

encrypts a 16-byte block with a 16-byte master key, using AES 128. The 16-byte output is returned.

There is currently no support for padding or modes of operation - this package only includes the cipher itself, which operates on blocks of exactly 16 bytes.

### File Structure

*aes.go*: Main package functions

*aes_test.go*: Test suite for functions in aes.go

*galois.go*: Support for multiplication in the field GF(2^8). This is not an efficient implmentation but provides a good insight into the underlying mathematics.

*galois_test.go*: Test suite for the galois functions

*keys.go*: Fuctions to provide the AES key schedule.

*c/sbox.c*: C program used to produce the s-box.

### Acknowledgements ###

Test cases taken from FIPS standard at https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.197.pdf

MixColumns test cases from Wikipedia: https://en.wikipedia.org/wiki/Rijndael_MixColumns

Reference sbox from https://ipfs.io/ipfs/QmXoypizjW3WknFiJnKLwHCnL72vedxjQkDDP1mXWo6uco/wiki/Rijndael_S-box.html

Code to generate sbox adapted from Wikipedia: https://en.wikipedia.org/wiki/Rijndael_S-box

### Contact

If you find a problem in the code or want to improve it, feel free to submit an issue or pull request.

If you'd like to get in contact, you can do so at:

xrmon **at** pm **dot** me
