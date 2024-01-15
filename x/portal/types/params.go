package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyProviderChannel = []byte("ProviderChannel")
	// TODO: Determine the default value
	DefaultProviderChannel string = "provider_channel"
)

var (
	KeyReserveChannel = []byte("ReserveChannel")
	// TODO: Determine the default value
	DefaultReserveChannel string = "reserve_channel"
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	providerChannel string,
	reserveChannel string,
) Params {
	return Params{
		ProviderChannel: providerChannel,
		ReserveChannel:  reserveChannel,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultProviderChannel,
		DefaultReserveChannel,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyProviderChannel, &p.ProviderChannel, validateProviderChannel),
		paramtypes.NewParamSetPair(KeyReserveChannel, &p.ReserveChannel, validateReserveChannel),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateProviderChannel(p.ProviderChannel); err != nil {
		return err
	}

	if err := validateReserveChannel(p.ReserveChannel); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateProviderChannel validates the ProviderChannel param
func validateProviderChannel(v interface{}) error {
	providerChannel, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = providerChannel

	return nil
}

// validateReserveChannel validates the ReserveChannel param
func validateReserveChannel(v interface{}) error {
	reserveChannel, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = reserveChannel

	return nil
}
