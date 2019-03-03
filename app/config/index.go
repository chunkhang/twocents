package config

import (
	"encoding/base64"
	"os"
)

const (
	// Port is port number for server
	Port = "3000"
)

var (
	// StoreKey is key for cookie store
	StoreKey, _ = base64.StdEncoding.DecodeString(os.Getenv("STORE_KEY"))
)
