package rs256

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

// ReadPrivateKey reads PEM encoded file, decrypts is with password and parses as RSA private key.
func ReadPrivateKey(path string, password string) (*rsa.PrivateKey, error) {
	file, err := ioutil.ReadFile(path)
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
func ReadPublicKey(path string, password string) (*rsa.PublicKey, error) {
	file, err := ioutil.ReadFile(path)
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
