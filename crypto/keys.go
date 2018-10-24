package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

func main() {
	privateKey, err := GeneratePrivateKey()

	keyPEM := PrivateKeyToPEM(privateKey)

	writeKeyToFile(keyPEM, "./crypto/privateKey")

	publicKey, err := GeneratePublicKeyPEM(privateKey)

	writeKeyToFile(publicKey, "./crypto/publicKey")

	if err != nil {
		panic(err)
	}
}

func GeneratePrivateKey() (*rsa.PrivateKey, error) {
	bits := 1024

	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}

	err = key.Validate()
	if err != nil {
		return nil, err
	}
	return key, err
}

func PrivateKeyToPEM(key *rsa.PrivateKey) []byte {
	// Get ASN.1 DER format
	privDER := x509.MarshalPKCS1PrivateKey(key)

	privBlock := &pem.Block{
		Type: "RSA PRIVATE KEY",
		Headers:nil,
		Bytes: privDER,
	}

	//// Encrypts the file with password (store this on drive prod)
	//encBlock, err := x509.EncryptPEMBlock(
	//	rand.Reader,
	//	privBlock.Type,
	//	privBlock.Bytes,
	//	[]byte("password"),
	//	x509.PEMCipherAES256)
	//if err != nil {
	//	panic(err)
	//}

	privatePEM := pem.EncodeToMemory(privBlock)

	return privatePEM
}

func writeKeyToFile(keyBytes []byte, saveToFile string) error {
	err := ioutil.WriteFile(saveToFile, keyBytes, 0600)
	if err != nil {
		return err
	}
	return nil
}

func GeneratePublicKeyPEM(privKey *rsa.PrivateKey) ([]byte, error) {
	asnKey, err := x509.MarshalPKIXPublicKey(&privKey.PublicKey)
	if err != nil {
		panic(err)
	}

	privBlock := &pem.Block{
		Type: "RSA PUBLIC KEY",
		Headers:nil,
		Bytes: asnKey,
	}

	publicPEM := pem.EncodeToMemory(privBlock)

	return publicPEM, nil

	//pubKey, err := ssh.NewPublicKey(privKey)
	//if err != nil {
	//	return nil, err
	//}
	//
	//pubKeyBytes := ssh.MarshalAuthorizedKey(pubKey)
	//
	//return pubKeyBytes, nil
}

func ReadPrivateKey(path string) *rsa.PrivateKey{
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(file)
	//// This decrypts password protected key
	//decBlock, _ := x509.DecryptPEMBlock(block, []byte("password"))
	//
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	return key
}