package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var (
	// KeyEarnRate is byte key for EarnRate param.
	KeyOnomyChannel = []byte("OnomyChannel") //nolint:gochecknoglobals // cosmos-sdk style
	// KeyReserveChannel is byte key for ReserveChannel param.
	KeyReserveChannel = []byte("ReserveChannel") //nolint:gochecknoglobals // cosmos-sdk style
)

var (
	// DefaultEarnRate is default value for the DefaultEarnRate param.
	DefaultOnomyChannel = "1" //nolint:gomnd,gochecknoglobals // cosmos-sdk style
	// DefaultReserveChannel is default value for the DefaultReserveChannel param.
	DefaultReserveChannel = "2" //nolint:gomnd,gochecknoglobals // cosmos-sdk style
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	onomyChannel string,
	reserveChannel string,
) Params {
	return Params{
		OnomyChannel:   onomyChannel,
		ReserveChannel: reserveChannel,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultOnomyChannel, DefaultReserveChannel)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyOnomyChannel, &p.OnomyChannel, validateOnomyChannel),
		paramtypes.NewParamSetPair(KeyReserveChannel, &p.ReserveChannel, validateReserveChannel),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateOnomyChannel(p.OnomyChannel); err != nil {
		return err
	}
	if err := validateReserveChannel(p.ReserveChannel); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p) //nolint:errcheck // error is not expected here
	return string(out)
}

func validateOnomyChannel(i interface{}) error {
	value, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	onomyChannel, ok := sdk.NewIntFromString(value)
	if !ok {
		return fmt.Errorf("invalid string number format: %q", value)
	}
	if onomyChannel.LTE(sdk.ZeroInt()) {
		return fmt.Errorf("onomy channel must be positive and greater than zero: %d", onomyChannel)
	}

	return nil
}

func validateReserveChannel(i interface{}) error {
	value, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	reserveChannel, ok := sdk.NewIntFromString(value)
	if !ok {
		return fmt.Errorf("invalid string number format: %q", value)
	}
	if reserveChannel.LTE(sdk.ZeroInt()) {
		return fmt.Errorf("reserve channel must be positive and greater than zero: %d", reserveChannel)
	}

	return nil
}
