package config

import (
	"time"
)

var (
	RsaPublicKey  = []byte(``)
	RsaPrivateKey = []byte(``)

	C2                          = "127.0.0.1:443"
	plainHTTP                   = "http://"
	sslHTTP                     = "https://"
	GetUrl                      = sslHTTP + C2 + "/load"
	PostUrl                     = sslHTTP + C2 + "/submit.php?id="
	WaitTime                    = 10000 * time.Millisecond
	VerifySSLCert               = true
	TimeOut       time.Duration = 10 //seconds

	IV        = []byte("abcdefghijklmnop")
	GlobalKey []byte
	AesKey    []byte
	HmacKey   []byte
	Counter   = 0
)

const (
	DebugMode = true
)
