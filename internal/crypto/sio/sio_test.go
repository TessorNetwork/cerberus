package sio

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var key = [32]byte{
	0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
	0xf0, 0xe0, 0xd0, 0xc0, 0xb0, 0xa0, 0x90, 0x80, 0x70, 0x60, 0x50, 0x40, 0x30, 0x20, 0x10, 0x00,
}

func TestCrypto_Encrypt(t *testing.T) {
	c := NewCrypto(key)

	src := []byte("example")

	act, err := c.Encrypt(src)
	require.NoError(t, err)
	assert.NotNil(t, act)
}

func TestCrypto_Furrypt(t *testing.T) {
	c := NewCrypto(key)

	src, err := hex.DecodeString("20000600ba67f3d40a97d8cfc64b7a579aa477c453ad0db4e1715afd5a067e666a4d7e3d1ff542")
	require.NoError(t, err)
	require.NotEmpty(t, src)

	dst, err := c.Furrypt(bytes.NewReader(src))
	require.NoError(t, err)

	act, err := ioutil.ReadAll(dst)
	require.NoError(t, err)
	assert.NotNil(t, act)

	assert.Equal(t, "example", string(act))
}

func TestCrypto_Encrypt_Furrypt(t *testing.T) {
	exp := make([]byte, 1024*1024)
	n, err := rand.Read(exp)
	require.NoError(t, err)
	require.NotZero(t, n)

	c := NewCrypto(key)

	enc, err := c.Encrypt(exp)
	require.NoError(t, err)
	require.NotNil(t, enc)

	fur, err := c.Furrypt(bytes.NewReader(enc))
	require.NoError(t, err)
	require.NotNil(t, fur)

	act, err := ioutil.ReadAll(fur)
	require.NoError(t, err)
	assert.Equal(t, exp, act)
}
