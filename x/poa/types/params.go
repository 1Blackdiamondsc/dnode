// Parameters store for PoA module.
package types

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/params"
)

// Default parameters.
const (
	DefaultMaxValidators uint16 = 11
	DefaultMinValidators uint16 = 3
)

// Key to store min and max validators parameters.
var (
	KeyMaxValidators = []byte("MaxValidators")
	KeyMinValidators = []byte("MinValidators")
)

// Describing parameters for PoA module, like: min and max validators amount.
type Params struct {
	MaxValidators uint16 `json:"max_validators"`
	MinValidators uint16 `json:"min_validators"`
}

// Create new instance to store parameters.
func NewParams(maxValidators, minValidators uint16) Params {
	return Params{
		MaxValidators: maxValidators,
		MinValidators: minValidators,
	}
}

func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{KeyMaxValidators, &p.MaxValidators},
		{KeyMinValidators, &p.MinValidators},
	}
}

func (p *Params) Equal(p2 Params) bool {
	return p.MinValidators == p2.MinValidators &&
		p.MaxValidators == p2.MaxValidators
}

func (p Params) Validate() error {
	if p.MinValidators < DefaultMinValidators {
		return fmt.Errorf("minimum amount of validators should be not less %d", DefaultMinValidators)
	}

	if p.MaxValidators > DefaultMaxValidators {
		return fmt.Errorf("maximum amount of validators should be not great then %d", DefaultMaxValidators)
	}

	return nil
}

func (p Params) String() string {
	return fmt.Sprintf("\tMax Validators: %d\tMin Validators: %d", p.MaxValidators, p.MaxValidators)
}

func DefaultParams() Params {
	return NewParams(DefaultMaxValidators, DefaultMinValidators)
}
