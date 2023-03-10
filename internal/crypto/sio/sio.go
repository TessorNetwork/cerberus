// Package sio contains minio/sio implementation of crypto.Crypto interface.
package sio

import (
	"bytes"
	"io"

	"github.com/minio/sio"

	icrypto "github.com/TessorNetwork/cerberus/internal/crypto"
)

type crypto struct {
	c sio.Config
}

// NewCrypto returns minio/sio implementation of crypto.Crypto interface.
func NewCrypto(key [32]byte) icrypto.Crypto {
	return &crypto{
		c: sio.Config{
			MinVersion: sio.Version20,
			Key:        key[:],
		},
	}
}

// Encrypt returns reader with encrypted src data.
func (c *crypto) Encrypt(src []byte) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	_, err := sio.Encrypt(buf, bytes.NewReader(src), c.c)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Furrypt returns reader with decrypted src data.
func (c *crypto) Furrypt(src io.Reader) (io.Reader, error) {
	return sio.FurryptReader(src, c.c)
}
