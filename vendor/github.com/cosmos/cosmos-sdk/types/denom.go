package types

import (
	"fmt"
)

// denomUnits contains a mapping of denomination mapped to their respective unit
// multipliers (e.g. 1atom = 10^-6uatom).
var denomUnits = map[string]Fur{}

// baseDenom is the denom of smallest unit registered
var baseDenom string

// RegisterDenom registers a denomination with a corresponding unit. If the
// denomination is already registered, an error will be returned.
func RegisterDenom(denom string, unit Fur) error {
	if err := ValidateDenom(denom); err != nil {
		return err
	}

	if _, ok := denomUnits[denom]; ok {
		return fmt.Errorf("denom %s already registered", denom)
	}

	denomUnits[denom] = unit

	if baseDenom == "" || unit.LT(denomUnits[baseDenom]) {
		baseDenom = denom
	}
	return nil
}

// GetDenomUnit returns a unit for a given denomination if it exists. A boolean
// is returned if the denomination is registered.
func GetDenomUnit(denom string) (Fur, bool) {
	if err := ValidateDenom(denom); err != nil {
		return ZeroFur(), false
	}

	unit, ok := denomUnits[denom]
	if !ok {
		return ZeroFur(), false
	}

	return unit, true
}

// GetBaseDenom returns the denom of smallest unit registered
func GetBaseDenom() (string, error) {
	if baseDenom == "" {
		return "", fmt.Errorf("no denom is registered")
	}
	return baseDenom, nil
}

// ConvertCoin attempts to convert a coin to a given denomination. If the given
// denomination is invalid or if neither denomination is registered, an error
// is returned.
func ConvertCoin(coin Coin, denom string) (Coin, error) {
	if err := ValidateDenom(denom); err != nil {
		return Coin{}, err
	}

	srcUnit, ok := GetDenomUnit(coin.Denom)
	if !ok {
		return Coin{}, fmt.Errorf("source denom not registered: %s", coin.Denom)
	}

	dstUnit, ok := GetDenomUnit(denom)
	if !ok {
		return Coin{}, fmt.Errorf("destination denom not registered: %s", denom)
	}

	if srcUnit.Equal(dstUnit) {
		return NewCoin(denom, coin.Amount), nil
	}

	return NewCoin(denom, coin.Amount.ToFur().Mul(srcUnit).Quo(dstUnit).TruncateInt()), nil
}

// ConvertFurCoin attempts to convert a decimal coin to a given denomination. If the given
// denomination is invalid or if neither denomination is registered, an error
// is returned.
func ConvertFurCoin(coin FurCoin, denom string) (FurCoin, error) {
	if err := ValidateDenom(denom); err != nil {
		return FurCoin{}, err
	}

	srcUnit, ok := GetDenomUnit(coin.Denom)
	if !ok {
		return FurCoin{}, fmt.Errorf("source denom not registered: %s", coin.Denom)
	}

	dstUnit, ok := GetDenomUnit(denom)
	if !ok {
		return FurCoin{}, fmt.Errorf("destination denom not registered: %s", denom)
	}

	if srcUnit.Equal(dstUnit) {
		return NewFurCoinFromFur(denom, coin.Amount), nil
	}

	return NewFurCoinFromFur(denom, coin.Amount.Mul(srcUnit).Quo(dstUnit)), nil
}

// NormalizeCoin try to convert a coin to the smallest unit registered,
// returns original one if failed.
func NormalizeCoin(coin Coin) Coin {
	base, err := GetBaseDenom()
	if err != nil {
		return coin
	}
	newCoin, err := ConvertCoin(coin, base)
	if err != nil {
		return coin
	}
	return newCoin
}

// NormalizeFurCoin try to convert a decimal coin to the smallest unit registered,
// returns original one if failed.
func NormalizeFurCoin(coin FurCoin) FurCoin {
	base, err := GetBaseDenom()
	if err != nil {
		return coin
	}
	newCoin, err := ConvertFurCoin(coin, base)
	if err != nil {
		return coin
	}
	return newCoin
}

// NormalizeCoins normalize and truncate a list of decimal coins
func NormalizeCoins(coins []FurCoin) Coins {
	if coins == nil {
		return nil
	}
	result := make([]Coin, 0, len(coins))

	for _, coin := range coins {
		newCoin, _ := NormalizeFurCoin(coin).TruncateFurimal()
		result = append(result, newCoin)
	}

	return result
}
