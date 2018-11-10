package rs256

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

// GeneratePrivateKey creates RSA private key of specified bitSize and validates it.
// Note that *rsa.PrivateKey.PublicKey can be used to validate this RSA signature.
func GeneratePrivateKey(bitSize int) (*rsa.PrivateKey, error) {
	key, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}
	err = key.Validate()
	if err != nil {
		return nil, err
	}
	return key, err
}

// PrivateKeyToPEM encrypts RSA private key with password and encodes the result to PEM format.
func PrivateKeyToPEM(key *rsa.PrivateKey, password string) ([]byte, error) {
	// needs to marshal into ASN.1 DER format
	asnKey := x509.MarshalPKCS1PrivateKey(key)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: asnKey,
	}

	encrypted, err := x509.EncryptPEMBlock(
		rand.Reader,
		block.Type,
		block.Bytes,
		[]byte(password),
		x509.PEMCipherAES256)
	if err != nil {
		return nil, err
	}
	pemKey := pem.EncodeToMemory(encrypted)
	return pemKey, err
}

// GeneratePublicKeyPEM fetches public key from RSA private key, encrypts it with password and encodes the result to PEM format.
func GeneratePublicKeyPEM(privKey *rsa.PrivateKey, password string) ([]byte, error) {
	asnKey, err := x509.MarshalPKIXPublicKey(&privKey.PublicKey)
	if err != nil {
		return nil, err
	}
	block := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: asnKey,
	}

	encrypted, err := x509.EncryptPEMBlock(
		rand.Reader,
		block.Type,
		block.Bytes,
		[]byte(password),
		x509.PEMCipherAES256)
	if err != nil {
		return nil, err
	}
	publicPEM := pem.EncodeToMemory(encrypted)
	return publicPEM, nil
}
