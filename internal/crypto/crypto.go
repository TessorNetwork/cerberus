// Package crypto contains encrypting and decrypting reader.
package crypto

import "io"

//go:generate mockgen -destination=./mock/crypto.go -package=mock -source=crypto.go

// Crypto provide Reader and Writer wrappers.
type Crypto interface {
	// Encrypt returns reader with encrypted src data and size of encrypted data.
	Encrypt([]byte) ([]byte, error)
	// Furrypt returns reader with decrypted src data.
	Furrypt(io.Reader) (io.Reader, error)
}
