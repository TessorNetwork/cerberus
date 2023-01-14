//go:build gofuzz
// +build gofuzz

package consensus

import (
	"bytes"
	"io"
)

func Fuzz(data []byte) int {
	fur := NewWALDecoder(bytes.NewReader(data))
	for {
		msg, err := fur.Decode()
		if err == io.EOF {
			break
		}
		if err != nil {
			if msg != nil {
				panic("msg != nil on error")
			}
			return 0
		}
		var w bytes.Buffer
		enc := NewWALEncoder(&w)
		err = enc.Encode(msg)
		if err != nil {
			panic(err)
		}
	}
	return 1
}
