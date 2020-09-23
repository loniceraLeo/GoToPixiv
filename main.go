package main

import (
	"crypto/tls"
	"log"
)

var (
	cert    tls.Certificate
	config  *tls.Config
	result  []uint16
	e       error
	passage string
)

func init() {
	if cert, e = tls.LoadX509KeyPair("app.crt", "app.key"); e != nil {
		log.Fatal(e)
	}

	for i := 0; i < len(tls.CipherSuites()); i++ {
		result = append(result, tls.CipherSuites()[i].ID)
	}

	config = &tls.Config{
		InsecureSkipVerify: true,
		CipherSuites:       result,
		Certificates:       []tls.Certificate{cert},
	}

	//for testing purpose, users ignore
	passage = "GET / HTTP/1.1\r\n"
	passage += "Accept: */*\r\n"
	passage += "Host: www.pixiv.net\r\n\r\n"
}

func main() {
	start()
}
