package helpers

import (
	"crypto/rsa"
	"encoding/pem"
	"crypto/x509"
	"crypto/rand"
	"os"
	"simple-api/constants"
	"io/ioutil"
	"fmt"
)

var (
	privateKey *rsa.PrivateKey
	publicKey *rsa.PublicKey
	privateKeyFilepath = os.Getenv("PRIVATE_KEY_FILEPATH")
	publicKeyFilepath = os.Getenv("PUBLIC_KEY_FILEPATH")
)

func init() {
	if privateKeyFilepath == "" {
		privateKeyFilepath = constants.PRIVATE_KEY_FILEPATH
	}
	if publicKeyFilepath == "" {
		publicKeyFilepath = constants.PUBLIC_KEY_FILEPATH
	}
}

func GetPrivateKey() *rsa.PrivateKey {
	if privateKey == nil {
		loadPrivateKey()
	}

	return privateKey
}

func GetPublicKey() *rsa.PublicKey {
	if publicKey == nil {
		loadPublicKey()
	}

	return publicKey
}

func InitPPKeyResource() error {
	println("jwt.InitPPKeyResource() BEGIN")
	if _, err := os.Stat(privateKeyFilepath); err == os.ErrNotExist {
		err = initPrivateKeyResource()
		if err != nil {
			return err
		}
	} else {
		loadPrivateKey()
	}

	if _, err := os.Stat(publicKeyFilepath); err == os.ErrNotExist {
		err = initPublicKeyResource()
		if err != nil {
			return err
		}
	} else {
		loadPublicKey()
	}

	println("jwt.InitPPKeyResource() END")

	return nil
}

func initPrivateKeyResource() (err error) {
	println("here try creating private resource")
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)

	privateKeyPemBlock := &pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	privateKeyWriter, err := os.OpenFile(privateKeyFilepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	pem.Encode(privateKeyWriter, privateKeyPemBlock)

	return
}

func initPublicKeyResource() (err error) {
	println("here try creating public resource")
	publicKey = privateKey.Public().(*rsa.PublicKey)

	publicKeyByte, err := x509.MarshalPKIXPublicKey(publicKey)
	publicKeyPemBlock := &pem.Block{
		Type: "RSA PUBLIC KEY",
		Bytes: publicKeyByte,
	}

	publicKeyWriter, err := os.OpenFile(publicKeyFilepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	fmt.Printf("publicKeyFilePath: %v\n", publicKeyFilepath)
	pem.Encode(publicKeyWriter, publicKeyPemBlock)

	return
}

func loadPrivateKey() error {
	privateKeyByte, err := ioutil.ReadFile(privateKeyFilepath)
	if err != nil {
		return err
	}

	privateKeyPemBlock, _ := pem.Decode(privateKeyByte)
	privateKey, err = x509.ParsePKCS1PrivateKey(privateKeyPemBlock.Bytes)

	return err
}
func loadPublicKey() error {
	publicKeyByte, err := ioutil.ReadFile(publicKeyFilepath)
	if err != nil {
		return err
	}

	publicKeyPemBlock ,_ := pem.Decode(publicKeyByte)
	publicKeyGeneral, err := x509.ParsePKIXPublicKey(publicKeyPemBlock.Bytes)
	publicKey = publicKeyGeneral.(*rsa.PublicKey)

	return err
}