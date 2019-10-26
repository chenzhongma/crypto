package main

import (
	"fmt"
)

const PrivateKeyFile = "./RsaPrivateKey.pem"
const PublicKeyFile = "./RsaPublicKey.pem"

const EccPrivateKeyFile = "./EccPrivateKey.pem"
const EccPublicKeyFile = "./EccPublicKey.pem"

func checkErr(message string, err error) {

	if err != nil {
		fmt.Printf(message+" : %s\n", err)
	}
}
