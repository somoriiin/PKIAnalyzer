// main
package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

func main() {

	data, err := os.ReadFile("testdata/server.crt")
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Println(string(data))

	bl, _ := pem.Decode(data)
	if bl == nil {
		log.Fatalf("Failed to decode PEM block")
	}

	cert, err := x509.ParseCertificate(bl.Bytes)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	fmt.Println("Subject: ", cert.Subject.String())
	fmt.Println("Issuer: ", cert.Issuer.String())
	fmt.Println("ValidFrom: ", cert.NotBefore.String())
	fmt.Println("ValidUntil: ", cert.NotAfter.String())
	fmt.Println("PublicKey: ", fmt.Sprintf("%v", cert.PublicKey))
}
