package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var (
	// KeyEarnRate is byte key for EarnRate param.
	KeyEarnRate = []byte("EarnRate") //nolint:gochecknoglobals // cosmos-sdk style
	// KeyBurnRate is byte key for BurnRate param.
	KeyBurnRate = []byte("BurnRate") //nolint:gochecknoglobals // cosmos-sdk style
)

var (
	// DefaultEarnRate is default value for the DefaultEarnRate param.
	DefaultEarnRate = []sdk.Int{sdk.NewInt(1), sdk.NewInt(10)} //nolint:gomnd,gochecknoglobals // cosmos-sdk style
	// DefaultBurnRate is default value for the DefaultBurnRate param.
	DefaultBurnRate = []sdk.Int{sdk.NewInt(1), sdk.NewInt(10)} //nolint:gomnd,gochecknoglobals // cosmos-sdk style
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	earnRate []sdk.Int,
	burnRate []sdk.Int,
) Params {
	return Params{
		EarnRate: earnRate,
		BurnRate: burnRate,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultEarnRate, DefaultBurnRate)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyEarnRate, &p.EarnRate, validateEarnRate),
		paramtypes.NewParamSetPair(KeyBurnRate, &p.BurnRate, validateBurnRate),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateEarnRate(p.EarnRate); err != nil {
		return err
	}
	if err := validateBurnRate(p.BurnRate); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p) //nolint:errcheck // error is not expected here
	return string(out)
}

func validateEarnRate(i interface{}) error {
	v, ok := i.([]sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if len(v) != 2 {
		return fmt.Errorf("earn rate must have 2 elements: %d", v)
	}

	if v[0].LTE(sdk.NewInt(0)) {
		return fmt.Errorf("earn rate numerator must be positive and greater than zero: %d", v)
	}

	if v[1].LTE(v[0]) {
		return fmt.Errorf("earn rate denominator must be greater than numerator: %d", v)
	}

	return nil
}

func validateBurnRate(i interface{}) error {
	v, ok := i.([]sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if len(v) != 2 {
		return fmt.Errorf("burn rate must have 2 elements: %d", v)
	}

	if v[0].LTE(sdk.NewInt(0)) {
		return fmt.Errorf("burn rate numerator must be positive and greater than zero: %d", v)
	}

	if v[1].LTE(v[0]) {
		return fmt.Errorf("burn rate denominator must be greater than numerator: %d", v)
	}

	return nil
}
