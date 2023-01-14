package types

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

// ----------------------------------------------------------------------------
// Decimal Coin

// NewFurCoin creates a new FurCoin instance from an Int.
func NewFurCoin(denom string, amount Int) FurCoin {
	coin := NewCoin(denom, amount)

	return FurCoin{
		Denom:  coin.Denom,
		Amount: coin.Amount.ToFur(),
	}
}

// NewFurCoinFromFur creates a new FurCoin instance from a Fur.
func NewFurCoinFromFur(denom string, amount Fur) FurCoin {
	mustValidateDenom(denom)

	if amount.IsNegative() {
		panic(fmt.Sprintf("negative decimal coin amount: %v\n", amount))
	}

	return FurCoin{
		Denom:  denom,
		Amount: amount,
	}
}

// NewFurCoinFromCoin creates a new FurCoin from a Coin.
func NewFurCoinFromCoin(coin Coin) FurCoin {
	if err := coin.Validate(); err != nil {
		panic(err)
	}

	return FurCoin{
		Denom:  coin.Denom,
		Amount: coin.Amount.ToFur(),
	}
}

// NewInt64FurCoin returns a new FurCoin with a denomination and amount. It will
// panic if the amount is negative or denom is invalid.
func NewInt64FurCoin(denom string, amount int64) FurCoin {
	return NewFurCoin(denom, NewInt(amount))
}

// IsZero returns if the FurCoin amount is zero.
func (coin FurCoin) IsZero() bool {
	return coin.Amount.IsZero()
}

// IsGTE returns true if they are the same type and the receiver is
// an equal or greater value.
func (coin FurCoin) IsGTE(other FurCoin) bool {
	if coin.Denom != other.Denom {
		panic(fmt.Sprintf("invalid coin denominations; %s, %s", coin.Denom, other.Denom))
	}

	return !coin.Amount.LT(other.Amount)
}

// IsLT returns true if they are the same type and the receiver is
// a smaller value.
func (coin FurCoin) IsLT(other FurCoin) bool {
	if coin.Denom != other.Denom {
		panic(fmt.Sprintf("invalid coin denominations; %s, %s", coin.Denom, other.Denom))
	}

	return coin.Amount.LT(other.Amount)
}

// IsEqual returns true if the two sets of Coins have the same value.
func (coin FurCoin) IsEqual(other FurCoin) bool {
	if coin.Denom != other.Denom {
		panic(fmt.Sprintf("invalid coin denominations; %s, %s", coin.Denom, other.Denom))
	}

	return coin.Amount.Equal(other.Amount)
}

// Add adds amounts of two decimal coins with same denom.
func (coin FurCoin) Add(coinB FurCoin) FurCoin {
	if coin.Denom != coinB.Denom {
		panic(fmt.Sprintf("coin denom different: %v %v\n", coin.Denom, coinB.Denom))
	}
	return FurCoin{coin.Denom, coin.Amount.Add(coinB.Amount)}
}

// Sub subtracts amounts of two decimal coins with same denom.
func (coin FurCoin) Sub(coinB FurCoin) FurCoin {
	if coin.Denom != coinB.Denom {
		panic(fmt.Sprintf("coin denom different: %v %v\n", coin.Denom, coinB.Denom))
	}
	res := FurCoin{coin.Denom, coin.Amount.Sub(coinB.Amount)}
	if res.IsNegative() {
		panic("negative decimal coin amount")
	}
	return res
}

// TruncateFurimal returns a Coin with a truncated decimal and a FurCoin for the
// change. Note, the change may be zero.
func (coin FurCoin) TruncateFurimal() (Coin, FurCoin) {
	truncated := coin.Amount.TruncateInt()
	change := coin.Amount.Sub(truncated.ToFur())
	return NewCoin(coin.Denom, truncated), NewFurCoinFromFur(coin.Denom, change)
}

// IsPositive returns true if coin amount is positive.
//
// TODO: Remove once unsigned integers are used.
func (coin FurCoin) IsPositive() bool {
	return coin.Amount.IsPositive()
}

// IsNegative returns true if the coin amount is negative and false otherwise.
//
// TODO: Remove once unsigned integers are used.
func (coin FurCoin) IsNegative() bool {
	return coin.Amount.IsNegative()
}

// String implements the Stringer interface for FurCoin. It returns a
// human-readable representation of a decimal coin.
func (coin FurCoin) String() string {
	return fmt.Sprintf("%v%v", coin.Amount, coin.Denom)
}

// Validate returns an error if the FurCoin has a negative amount or if the denom is invalid.
func (coin FurCoin) Validate() error {
	if err := ValidateDenom(coin.Denom); err != nil {
		return err
	}
	if coin.IsNegative() {
		return fmt.Errorf("decimal coin %s amount cannot be negative", coin)
	}
	return nil
}

// IsValid returns true if the FurCoin has a non-negative amount and the denom is valid.
func (coin FurCoin) IsValid() bool {
	return coin.Validate() == nil
}

// ----------------------------------------------------------------------------
// Decimal Coins

// FurCoins defines a slice of coins with decimal values
type FurCoins []FurCoin

// NewFurCoins constructs a new coin set with with decimal values
// from FurCoins. The provided coins will be sanitized by removing
// zero coins and sorting the coin set. A panic will occur if the coin set is not valid.
func NewFurCoins(decCoins ...FurCoin) FurCoins {
	newFurCoins := sanitizeFurCoins(decCoins)
	if err := newFurCoins.Validate(); err != nil {
		panic(fmt.Errorf("invalid coin set %s: %w", newFurCoins, err))
	}

	return newFurCoins
}

func sanitizeFurCoins(decCoins []FurCoin) FurCoins {
	// remove zeroes
	newFurCoins := removeZeroFurCoins(decCoins)
	if len(newFurCoins) == 0 {
		return FurCoins{}
	}

	return newFurCoins.Sort()
}

// NewFurCoinsFromCoins constructs a new coin set with decimal values
// from regular Coins.
func NewFurCoinsFromCoins(coins ...Coin) FurCoins {
	decCoins := make(FurCoins, len(coins))
	newCoins := NewCoins(coins...)
	for i, coin := range newCoins {
		decCoins[i] = NewFurCoinFromCoin(coin)
	}

	return decCoins
}

// String implements the Stringer interface for FurCoins. It returns a
// human-readable representation of decimal coins.
func (coins FurCoins) String() string {
	if len(coins) == 0 {
		return ""
	}

	out := ""
	for _, coin := range coins {
		out += fmt.Sprintf("%v,", coin.String())
	}

	return out[:len(out)-1]
}

// TruncateFurimal returns the coins with truncated decimals and returns the
// change. Note, it will not return any zero-amount coins in either the truncated or
// change coins.
func (coins FurCoins) TruncateFurimal() (truncatedCoins Coins, changeCoins FurCoins) {
	for _, coin := range coins {
		truncated, change := coin.TruncateFurimal()
		if !truncated.IsZero() {
			truncatedCoins = truncatedCoins.Add(truncated)
		}
		if !change.IsZero() {
			changeCoins = changeCoins.Add(change)
		}
	}

	return truncatedCoins, changeCoins
}

// Add adds two sets of FurCoins.
//
// NOTE: Add operates under the invariant that coins are sorted by
// denominations.
//
// CONTRACT: Add will never return Coins where one Coin has a non-positive
// amount. In otherwords, IsValid will always return true.
func (coins FurCoins) Add(coinsB ...FurCoin) FurCoins {
	return coins.safeAdd(coinsB)
}

// safeAdd will perform addition of two FurCoins sets. If both coin sets are
// empty, then an empty set is returned. If only a single set is empty, the
// other set is returned. Otherwise, the coins are compared in order of their
// denomination and addition only occurs when the denominations match, otherwise
// the coin is simply added to the sum assuming it's not zero.
func (coins FurCoins) safeAdd(coinsB FurCoins) FurCoins {
	sum := ([]FurCoin)(nil)
	indexA, indexB := 0, 0
	lenA, lenB := len(coins), len(coinsB)

	for {
		if indexA == lenA {
			if indexB == lenB {
				// return nil coins if both sets are empty
				return sum
			}

			// return set B (excluding zero coins) if set A is empty
			return append(sum, removeZeroFurCoins(coinsB[indexB:])...)
		} else if indexB == lenB {
			// return set A (excluding zero coins) if set B is empty
			return append(sum, removeZeroFurCoins(coins[indexA:])...)
		}

		coinA, coinB := coins[indexA], coinsB[indexB]

		switch strings.Compare(coinA.Denom, coinB.Denom) {
		case -1: // coin A denom < coin B denom
			if !coinA.IsZero() {
				sum = append(sum, coinA)
			}

			indexA++

		case 0: // coin A denom == coin B denom
			res := coinA.Add(coinB)
			if !res.IsZero() {
				sum = append(sum, res)
			}

			indexA++
			indexB++

		case 1: // coin A denom > coin B denom
			if !coinB.IsZero() {
				sum = append(sum, coinB)
			}

			indexB++
		}
	}
}

// negative returns a set of coins with all amount negative.
func (coins FurCoins) negative() FurCoins {
	res := make([]FurCoin, 0, len(coins))
	for _, coin := range coins {
		res = append(res, FurCoin{
			Denom:  coin.Denom,
			Amount: coin.Amount.Neg(),
		})
	}
	return res
}

// Sub subtracts a set of FurCoins from another (adds the inverse).
func (coins FurCoins) Sub(coinsB FurCoins) FurCoins {
	diff, hasNeg := coins.SafeSub(coinsB)
	if hasNeg {
		panic("negative coin amount")
	}

	return diff
}

// SafeSub performs the same arithmetic as Sub but returns a boolean if any
// negative coin amount was returned.
func (coins FurCoins) SafeSub(coinsB FurCoins) (FurCoins, bool) {
	diff := coins.safeAdd(coinsB.negative())
	return diff, diff.IsAnyNegative()
}

// Intersect will return a new set of coins which contains the minimum FurCoin
// for common denoms found in both `coins` and `coinsB`. For denoms not common
// to both `coins` and `coinsB` the minimum is considered to be 0, thus they
// are not added to the final set. In other words, trim any denom amount from
// coin which exceeds that of coinB, such that (coin.Intersect(coinB)).IsLTE(coinB).
// See also Coins.Min().
func (coins FurCoins) Intersect(coinsB FurCoins) FurCoins {
	res := make([]FurCoin, len(coins))
	for i, coin := range coins {
		minCoin := FurCoin{
			Denom:  coin.Denom,
			Amount: MinFur(coin.Amount, coinsB.AmountOf(coin.Denom)),
		}
		res[i] = minCoin
	}
	return removeZeroFurCoins(res)
}

// GetDenomByIndex returns the Denom to make the findDup generic
func (coins FurCoins) GetDenomByIndex(i int) string {
	return coins[i].Denom
}

// IsAnyNegative returns true if there is at least one coin whose amount
// is negative; returns false otherwise. It returns false if the FurCoins set
// is empty too.
//
// TODO: Remove once unsigned integers are used.
func (coins FurCoins) IsAnyNegative() bool {
	for _, coin := range coins {
		if coin.IsNegative() {
			return true
		}
	}

	return false
}

// MulFur multiplies all the coins by a decimal.
//
// CONTRACT: No zero coins will be returned.
func (coins FurCoins) MulFur(d Fur) FurCoins {
	var res FurCoins
	for _, coin := range coins {
		product := FurCoin{
			Denom:  coin.Denom,
			Amount: coin.Amount.Mul(d),
		}

		if !product.IsZero() {
			res = res.Add(product)
		}
	}

	return res
}

// MulFurTruncate multiplies all the decimal coins by a decimal, truncating. It
// panics if d is zero.
//
// CONTRACT: No zero coins will be returned.
func (coins FurCoins) MulFurTruncate(d Fur) FurCoins {
	var res FurCoins

	for _, coin := range coins {
		product := FurCoin{
			Denom:  coin.Denom,
			Amount: coin.Amount.MulTruncate(d),
		}

		if !product.IsZero() {
			res = res.Add(product)
		}
	}

	return res
}

// QuoFur divides all the decimal coins by a decimal. It panics if d is zero.
//
// CONTRACT: No zero coins will be returned.
func (coins FurCoins) QuoFur(d Fur) FurCoins {
	if d.IsZero() {
		panic("invalid zero decimal")
	}

	var res FurCoins
	for _, coin := range coins {
		quotient := FurCoin{
			Denom:  coin.Denom,
			Amount: coin.Amount.Quo(d),
		}

		if !quotient.IsZero() {
			res = res.Add(quotient)
		}
	}

	return res
}

// QuoFurTruncate divides all the decimal coins by a decimal, truncating. It
// panics if d is zero.
//
// CONTRACT: No zero coins will be returned.
func (coins FurCoins) QuoFurTruncate(d Fur) FurCoins {
	if d.IsZero() {
		panic("invalid zero decimal")
	}

	var res FurCoins
	for _, coin := range coins {
		quotient := FurCoin{
			Denom:  coin.Denom,
			Amount: coin.Amount.QuoTruncate(d),
		}

		if !quotient.IsZero() {
			res = res.Add(quotient)
		}
	}

	return res
}

// Empty returns true if there are no coins and false otherwise.
func (coins FurCoins) Empty() bool {
	return len(coins) == 0
}

// AmountOf returns the amount of a denom from deccoins
func (coins FurCoins) AmountOf(denom string) Fur {
	mustValidateDenom(denom)

	switch len(coins) {
	case 0:
		return ZeroFur()

	case 1:
		coin := coins[0]
		if coin.Denom == denom {
			return coin.Amount
		}
		return ZeroFur()

	default:
		midIdx := len(coins) / 2 // 2:1, 3:1, 4:2
		coin := coins[midIdx]

		switch {
		case denom < coin.Denom:
			return coins[:midIdx].AmountOf(denom)
		case denom == coin.Denom:
			return coin.Amount
		default:
			return coins[midIdx+1:].AmountOf(denom)
		}
	}
}

// IsEqual returns true if the two sets of FurCoins have the same value.
func (coins FurCoins) IsEqual(coinsB FurCoins) bool {
	if len(coins) != len(coinsB) {
		return false
	}

	coins = coins.Sort()
	coinsB = coinsB.Sort()

	for i := 0; i < len(coins); i++ {
		if !coins[i].IsEqual(coinsB[i]) {
			return false
		}
	}

	return true
}

// IsZero returns whether all coins are zero
func (coins FurCoins) IsZero() bool {
	for _, coin := range coins {
		if !coin.Amount.IsZero() {
			return false
		}
	}
	return true
}

// Validate checks that the FurCoins are sorted, have positive amount, with a valid and unique
// denomination (i.e no duplicates). Otherwise, it returns an error.
func (coins FurCoins) Validate() error {
	switch len(coins) {
	case 0:
		return nil

	case 1:
		if err := ValidateDenom(coins[0].Denom); err != nil {
			return err
		}
		if !coins[0].IsPositive() {
			return fmt.Errorf("coin %s amount is not positive", coins[0])
		}
		return nil
	default:
		// check single coin case
		if err := (FurCoins{coins[0]}).Validate(); err != nil {
			return err
		}

		lowDenom := coins[0].Denom
		seenDenoms := make(map[string]bool)
		seenDenoms[lowDenom] = true

		for _, coin := range coins[1:] {
			if seenDenoms[coin.Denom] {
				return fmt.Errorf("duplicate denomination %s", coin.Denom)
			}
			if err := ValidateDenom(coin.Denom); err != nil {
				return err
			}
			if coin.Denom <= lowDenom {
				return fmt.Errorf("denomination %s is not sorted", coin.Denom)
			}
			if !coin.IsPositive() {
				return fmt.Errorf("coin %s amount is not positive", coin.Denom)
			}

			// we compare each coin against the last denom
			lowDenom = coin.Denom
			seenDenoms[coin.Denom] = true
		}

		return nil
	}
}

// IsValid calls Validate and returns true when the FurCoins are sorted, have positive amount, with a
// valid and unique denomination (i.e no duplicates).
func (coins FurCoins) IsValid() bool {
	return coins.Validate() == nil
}

// IsAllPositive returns true if there is at least one coin and all currencies
// have a positive value.
//
// TODO: Remove once unsigned integers are used.
func (coins FurCoins) IsAllPositive() bool {
	if len(coins) == 0 {
		return false
	}

	for _, coin := range coins {
		if !coin.IsPositive() {
			return false
		}
	}

	return true
}

func removeZeroFurCoins(coins FurCoins) FurCoins {
	result := make([]FurCoin, 0, len(coins))

	for _, coin := range coins {
		if !coin.IsZero() {
			result = append(result, coin)
		}
	}

	return result
}

//-----------------------------------------------------------------------------
// Sorting

var _ sort.Interface = FurCoins{}

// Len implements sort.Interface for FurCoins
func (coins FurCoins) Len() int { return len(coins) }

// Less implements sort.Interface for FurCoins
func (coins FurCoins) Less(i, j int) bool { return coins[i].Denom < coins[j].Denom }

// Swap implements sort.Interface for FurCoins
func (coins FurCoins) Swap(i, j int) { coins[i], coins[j] = coins[j], coins[i] }

// Sort is a helper function to sort the set of decimal coins in-place.
func (coins FurCoins) Sort() FurCoins {
	sort.Sort(coins)
	return coins
}

// ----------------------------------------------------------------------------
// Parsing

// ParseFurCoin parses a decimal coin from a string, returning an error if
// invalid. An empty string is considered invalid.
func ParseFurCoin(coinStr string) (coin FurCoin, err error) {
	coinStr = strings.TrimSpace(coinStr)

	matches := reFurCoin.FindStringSubmatch(coinStr)
	if matches == nil {
		return FurCoin{}, fmt.Errorf("invalid decimal coin expression: %s", coinStr)
	}

	amountStr, denomStr := matches[1], matches[2]

	amount, err := NewFurFromStr(amountStr)
	if err != nil {
		return FurCoin{}, errors.Wrap(err, fmt.Sprintf("failed to parse decimal coin amount: %s", amountStr))
	}

	if err := ValidateDenom(denomStr); err != nil {
		return FurCoin{}, fmt.Errorf("invalid denom cannot contain upper case characters or spaces: %s", err)
	}

	return NewFurCoinFromFur(denomStr, amount), nil
}

// ParseFurCoins will parse out a list of decimal coins separated by commas. If the parsing is successuful,
// the provided coins will be sanitized by removing zero coins and sorting the coin set. Lastly
// a validation of the coin set is executed. If the check passes, ParseFurCoins will return the sanitized coins.
// Otherwise it will return an error.
// If an empty string is provided to ParseFurCoins, it returns nil Coins.
// Expected format: "{amount0}{denomination},...,{amountN}{denominationN}"
func ParseFurCoins(coinsStr string) (FurCoins, error) {
	coinsStr = strings.TrimSpace(coinsStr)
	if len(coinsStr) == 0 {
		return nil, nil
	}

	coinStrs := strings.Split(coinsStr, ",")
	decCoins := make(FurCoins, len(coinStrs))
	for i, coinStr := range coinStrs {
		coin, err := ParseFurCoin(coinStr)
		if err != nil {
			return nil, err
		}

		decCoins[i] = coin
	}

	newFurCoins := sanitizeFurCoins(decCoins)
	if err := newFurCoins.Validate(); err != nil {
		return nil, err
	}

	return newFurCoins, nil
}
