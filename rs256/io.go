package rs256

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

const (
	PathToRSAPriv = "./rs256/keys/id_rsa"
	PathToRSAPub = "./rs256/keys/id_rsa_pub"
)

func MakeRSAKeysToDisk(password string) error {
	privateKey, err := GeneratePrivateKey(1024)
	if err != nil {
		return err
	}
	keyPEM, err := PrivateKeyToPEM(privateKey, password)
	if err != nil {
		return err
	}
	WriteKeyToFile(keyPEM, PathToRSAPriv)
	publicKey, err := GeneratePublicKeyPEM(privateKey, password)
	if err != nil {
		return err
	}
	WriteKeyToFile(publicKey, PathToRSAPub)
	return err
}

// ReadPrivateKey reads PEM encoded file, decrypts is with password and parses as RSA private key.
func ReadPrivateKey(password string) (*rsa.PrivateKey, error) {
	file, err := ioutil.ReadFile(PathToRSAPriv)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(file)
	decBlock, err := x509.DecryptPEMBlock(block, []byte(password))
	if err != nil {
		return nil, err
	}
	key, err := x509.ParsePKCS1PrivateKey(decBlock)
	if err != nil {
		return nil, err
	}
	return key, err
}

// ReadPublicKey reads PEM encoded file, decrypts is with password and parses as SHA256 public key.
func ReadPublicKey(password string) (*rsa.PublicKey, error) {
	file, err := ioutil.ReadFile(PathToRSAPub)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(file)
	decrypted, err := x509.DecryptPEMBlock(block, []byte(password))
	if err != nil {
		return nil, err
	}
	parsed, err := x509.ParsePKIXPublicKey(decrypted)
	if err != nil {
		return nil, err
	}
	key := parsed.(*rsa.PublicKey) // type assertion of ambiguous return??
	return key, nil
}

// WriteKeyToFile writes key to file.
func WriteKeyToFile(keyBytes []byte, saveToFile string) error {
	err := ioutil.WriteFile(saveToFile, keyBytes, 0600)
	return err
}
