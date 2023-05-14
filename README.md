# ECDH with go

This client uses ECDH (Elliptic Curve Diffie-Hellman) key exchange to encrypt a file as a secret between two party.

## Set keys

Fill .env to set your private key and the public key of the other party.

Note: the public key needs to be in uncompressed format, and without "0x" prefix.

## Choose file to decrypt/encrypt

If you want to encrypt a file, replace to_encrypt file in /src by the file to encrypt.

If you want to decrypt a message, fill to_decrypt with the message.

## 🔧 todo

- encrypt same file for multiple recipient
- use HKDF