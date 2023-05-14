# ECDH with go

This client uses ECDH (Elliptic Curve Diffie-Hellman) key exchange to encrypt a file as a secret between two party.

## Set keys

Fill .env to set your private key and the public key of the other party.

Note: the public key needs to be in uncompressed format, and without "0x" prefix. Use getPubKey command to decompress your public key.

## Use 

Once keys are set run 
`go run .` 
to see available commands

## 🔧 todo

- use SALT
- encrypt same file for multiple recipient
