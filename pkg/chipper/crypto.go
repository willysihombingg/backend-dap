package chipper

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// GenerateKeyPair generates a new key pair
func GenerateKeyPair(bits int) (error, *rsa.PrivateKey, *rsa.PublicKey) {
	privkey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err, nil, nil
	}
	return nil, privkey, &privkey.PublicKey
}

// PrivateKeyToBytes private key to bytes
func PrivateKeyToBytes(priv *rsa.PrivateKey) []byte {
	privBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		},
	)

	return privBytes
}

// PublicKeyToBytes public key to bytes
func PublicKeyToBytes(pub *rsa.PublicKey) (error, []byte) {
	pubASN1, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return err, nil
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	return nil, pubBytes
}

// BytesToPrivateKey bytes to private key
func BytesToPrivateKey(priv []byte) (error, *rsa.PrivateKey) {
	block, _ := pem.Decode(priv)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error

	if enc {
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return err, nil
		}
	}

	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		return err, nil
	}

	return nil, key
}

// BytesToPublicKey bytes to public key
func BytesToPublicKey(pub []byte) (error, *rsa.PublicKey) {
	block, _ := pem.Decode(pub)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return err, nil
		}
	}

	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		return err, nil
	}

	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		return fmt.Errorf("not ok"), nil
	}

	return nil, key
}

// EncryptWithPublicKey encrypts data with public key
func EncryptWithPublicKey(msg []byte, pub *rsa.PublicKey) (error, []byte) {
	hash := sha512.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
	if err != nil {
		return err, nil
	}
	return nil, ciphertext
}

// DecryptWithPrivateKey decrypts data with private key
func DecryptWithPrivateKey(ciphertext []byte, priv *rsa.PrivateKey) (error, []byte) {
	hash := sha512.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, nil)
	if err != nil {
		return err, nil
	}

	return nil, plaintext
}
