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
)

func Encrypt(privKey *secp256k1.PrivateKey, pubKey *secp256k1.PublicKey, pathToFile string) {

	//generate DH shared secret using keys
	//GenerateSharedSecret returns x coordinate of the generated point
	sharedPkey := GenerateSharedPkey(privKey, pubKey)
	fmt.Println("🔓 sharedPkey: ", sharedPkey)

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
	ioutil.WriteFile("src/encrypted", encryptedMessage, 0644)
	fmt.Println("✅ file encrypted with ECDH to /src/encrypted")

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

func GenerateSharedPkey(pkey *secp256k1.PrivateKey, pubkey *secp256k1.PublicKey) *secp256k1.PrivateKey {
	sharePkeyBytes := secp256k1.GenerateSharedSecret(pkey, pubkey)

	// we will use the x as pkey
	// ideally we would use HKDF(x) for more security
	sharedPkey, _ := secp256k1.PrivKeyFromBytes(sharePkeyBytes)

	return sharedPkey
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
	pubKeyStr := strings.Split(os.Getenv("RECIPIENT_PUBKEY"), " ")

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

	pubKey2 := secp256k1.NewPublicKey(x, y)

	return privKey, pubKey2
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
