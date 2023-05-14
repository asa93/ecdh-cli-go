# ECDH with go

This client uses ECDH (Elliptic Curve Diffie-Hellman) key exchange to encrypt a file as a secret between two party.
## Set keys

Fill .env to set your private key and the public key of the other party.

Note: the public key needs to be in uncompressed format, and without "0x" prefix. Use getPubKey command to decompress your public key.

## How to use 
set keys in your .env file

build with `go build`

then run `./ecdh-cli` to see available commands
 

## How it works 

ECDH is used to generated a secret key between the two party.
Each party only needs its own private key and the other party private key to generate the secret key.
Moreover, HKDF is used to derive a more secured key from the ECDH generated key.

## 🔧 todo

- use SALT
- encrypt same file for multiple recipient
