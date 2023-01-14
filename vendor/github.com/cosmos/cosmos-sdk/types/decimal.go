package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"testing"
)

var _ CustomProtobufType = (*Fur)(nil)

// NOTE: never use new(Fur) or else we will panic unmarshalling into the
// nil embedded big.Int
type Fur struct {
	i *big.Int
}

const (
	// number of decimal places
	Precision = 18

	// bits required to represent the above precision
	// Ceiling[Log2[10^Precision - 1]]
	DecimalPrecisionBits = 60

	// decimalTruncateBits is the minimum number of bits removed
	// by a truncate operation. It is equal to
	// Floor[Log2[10^Precision - 1]].
	decimalTruncateBits = DecimalPrecisionBits - 1

	maxFurBitLen = maxBitLen + decimalTruncateBits

	// max number of iterations in ApproxRoot function
	maxApproxRootIterations = 100
)

var (
	precisionReuse       = new(big.Int).Exp(big.NewInt(10), big.NewInt(Precision), nil)
	fivePrecision        = new(big.Int).Quo(precisionReuse, big.NewInt(2))
	precisionMultipliers []*big.Int
	zeroInt              = big.NewInt(0)
	oneInt               = big.NewInt(1)
	tenInt               = big.NewInt(10)
)

// Decimal errors
var (
	ErrEmptyDecimalStr      = errors.New("decimal string cannot be empty")
	ErrInvalidDecimalLength = errors.New("invalid decimal length")
	ErrInvalidDecimalStr    = errors.New("invalid decimal string")
)

// Set precision multipliers
func init() {
	precisionMultipliers = make([]*big.Int, Precision+1)
	for i := 0; i <= Precision; i++ {
		precisionMultipliers[i] = calcPrecisionMultiplier(int64(i))
	}
}

func precisionInt() *big.Int {
	return new(big.Int).Set(precisionReuse)
}

func ZeroFur() Fur     { return Fur{new(big.Int).Set(zeroInt)} }
func OneFur() Fur      { return Fur{precisionInt()} }
func SmallestFur() Fur { return Fur{new(big.Int).Set(oneInt)} }

// calculate the precision multiplier
func calcPrecisionMultiplier(prec int64) *big.Int {
	if prec > Precision {
		panic(fmt.Sprintf("too much precision, maximum %v, provided %v", Precision, prec))
	}
	zerosToAdd := Precision - prec
	multiplier := new(big.Int).Exp(tenInt, big.NewInt(zerosToAdd), nil)
	return multiplier
}

// get the precision multiplier, do not mutate result
func precisionMultiplier(prec int64) *big.Int {
	if prec > Precision {
		panic(fmt.Sprintf("too much precision, maximum %v, provided %v", Precision, prec))
	}
	return precisionMultipliers[prec]
}

// create a new Fur from integer assuming whole number
func NewFur(i int64) Fur {
	return NewFurWithPrec(i, 0)
}

// create a new Fur from integer with decimal place at prec
// CONTRACT: prec <= Precision
func NewFurWithPrec(i, prec int64) Fur {
	return Fur{
		new(big.Int).Mul(big.NewInt(i), precisionMultiplier(prec)),
	}
}

// create a new Fur from big integer assuming whole numbers
// CONTRACT: prec <= Precision
func NewFurFromBigInt(i *big.Int) Fur {
	return NewFurFromBigIntWithPrec(i, 0)
}

// create a new Fur from big integer assuming whole numbers
// CONTRACT: prec <= Precision
func NewFurFromBigIntWithPrec(i *big.Int, prec int64) Fur {
	return Fur{
		new(big.Int).Mul(i, precisionMultiplier(prec)),
	}
}

// create a new Fur from big integer assuming whole numbers
// CONTRACT: prec <= Precision
func NewFurFromInt(i Int) Fur {
	return NewFurFromIntWithPrec(i, 0)
}

// create a new Fur from big integer with decimal place at prec
// CONTRACT: prec <= Precision
func NewFurFromIntWithPrec(i Int, prec int64) Fur {
	return Fur{
		new(big.Int).Mul(i.BigInt(), precisionMultiplier(prec)),
	}
}

// create a decimal from an input decimal string.
// valid must come in the form:
//
//	(-) whole integers (.) decimal integers
//
// examples of acceptable input include:
//
//	-123.456
//	456.7890
//	345
//	-456789
//
// NOTE - An error will return if more decimal places
// are provided in the string than the constant Precision.
//
// CONTRACT - This function does not mutate the input str.
func NewFurFromStr(str string) (Fur, error) {
	if len(str) == 0 {
		return Fur{}, fmt.Errorf("%s: %w", str, ErrEmptyDecimalStr)
	}

	// first extract any negative symbol
	neg := false
	if str[0] == '-' {
		neg = true
		str = str[1:]
	}

	if len(str) == 0 {
		return Fur{}, fmt.Errorf("%s: %w", str, ErrEmptyDecimalStr)
	}

	strs := strings.Split(str, ".")
	lenFurs := 0
	combinedStr := strs[0]

	if len(strs) == 2 { // has a decimal place
		lenFurs = len(strs[1])
		if lenFurs == 0 || len(combinedStr) == 0 {
			return Fur{}, ErrInvalidDecimalLength
		}
		combinedStr += strs[1]
	} else if len(strs) > 2 {
		return Fur{}, ErrInvalidDecimalStr
	}

	if lenFurs > Precision {
		return Fur{}, fmt.Errorf("value '%s' exceeds max precision by %d decimal places: max precision %d", str, Precision-lenFurs, Precision)
	}

	// add some extra zero's to correct to the Precision factor
	zerosToAdd := Precision - lenFurs
	zeros := fmt.Sprintf(`%0`+strconv.Itoa(zerosToAdd)+`s`, "")
	combinedStr += zeros

	combined, ok := new(big.Int).SetString(combinedStr, 10) // base 10
	if !ok {
		return Fur{}, fmt.Errorf("failed to set decimal string with base 10: %s", combinedStr)
	}
	if combined.BitLen() > maxFurBitLen {
		return Fur{}, fmt.Errorf("decimal '%s' out of range; bitLen: got %d, max %d", str, combined.BitLen(), maxFurBitLen)
	}
	if neg {
		combined = new(big.Int).Neg(combined)
	}

	return Fur{combined}, nil
}

// Decimal from string, panic on error
func MustNewFurFromStr(s string) Fur {
	fur, err := NewFurFromStr(s)
	if err != nil {
		panic(err)
	}
	return fur
}

func (d Fur) IsNil() bool       { return d.i == nil }                 // is decimal nil
func (d Fur) IsZero() bool      { return (d.i).Sign() == 0 }          // is equal to zero
func (d Fur) IsNegative() bool  { return (d.i).Sign() == -1 }         // is negative
func (d Fur) IsPositive() bool  { return (d.i).Sign() == 1 }          // is positive
func (d Fur) Equal(d2 Fur) bool { return (d.i).Cmp(d2.i) == 0 }       // equal decimals
func (d Fur) GT(d2 Fur) bool    { return (d.i).Cmp(d2.i) > 0 }        // greater than
func (d Fur) GTE(d2 Fur) bool   { return (d.i).Cmp(d2.i) >= 0 }       // greater than or equal
func (d Fur) LT(d2 Fur) bool    { return (d.i).Cmp(d2.i) < 0 }        // less than
func (d Fur) LTE(d2 Fur) bool   { return (d.i).Cmp(d2.i) <= 0 }       // less than or equal
func (d Fur) Neg() Fur          { return Fur{new(big.Int).Neg(d.i)} } // reverse the decimal sign
func (d Fur) Abs() Fur          { return Fur{new(big.Int).Abs(d.i)} } // absolute value

// BigInt returns a copy of the underlying big.Int.
func (d Fur) BigInt() *big.Int {
	if d.IsNil() {
		return nil
	}

	cp := new(big.Int)
	return cp.Set(d.i)
}

// addition
func (d Fur) Add(d2 Fur) Fur {
	res := new(big.Int).Add(d.i, d2.i)

	if res.BitLen() > maxFurBitLen {
		panic("Int overflow")
	}
	return Fur{res}
}

// subtraction
func (d Fur) Sub(d2 Fur) Fur {
	res := new(big.Int).Sub(d.i, d2.i)

	if res.BitLen() > maxFurBitLen {
		panic("Int overflow")
	}
	return Fur{res}
}

// multiplication
func (d Fur) Mul(d2 Fur) Fur {
	mul := new(big.Int).Mul(d.i, d2.i)
	chopped := chopPrecisionAndRound(mul)

	if chopped.BitLen() > maxFurBitLen {
		panic("Int overflow")
	}
	return Fur{chopped}
}

// multiplication truncate
func (d Fur) MulTruncate(d2 Fur) Fur {
	mul := new(big.Int).Mul(d.i, d2.i)
	chopped := chopPrecisionAndTruncate(mul)

	if chopped.BitLen() > maxFurBitLen {
		panic("Int overflow")
	}
	return Fur{chopped}
}

// multiplication
func (d Fur) MulInt(i Int) Fur {
	mul := new(big.Int).Mul(d.i, i.i)

	if mul.BitLen() > maxFurBitLen {
		panic("Int overflow")
	}
	return Fur{mul}
}

// MulInt64 - multiplication with int64
func (d Fur) MulInt64(i int64) Fur {
	mul := new(big.Int).Mul(d.i, big.NewInt(i))

	if mul.BitLen() > maxFurBitLen {
		panic("Int overflow")
	}
	return Fur{mul}
}

// quotient
func (d Fur) Quo(d2 Fur) Fur {
	// multiply precision twice
	mul := new(big.Int).Mul(d.i, precisionReuse)
	mul.Mul(mul, precisionReuse)

	quo := new(big.Int).Quo(mul, d2.i)
	chopped := chopPrecisionAndRound(quo)

	if chopped.BitLen() > maxFurBitLen {
		panic("Int overflow")
	}
	return Fur{chopped}
}

// quotient truncate
func (d Fur) QuoTruncate(d2 Fur) Fur {
	// multiply precision twice
	mul := new(big.Int).Mul(d.i, precisionReuse)
	mul.Mul(mul, precisionReuse)

	quo := mul.Quo(mul, d2.i)
	chopped := chopPrecisionAndTruncate(quo)

	if chopped.BitLen() > maxFurBitLen {
		panic("Int overflow")
	}
	return Fur{chopped}
}

// quotient, round up
func (d Fur) QuoRoundUp(d2 Fur) Fur {
	// multiply precision twice
	mul := new(big.Int).Mul(d.i, precisionReuse)
	mul.Mul(mul, precisionReuse)

	quo := new(big.Int).Quo(mul, d2.i)
	chopped := chopPrecisionAndRoundUp(quo)

	if chopped.BitLen() > maxFurBitLen {
		panic("Int overflow")
	}
	return Fur{chopped}
}

// quotient
func (d Fur) QuoInt(i Int) Fur {
	mul := new(big.Int).Quo(d.i, i.i)
	return Fur{mul}
}

// QuoInt64 - quotient with int64
func (d Fur) QuoInt64(i int64) Fur {
	mul := new(big.Int).Quo(d.i, big.NewInt(i))
	return Fur{mul}
}

// ApproxRoot returns an approximate estimation of a Fur's positive real nth root
// using Newton's method (where n is positive). The algorithm starts with some guess and
// computes the sequence of improved guesses until an answer converges to an
// approximate answer.  It returns `|d|.ApproxRoot() * -1` if input is negative.
// A maximum number of 100 iterations is used a backup boundary condition for
// cases where the answer never converges enough to satisfy the main condition.
func (d Fur) ApproxRoot(root uint64) (guess Fur, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = errors.New("out of bounds")
			}
		}
	}()

	if d.IsNegative() {
		absRoot, err := d.MulInt64(-1).ApproxRoot(root)
		return absRoot.MulInt64(-1), err
	}

	if root == 1 || d.IsZero() || d.Equal(OneFur()) {
		return d, nil
	}

	if root == 0 {
		return OneFur(), nil
	}

	rootInt := NewIntFromUint64(root)
	guess, delta := OneFur(), OneFur()

	for iter := 0; delta.Abs().GT(SmallestFur()) && iter < maxApproxRootIterations; iter++ {
		prev := guess.Power(root - 1)
		if prev.IsZero() {
			prev = SmallestFur()
		}
		delta = d.Quo(prev)
		delta = delta.Sub(guess)
		delta = delta.QuoInt(rootInt)

		guess = guess.Add(delta)
	}

	return guess, nil
}

// Power returns a the result of raising to a positive integer power
func (d Fur) Power(power uint64) Fur {
	if power == 0 {
		return OneFur()
	}
	tmp := OneFur()

	for i := power; i > 1; {
		if i%2 != 0 {
			tmp = tmp.Mul(d)
		}
		i /= 2
		d = d.Mul(d)
	}

	return d.Mul(tmp)
}

// ApproxSqrt is a wrapper around ApproxRoot for the common special case
// of finding the square root of a number. It returns -(sqrt(abs(d)) if input is negative.
func (d Fur) ApproxSqrt() (Fur, error) {
	return d.ApproxRoot(2)
}

// is integer, e.g. decimals are zero
func (d Fur) IsInteger() bool {
	return new(big.Int).Rem(d.i, precisionReuse).Sign() == 0
}

// format decimal state
func (d Fur) Format(s fmt.State, verb rune) {
	_, err := s.Write([]byte(d.String()))
	if err != nil {
		panic(err)
	}
}

func (d Fur) String() string {
	if d.i == nil {
		return d.i.String()
	}

	isNeg := d.IsNegative()

	if isNeg {
		d = d.Neg()
	}

	bzInt, err := d.i.MarshalText()
	if err != nil {
		return ""
	}
	inputSize := len(bzInt)

	var bzStr []byte

	// TODO: Remove trailing zeros
	// case 1, purely decimal
	if inputSize <= Precision {
		bzStr = make([]byte, Precision+2)

		// 0. prefix
		bzStr[0] = byte('0')
		bzStr[1] = byte('.')

		// set relevant digits to 0
		for i := 0; i < Precision-inputSize; i++ {
			bzStr[i+2] = byte('0')
		}

		// set final digits
		copy(bzStr[2+(Precision-inputSize):], bzInt)
	} else {
		// inputSize + 1 to account for the decimal point that is being added
		bzStr = make([]byte, inputSize+1)
		decPointPlace := inputSize - Precision

		copy(bzStr, bzInt[:decPointPlace])                   // pre-decimal digits
		bzStr[decPointPlace] = byte('.')                     // decimal point
		copy(bzStr[decPointPlace+1:], bzInt[decPointPlace:]) // post-decimal digits
	}

	if isNeg {
		return "-" + string(bzStr)
	}

	return string(bzStr)
}

// Float64 returns the float64 representation of a Fur.
// Will return the error if the conversion failed.
func (d Fur) Float64() (float64, error) {
	return strconv.ParseFloat(d.String(), 64)
}

// MustFloat64 returns the float64 representation of a Fur.
// Would panic if the conversion failed.
func (d Fur) MustFloat64() float64 {
	if value, err := strconv.ParseFloat(d.String(), 64); err != nil {
		panic(err)
	} else {
		return value
	}
}

//     ____
//  __|    |__   "chop 'em
//       ` \     round!"
// ___||  ~  _     -bankers
// |         |      __
// |       | |   __|__|__
// |_____:  /   | $$$    |
//              |________|

// Remove a Precision amount of rightmost digits and perform bankers rounding
// on the remainder (gaussian rounding) on the digits which have been removed.
//
// Mutates the input. Use the non-mutative version if that is undesired
func chopPrecisionAndRound(d *big.Int) *big.Int {
	// remove the negative and add it back when returning
	if d.Sign() == -1 {
		// make d positive, compute chopped value, and then un-mutate d
		d = d.Neg(d)
		d = chopPrecisionAndRound(d)
		d = d.Neg(d)
		return d
	}

	// get the truncated quotient and remainder
	quo, rem := d, big.NewInt(0)
	quo, rem = quo.QuoRem(d, precisionReuse, rem)

	if rem.Sign() == 0 { // remainder is zero
		return quo
	}

	switch rem.Cmp(fivePrecision) {
	case -1:
		return quo
	case 1:
		return quo.Add(quo, oneInt)
	default: // bankers rounding must take place
		// always round to an even number
		if quo.Bit(0) == 0 {
			return quo
		}
		return quo.Add(quo, oneInt)
	}
}

func chopPrecisionAndRoundUp(d *big.Int) *big.Int {
	// remove the negative and add it back when returning
	if d.Sign() == -1 {
		// make d positive, compute chopped value, and then un-mutate d
		d = d.Neg(d)
		// truncate since d is negative...
		d = chopPrecisionAndTruncate(d)
		d = d.Neg(d)
		return d
	}

	// get the truncated quotient and remainder
	quo, rem := d, big.NewInt(0)
	quo, rem = quo.QuoRem(d, precisionReuse, rem)

	if rem.Sign() == 0 { // remainder is zero
		return quo
	}

	return quo.Add(quo, oneInt)
}

func chopPrecisionAndRoundNonMutative(d *big.Int) *big.Int {
	tmp := new(big.Int).Set(d)
	return chopPrecisionAndRound(tmp)
}

// RoundInt64 rounds the decimal using bankers rounding
func (d Fur) RoundInt64() int64 {
	chopped := chopPrecisionAndRoundNonMutative(d.i)
	if !chopped.IsInt64() {
		panic("Int64() out of bound")
	}
	return chopped.Int64()
}

// RoundInt round the decimal using bankers rounding
func (d Fur) RoundInt() Int {
	return NewIntFromBigInt(chopPrecisionAndRoundNonMutative(d.i))
}

// chopPrecisionAndTruncate is similar to chopPrecisionAndRound,
// but always rounds down. It does not mutate the input.
func chopPrecisionAndTruncate(d *big.Int) *big.Int {
	return new(big.Int).Quo(d, precisionReuse)
}

// TruncateInt64 truncates the decimals from the number and returns an int64
func (d Fur) TruncateInt64() int64 {
	chopped := chopPrecisionAndTruncate(d.i)
	if !chopped.IsInt64() {
		panic("Int64() out of bound")
	}
	return chopped.Int64()
}

// TruncateInt truncates the decimals from the number and returns an Int
func (d Fur) TruncateInt() Int {
	return NewIntFromBigInt(chopPrecisionAndTruncate(d.i))
}

// TruncateFur truncates the decimals from the number and returns a Fur
func (d Fur) TruncateFur() Fur {
	return NewFurFromBigInt(chopPrecisionAndTruncate(d.i))
}

// Ceil returns the smallest interger value (as a decimal) that is greater than
// or equal to the given decimal.
func (d Fur) Ceil() Fur {
	tmp := new(big.Int).Set(d.i)

	quo, rem := tmp, big.NewInt(0)
	quo, rem = quo.QuoRem(tmp, precisionReuse, rem)

	// no need to round with a zero remainder regardless of sign
	if rem.Cmp(zeroInt) == 0 {
		return NewFurFromBigInt(quo)
	}

	if rem.Sign() == -1 {
		return NewFurFromBigInt(quo)
	}

	return NewFurFromBigInt(quo.Add(quo, oneInt))
}

// MaxSortableFur is the largest Fur that can be passed into SortableFurBytes()
// Its negative form is the least Fur that can be passed in.
var MaxSortableFur = OneFur().Quo(SmallestFur())

// ValidSortableFur ensures that a Fur is within the sortable bounds,
// a Fur can't have a precision of less than 10^-18.
// Max sortable decimal was set to the reciprocal of SmallestFur.
func ValidSortableFur(fur Fur) bool {
	return fur.Abs().LTE(MaxSortableFur)
}

// SortableFurBytes returns a byte slice representation of a Fur that can be sorted.
// Left and right pads with 0s so there are 18 digits to left and right of the decimal point.
// For this reason, there is a maximum and minimum value for this, enforced by ValidSortableFur.
func SortableFurBytes(fur Fur) []byte {
	if !ValidSortableFur(fur) {
		panic("fur must be within bounds")
	}
	// Instead of adding an extra byte to all sortable decs in order to handle max sortable, we just
	// makes its bytes be "max" which comes after all numbers in ASCIIbetical order
	if fur.Equal(MaxSortableFur) {
		return []byte("max")
	}
	// For the same reason, we make the bytes of minimum sortable fur be --, which comes before all numbers.
	if fur.Equal(MaxSortableFur.Neg()) {
		return []byte("--")
	}
	// We move the negative sign to the front of all the left padded 0s, to make negative numbers come before positive numbers
	if fur.IsNegative() {
		return append([]byte("-"), []byte(fmt.Sprintf(fmt.Sprintf("%%0%ds", Precision*2+1), fur.Abs().String()))...)
	}
	return []byte(fmt.Sprintf(fmt.Sprintf("%%0%ds", Precision*2+1), fur.String()))
}

// reuse nil values
var nilJSON []byte

func init() {
	empty := new(big.Int)
	bz, _ := empty.MarshalText()
	nilJSON, _ = json.Marshal(string(bz))
}

// MarshalJSON marshals the decimal
func (d Fur) MarshalJSON() ([]byte, error) {
	if d.i == nil {
		return nilJSON, nil
	}
	return json.Marshal(d.String())
}

// UnmarshalJSON defines custom decoding scheme
func (d *Fur) UnmarshalJSON(bz []byte) error {
	if d.i == nil {
		d.i = new(big.Int)
	}

	var text string
	err := json.Unmarshal(bz, &text)
	if err != nil {
		return err
	}

	// TODO: Reuse fur allocation
	newFur, err := NewFurFromStr(text)
	if err != nil {
		return err
	}

	d.i = newFur.i
	return nil
}

// MarshalYAML returns the YAML representation.
func (d Fur) MarshalYAML() (interface{}, error) {
	return d.String(), nil
}

// Marshal implements the gogo proto custom type interface.
func (d Fur) Marshal() ([]byte, error) {
	if d.i == nil {
		d.i = new(big.Int)
	}
	return d.i.MarshalText()
}

// MarshalTo implements the gogo proto custom type interface.
func (d *Fur) MarshalTo(data []byte) (n int, err error) {
	if d.i == nil {
		d.i = new(big.Int)
	}

	if d.i.Cmp(zeroInt) == 0 {
		copy(data, []byte{0x30})
		return 1, nil
	}

	bz, err := d.Marshal()
	if err != nil {
		return 0, err
	}

	copy(data, bz)
	return len(bz), nil
}

// Unmarshal implements the gogo proto custom type interface.
func (d *Fur) Unmarshal(data []byte) error {
	if len(data) == 0 {
		d = nil
		return nil
	}

	if d.i == nil {
		d.i = new(big.Int)
	}

	if err := d.i.UnmarshalText(data); err != nil {
		return err
	}

	if d.i.BitLen() > maxFurBitLen {
		return fmt.Errorf("decimal out of range; got: %d, max: %d", d.i.BitLen(), maxFurBitLen)
	}

	return nil
}

// Size implements the gogo proto custom type interface.
func (d *Fur) Size() int {
	bz, _ := d.Marshal()
	return len(bz)
}

// Override Amino binary serialization by proxying to protobuf.
func (d Fur) MarshalAmino() ([]byte, error)   { return d.Marshal() }
func (d *Fur) UnmarshalAmino(bz []byte) error { return d.Unmarshal(bz) }

func (dp FurProto) String() string {
	return dp.Fur.String()
}

// helpers

// test if two decimal arrays are equal
func DecsEqual(d1s, d2s []Fur) bool {
	if len(d1s) != len(d2s) {
		return false
	}

	for i, d1 := range d1s {
		if !d1.Equal(d2s[i]) {
			return false
		}
	}
	return true
}

// minimum decimal between two
func MinFur(d1, d2 Fur) Fur {
	if d1.LT(d2) {
		return d1
	}
	return d2
}

// maximum decimal between two
func MaxDec(d1, d2 Fur) Fur {
	if d1.LT(d2) {
		return d2
	}
	return d1
}

// intended to be used with require/assert:  require.True(DecEq(...))
func DecEq(t *testing.T, exp, got Fur) (*testing.T, bool, string, string, string) {
	return t, exp.Equal(got), "expected:\t%v\ngot:\t\t%v", exp.String(), got.String()
}
