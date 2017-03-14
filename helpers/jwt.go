package helpers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"
	"simple-api/constants"
	"time"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jwt"
	"fmt"
	"strings"
)

var (
	privateKey         *rsa.PrivateKey
	publicKey          *rsa.PublicKey
	privateKeyFilepath = os.Getenv("PRIVATE_KEY_FILEPATH")
	publicKeyFilepath  = os.Getenv("PUBLIC_KEY_FILEPATH")
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
	if _, err := os.Stat(privateKeyFilepath); os.IsNotExist(err) {
		err = initPrivateKeyResource()
		if err != nil {
			return err
		}
	} else {
		err = loadPrivateKey()
		if err != nil {
			return err
		}
	}

	if _, err := os.Stat(publicKeyFilepath); os.IsNotExist(err) {
		err = initPublicKeyResource()
		if err != nil {
			return err
		}
	} else {
		err = loadPublicKey()
		if err != nil {
			return err
		}
	}

	return nil
}

func initPrivateKeyResource() (err error) {
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)

	privateKeyPemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	privateKeyWriter, err := os.OpenFile(privateKeyFilepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	pem.Encode(privateKeyWriter, privateKeyPemBlock)

	return
}

func initPublicKeyResource() (err error) {
	publicKey = privateKey.Public().(*rsa.PublicKey)

	publicKeyByte, err := x509.MarshalPKIXPublicKey(publicKey)
	publicKeyPemBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyByte,
	}

	publicKeyWriter, err := os.OpenFile(publicKeyFilepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
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

	publicKeyPemBlock, _ := pem.Decode(publicKeyByte)
	publicKeyGeneral, err := x509.ParsePKIXPublicKey(publicKeyPemBlock.Bytes)
	publicKey = publicKeyGeneral.(*rsa.PublicKey)

	return err
}

func CreateJWT(claimsSet map[string]interface{}, expiration time.Time) (string, error) {
	claims := make(jws.Claims)
	claims.SetExpiration(expiration)
	for key := range claimsSet {
		claims.Set(key, claimsSet[key])
	}
	jwtObj := jws.NewJWT(claims, crypto.SigningMethodRS512)
	token, err := jwtObj.Serialize(GetPrivateKey())
	if err != nil {
		return "", err
	}
	return string(token), nil
}

func ValidateJWT(token string) (claims map[string]interface{}, err error) {
	var jwtObj jwt.JWT
	if jwtObj, err = jws.ParseJWT([]byte(token)); err != nil {
		return
	}
	validator := jws.NewValidator(nil, 0, 0, func(c jws.Claims) error {
		if c["username"] == nil {
			err = fmt.Errorf("token contains no username")
			return err
		}
		if c["id"] == nil {
			err = fmt.Errorf("token contains no id")
			return err
		}
		claims = c
		return nil
	})
	err = jwtObj.Validate(GetPublicKey(), crypto.SigningMethodRS512, validator)
	return
}

func ParseHeaderJWT(header string) (token string, err error) {
	splitted := strings.Split(header, " ")

	if len(splitted) != 2 {
		err = fmt.Errorf("Invalid authorization header")
		return
	}

	if strings.ToLower(splitted[0]) != "bearer" {
		err = fmt.Errorf("Invalid token type")
		return
	}
	token = splitted[1]
	return
}