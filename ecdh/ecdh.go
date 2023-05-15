package ecdh

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"os"
	"strings"

	"github.com/decred/dcrd/dcrec/secp256k1"
	"github.com/joho/godotenv"

	"crypto/rand"
	"crypto/sha256"

	"golang.org/x/crypto/hkdf"
)

func Encrypt(privKey *secp256k1.PrivateKey, pubKey *secp256k1.PublicKey, pathToFile string, pathOut string) {

	//generate DH shared secret using keys
	sharedPkey := GenerateSharedPkey(privKey, pubKey)
	fmt.Println("🔑 sharedPkey generated: ", sharedPkey)

	// should turn into private key with HKDF ?

	//read file to encrypt
	message, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	//encrypt message with shared secret
	encryptedMessage, err := secp256k1.Encrypt(sharedPkey.PubKey(), []byte(message))
	if err != nil {
		fmt.Println(err)
		return
	}

	//save encrypted file
	ioutil.WriteFile(pathOut, encryptedMessage, 0644)
	fmt.Println("🔓 file encrypted with ECDH to " + pathOut)

}

func Decrypt(privKey *secp256k1.PrivateKey, pubKey *secp256k1.PublicKey, pathToFile string) {
	//generate DH shared secret using keys
	//GenerateSharedSecret returns x coordinate of the generated point
	sharedPkey := GenerateSharedPkey(privKey, pubKey)
	fmt.Println("🔓 sharedPkey: ", sharedPkey)

	//read bytearray from file
	encrypted := ReadByteFile(pathToFile)

	decrypted, err := secp256k1.Decrypt(sharedPkey, encrypted)
	if err != nil {
		fmt.Println(err)
		return
	}

	ioutil.WriteFile("src/decrypted", decrypted, 0644)

	fmt.Println(`
	================
	🔓 decrypted file
	================
	`, string(decrypted))
}

// generates a secure private key using one private key and a public key
func GenerateSharedPkey(pkey *secp256k1.PrivateKey, pubkey *secp256k1.PublicKey) *secp256k1.PrivateKey {

	//GenerateSharedSecret returns the x coordinate of the shared public key
	// Qs = dA*dB*G = dA*Qb = dB*Qa
	sharePkeyBytes := secp256k1.GenerateSharedSecret(pkey, pubkey)

	//use HKDF to generate a more secure key
	derivedKeyBytes := genHKDF(sharePkeyBytes)

	sharedPkey, _ := secp256k1.PrivKeyFromBytes(derivedKeyBytes)

	return sharedPkey
}

// derives a secure private key from an array of bytes usin HKDF
func genHKDF(input []byte) []byte {

	// Underlying hash function for HMAC.
	hash := sha256.New

	// Non-secret salt, optional (can be nil).
	// Recommended: hash-length random value.
	salt := make([]byte, hash().Size())
	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}

	// Generate a 32-bytes derived key.
	hkdf := hkdf.New(hash, input, nil, nil)

	key := make([]byte, 32)
	if _, err := io.ReadFull(hkdf, key); err != nil {
		panic(err)
	}

	return key
}

func GetKeys() (*secp256k1.PrivateKey, *secp256k1.PublicKey) {
	godotenv.Load()
	// Decode the hex-encoded private key
	pkBytes, err := hex.DecodeString(os.Getenv("SENDER_PKEY"))
	if err != nil {
		fmt.Println(err)
	}
	privKey, _ := secp256k1.PrivKeyFromBytes(pkBytes)

	///decode uncompressed pubkey of other party

	pubKey2 := ParsePublicKey(os.Getenv("RECIPIENT_PUBKEY"))

	return privKey, pubKey2
}

// parse & return an uncompressed public key from a string with x,y coordinates
func ParsePublicKey(keyStr string) *secp256k1.PublicKey {
	pubKeyStr := strings.Split(keyStr, " ")

	x := new(big.Int)
	x, ok := x.SetString(pubKeyStr[0], 10)
	if !ok {
		fmt.Println("SetString: error")
	}

	y := new(big.Int)
	y, ok_ := y.SetString(pubKeyStr[1], 10)
	if !ok_ {
		fmt.Println("SetString: error")
	}

	return secp256k1.NewPublicKey(x, y)

}
func ReadByteFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	// Get the file size
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Read the file into a byte slice
	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return nil
	}

	return bs

}
