package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var (
	// KeyEarnRate is byte key for EarnRate param.
	KeyEarnRates = []byte("EarnRates") //nolint:gochecknoglobals // cosmos-sdk style
	// KeyBurnRate is byte key for BurnRate param.
	KeyBurnRate = []byte("BurnRate") //nolint:gochecknoglobals // cosmos-sdk style
	// KeyBurnCoin is byte key for BurnCoin param.
	KeyBurnCoin = []byte("BurnCoin") //nolint:gochecknoglobals // cosmos-sdk style
	// KeyBurnCoin is byte key for BurnCoin param.
	KeyMarketFee = []byte("MarketFee") //nolint:gochecknoglobals // cosmos-sdk style
)

var (
	// DefaultEarnRate is default value for the DefaultEarnRate param.
	DefaultEarnRates = "0500,0300,0200" //nolint:gomnd,gochecknoglobals // cosmos-sdk style
	// DefaultBurnRate is default value for the DefaultBurnRate param.
	DefaultBurnRate = "1000" //nolint:gomnd,gochecknoglobals // cosmos-sdk style
	// DefaultBurnCoin is default value for the DefaultBurnCoin param.
	DefaultBurnCoin = "stake" //nolint:gomnd,gochecknoglobals // cosmos-sdk style
	// DefaultMarketFee is default value for the MarketFee param.
	DefaultMarketFee = "0030" //nolint:gomnd,gochecknoglobals // cosmos-sdk style

)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	earnRates string,
	burnRate string,
	burnCoin string,
	marketFee string,
) Params {
	return Params{
		EarnRates: earnRates,
		BurnRate:  burnRate,
		BurnCoin:  burnCoin,
		MarketFee: marketFee,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultEarnRates, DefaultBurnRate, DefaultBurnCoin, DefaultMarketFee)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyEarnRates, &p.EarnRates, validateEarnRates),
		paramtypes.NewParamSetPair(KeyBurnRate, &p.BurnRate, validateBurnRate),
		paramtypes.NewParamSetPair(KeyBurnCoin, &p.BurnCoin, validateBurnCoin),
		paramtypes.NewParamSetPair(KeyMarketFee, &p.MarketFee, validateMarketFee),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateEarnRates(p.EarnRates); err != nil {
		return err
	}
	if err := validateBurnRate(p.BurnRate); err != nil {
		return err
	}
	if err := validateBurnCoin(p.BurnCoin); err != nil {
		return err
	}
	if err := validateMarketFee(p.MarketFee); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p) //nolint:errcheck // error is not expected here
	return string(out)
}

func validateEarnRates(i interface{}) error {
	value, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	earnRatesStringArray := strings.Split(value, ",")

	if len(earnRatesStringArray) > 10 {
		return fmt.Errorf("the maximum number of rate values is 10")
	}

	var earnRates [10]sdk.Int
	for i, v := range earnRatesStringArray {
		earnRates[i], ok = sdk.NewIntFromString(v)
		if !ok {
			return fmt.Errorf("invalid string number format: %q", v)
		}
		if earnRates[i].LTE(sdk.ZeroInt()) {
			return fmt.Errorf("earn rate numerator must be positive and greater than zero: %d", earnRates[i])
		}
		if earnRates[i].GTE(sdk.NewInt(10000)) {
			return fmt.Errorf("earn rate numerator must be less than 10000: %d", earnRates[i])
		}
		if i > 0 {
			if earnRates[i].GT(earnRates[i-1]) {
				return fmt.Errorf("earn rates must not increase")
			}
		}
	}

	return nil
}

func validateBurnRate(i interface{}) error {
	value, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	burnRate, ok := sdk.NewIntFromString(value)
	if !ok {
		return fmt.Errorf("invalid string number format: %q", value)
	}
	if burnRate.LTE(sdk.ZeroInt()) {
		return fmt.Errorf("burn rate numerator must be positive and greater than zero: %d", burnRate)
	}
	if burnRate.GTE(sdk.NewInt(10000)) {
		return fmt.Errorf("burn rate numerator must be less than 10000: %d", burnRate)
	}

	return nil
}

func validateBurnCoin(i interface{}) error {
	_, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateMarketFee(i interface{}) error {
	value, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	marketFee, ok := sdk.NewIntFromString(value)
	if !ok {
		return fmt.Errorf("invalid string number format: %q", value)
	}
	if marketFee.LTE(sdk.ZeroInt()) {
		return fmt.Errorf("market fee numerator must be positive and greater than zero: %d", marketFee)
	}
	if marketFee.GTE(sdk.NewInt(10000)) {
		return fmt.Errorf("market fee numerator must be less than 10000: %d", marketFee)
	}

	return nil
}
