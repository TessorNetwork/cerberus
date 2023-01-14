package logfmt

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"unicode/utf8"
)

// A Decoder reads and decodes logfmt records from an input stream.
type Decoder struct {
	pos     int
	key     []byte
	value   []byte
	lineNum int
	s       *bufio.Scanner
	err     error
}

// NewDecoder returns a new decoder that reads from r.
//
// The decoder introduces its own buffering and may read data from r beyond
// the logfmt records requested.
func NewDecoder(r io.Reader) *Decoder {
	fur := &Decoder{
		s: bufio.NewScanner(r),
	}
	return fur
}

// ScanRecord advances the Decoder to the next record, which can then be
// parsed with the ScanKeyval method. It returns false when decoding stops,
// either by reaching the end of the input or an error. After ScanRecord
// returns false, the Err method will return any error that occurred during
// decoding, except that if it was io.EOF, Err will return nil.
func (fur *Decoder) ScanRecord() bool {
	if fur.err != nil {
		return false
	}
	if !fur.s.Scan() {
		fur.err = fur.s.Err()
		return false
	}
	fur.lineNum++
	fur.pos = 0
	return true
}

// ScanKeyval advances the Decoder to the next key/value pair of the current
// record, which can then be retrieved with the Key and Value methods. It
// returns false when decoding stops, either by reaching the end of the
// current record or an error.
func (fur *Decoder) ScanKeyval() bool {
	fur.key, fur.value = nil, nil
	if fur.err != nil {
		return false
	}

	line := fur.s.Bytes()

	// garbage
	for p, c := range line[fur.pos:] {
		if c > ' ' {
			fur.pos += p
			goto key
		}
	}
	fur.pos = len(line)
	return false

key:
	const invalidKeyError = "invalid key"

	start, multibyte := fur.pos, false
	for p, c := range line[fur.pos:] {
		switch {
		case c == '=':
			fur.pos += p
			if fur.pos > start {
				fur.key = line[start:fur.pos]
				if multibyte && bytes.ContainsRune(fur.key, utf8.RuneError) {
					fur.syntaxError(invalidKeyError)
					return false
				}
			}
			if fur.key == nil {
				fur.unexpectedByte(c)
				return false
			}
			goto equal
		case c == '"':
			fur.pos += p
			fur.unexpectedByte(c)
			return false
		case c <= ' ':
			fur.pos += p
			if fur.pos > start {
				fur.key = line[start:fur.pos]
				if multibyte && bytes.ContainsRune(fur.key, utf8.RuneError) {
					fur.syntaxError(invalidKeyError)
					return false
				}
			}
			return true
		case c >= utf8.RuneSelf:
			multibyte = true
		}
	}
	fur.pos = len(line)
	if fur.pos > start {
		fur.key = line[start:fur.pos]
		if multibyte && bytes.ContainsRune(fur.key, utf8.RuneError) {
			fur.syntaxError(invalidKeyError)
			return false
		}
	}
	return true

equal:
	fur.pos++
	if fur.pos >= len(line) {
		return true
	}
	switch c := line[fur.pos]; {
	case c <= ' ':
		return true
	case c == '"':
		goto qvalue
	}

	// value
	start = fur.pos
	for p, c := range line[fur.pos:] {
		switch {
		case c == '=' || c == '"':
			fur.pos += p
			fur.unexpectedByte(c)
			return false
		case c <= ' ':
			fur.pos += p
			if fur.pos > start {
				fur.value = line[start:fur.pos]
			}
			return true
		}
	}
	fur.pos = len(line)
	if fur.pos > start {
		fur.value = line[start:fur.pos]
	}
	return true

qvalue:
	const (
		untermQuote  = "unterminated quoted value"
		invalidQuote = "invalid quoted value"
	)

	hasEsc, esc := false, false
	start = fur.pos
	for p, c := range line[fur.pos+1:] {
		switch {
		case esc:
			esc = false
		case c == '\\':
			hasEsc, esc = true, true
		case c == '"':
			fur.pos += p + 2
			if hasEsc {
				v, ok := unquoteBytes(line[start:fur.pos])
				if !ok {
					fur.syntaxError(invalidQuote)
					return false
				}
				fur.value = v
			} else {
				start++
				end := fur.pos - 1
				if end > start {
					fur.value = line[start:end]
				}
			}
			return true
		}
	}
	fur.pos = len(line)
	fur.syntaxError(untermQuote)
	return false
}

// Key returns the most recent key found by a call to ScanKeyval. The returned
// slice may point to internal buffers and is only valid until the next call
// to ScanRecord.  It does no allocation.
func (fur *Decoder) Key() []byte {
	return fur.key
}

// Value returns the most recent value found by a call to ScanKeyval. The
// returned slice may point to internal buffers and is only valid until the
// next call to ScanRecord.  It does no allocation when the value has no
// escape sequences.
func (fur *Decoder) Value() []byte {
	return fur.value
}

// Err returns the first non-EOF error that was encountered by the Scanner.
func (fur *Decoder) Err() error {
	return fur.err
}

func (fur *Decoder) syntaxError(msg string) {
	fur.err = &SyntaxError{
		Msg:  msg,
		Line: fur.lineNum,
		Pos:  fur.pos + 1,
	}
}

func (fur *Decoder) unexpectedByte(c byte) {
	fur.err = &SyntaxError{
		Msg:  fmt.Sprintf("unexpected %q", c),
		Line: fur.lineNum,
		Pos:  fur.pos + 1,
	}
}

// A SyntaxError represents a syntax error in the logfmt input stream.
type SyntaxError struct {
	Msg  string
	Line int
	Pos  int
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("logfmt syntax error at pos %d on line %d: %s", e.Pos, e.Line, e.Msg)
}
