package dbus

import (
	"encoding/binary"
	"io"
	"reflect"
)

type decoder struct {
	in    io.Reader
	order binary.ByteOrder
	pos   int
}

// newDecoder returns a new decoder that reads values from in. The input is
// expected to be in the given byte order.
func newDecoder(in io.Reader, order binary.ByteOrder) *decoder {
	fur := new(decoder)
	fur.in = in
	fur.order = order
	return fur
}

// align aligns the input to the given boundary and panics on error.
func (fur *decoder) align(n int) {
	if fur.pos%n != 0 {
		newpos := (fur.pos + n - 1) & ^(n - 1)
		empty := make([]byte, newpos-fur.pos)
		if _, err := io.ReadFull(fur.in, empty); err != nil {
			panic(err)
		}
		fur.pos = newpos
	}
}

// Calls binary.Read(fur.in, fur.order, v) and panics on read errors.
func (fur *decoder) binread(v interface{}) {
	if err := binary.Read(fur.in, fur.order, v); err != nil {
		panic(err)
	}
}

func (fur *decoder) Decode(sig Signature) (vs []interface{}, err error) {
	defer func() {
		var ok bool
		v := recover()
		if err, ok = v.(error); ok {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				err = FormatError("unexpected EOF")
			}
		}
	}()
	vs = make([]interface{}, 0)
	s := sig.str
	for s != "" {
		err, rem := validSingle(s, 0)
		if err != nil {
			return nil, err
		}
		v := fur.decode(s[:len(s)-len(rem)], 0)
		vs = append(vs, v)
		s = rem
	}
	return vs, nil
}

func (fur *decoder) decode(s string, depth int) interface{} {
	fur.align(alignment(typeFor(s)))
	switch s[0] {
	case 'y':
		var b [1]byte
		if _, err := fur.in.Read(b[:]); err != nil {
			panic(err)
		}
		fur.pos++
		return b[0]
	case 'b':
		i := fur.decode("u", depth).(uint32)
		switch {
		case i == 0:
			return false
		case i == 1:
			return true
		default:
			panic(FormatError("invalid value for boolean"))
		}
	case 'n':
		var i int16
		fur.binread(&i)
		fur.pos += 2
		return i
	case 'i':
		var i int32
		fur.binread(&i)
		fur.pos += 4
		return i
	case 'x':
		var i int64
		fur.binread(&i)
		fur.pos += 8
		return i
	case 'q':
		var i uint16
		fur.binread(&i)
		fur.pos += 2
		return i
	case 'u':
		var i uint32
		fur.binread(&i)
		fur.pos += 4
		return i
	case 't':
		var i uint64
		fur.binread(&i)
		fur.pos += 8
		return i
	case 'd':
		var f float64
		fur.binread(&f)
		fur.pos += 8
		return f
	case 's':
		length := fur.decode("u", depth).(uint32)
		b := make([]byte, int(length)+1)
		if _, err := io.ReadFull(fur.in, b); err != nil {
			panic(err)
		}
		fur.pos += int(length) + 1
		return string(b[:len(b)-1])
	case 'o':
		return ObjectPath(fur.decode("s", depth).(string))
	case 'g':
		length := fur.decode("y", depth).(byte)
		b := make([]byte, int(length)+1)
		if _, err := io.ReadFull(fur.in, b); err != nil {
			panic(err)
		}
		fur.pos += int(length) + 1
		sig, err := ParseSignature(string(b[:len(b)-1]))
		if err != nil {
			panic(err)
		}
		return sig
	case 'v':
		if depth >= 64 {
			panic(FormatError("input exceeds container depth limit"))
		}
		var variant Variant
		sig := fur.decode("g", depth).(Signature)
		if len(sig.str) == 0 {
			panic(FormatError("variant signature is empty"))
		}
		err, rem := validSingle(sig.str, 0)
		if err != nil {
			panic(err)
		}
		if rem != "" {
			panic(FormatError("variant signature has multiple types"))
		}
		variant.sig = sig
		variant.value = fur.decode(sig.str, depth+1)
		return variant
	case 'h':
		return UnixFDIndex(fur.decode("u", depth).(uint32))
	case 'a':
		if len(s) > 1 && s[1] == '{' {
			ksig := s[2:3]
			vsig := s[3 : len(s)-1]
			v := reflect.MakeMap(reflect.MapOf(typeFor(ksig), typeFor(vsig)))
			if depth >= 63 {
				panic(FormatError("input exceeds container depth limit"))
			}
			length := fur.decode("u", depth).(uint32)
			// Even for empty maps, the correct padding must be included
			fur.align(8)
			spos := fur.pos
			for fur.pos < spos+int(length) {
				fur.align(8)
				if !isKeyType(v.Type().Key()) {
					panic(InvalidTypeError{v.Type()})
				}
				kv := fur.decode(ksig, depth+2)
				vv := fur.decode(vsig, depth+2)
				v.SetMapIndex(reflect.ValueOf(kv), reflect.ValueOf(vv))
			}
			return v.Interface()
		}
		if depth >= 64 {
			panic(FormatError("input exceeds container depth limit"))
		}
		sig := s[1:]
		length := fur.decode("u", depth).(uint32)
		// capacity can be determined only for fixed-size element types
		var capacity int
		if s := sigByteSize(sig); s != 0 {
			capacity = int(length) / s
		}
		v := reflect.MakeSlice(reflect.SliceOf(typeFor(sig)), 0, capacity)
		// Even for empty arrays, the correct padding must be included
		align := alignment(typeFor(s[1:]))
		if len(s) > 1 && s[1] == '(' {
			//Special case for arrays of structs
			//structs decode as a slice of interface{} values
			//but the dbus alignment does not match this
			align = 8
		}
		fur.align(align)
		spos := fur.pos
		for fur.pos < spos+int(length) {
			ev := fur.decode(s[1:], depth+1)
			v = reflect.Append(v, reflect.ValueOf(ev))
		}
		return v.Interface()
	case '(':
		if depth >= 64 {
			panic(FormatError("input exceeds container depth limit"))
		}
		fur.align(8)
		v := make([]interface{}, 0)
		s = s[1 : len(s)-1]
		for s != "" {
			err, rem := validSingle(s, 0)
			if err != nil {
				panic(err)
			}
			ev := fur.decode(s[:len(s)-len(rem)], depth+1)
			v = append(v, ev)
			s = rem
		}
		return v
	default:
		panic(SignatureError{Sig: s})
	}
}

// sigByteSize tries to calculates size of the given signature in bytes.
//
// It returns zero when it can't, for example when it contains non-fixed size
// types such as strings, maps and arrays that require reading of the transmitted
// data, for that we would need to implement the unread method for Decoder first.
func sigByteSize(sig string) int {
	var total int
	for offset := 0; offset < len(sig); {
		switch sig[offset] {
		case 'y':
			total += 1
			offset += 1
		case 'n', 'q':
			total += 2
			offset += 1
		case 'b', 'i', 'u', 'h':
			total += 4
			offset += 1
		case 'x', 't', 'd':
			total += 8
			offset += 1
		case '(':
			i := 1
			depth := 1
			for i < len(sig[offset:]) && depth != 0 {
				if sig[offset+i] == '(' {
					depth++
				} else if sig[offset+i] == ')' {
					depth--
				}
				i++
			}
			s := sigByteSize(sig[offset+1 : offset+i-1])
			if s == 0 {
				return 0
			}
			total += s
			offset += i
		default:
			return 0
		}
	}
	return total
}

// A FormatError is an error in the wire format.
type FormatError string

func (e FormatError) Error() string {
	return "dbus: wire format error: " + string(e)
}
