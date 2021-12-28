package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}

func (s Sequence) ValidateBasic() error {
	if len(s.IDKey) == 0 {
		return sdkerrors.Wrap(ErrEmpty, "id key")
	}
	return nil
}

func (s GenesisState) ValidateBasic() error {
	if err := s.Params.ValidateBasic(); err != nil {
		return sdkerrors.Wrap(err, "params")
	}
	for i := range s.Codes {
		if err := s.Codes[i].ValidateBasic(); err != nil {
			return sdkerrors.Wrapf(err, "code: %d", i)
		}
	}
	for i := range s.Contracts {
		if err := s.Contracts[i].ValidateBasic(); err != nil {
			return sdkerrors.Wrapf(err, "contract: %d", i)
		}
	}
	for i := range s.Sequences {
		if err := s.Sequences[i].ValidateBasic(); err != nil {
			return sdkerrors.Wrapf(err, "sequence: %d", i)
		}
	}
	for i := range s.GenMsgs {
		if err := s.GenMsgs[i].ValidateBasic(); err != nil {
			return sdkerrors.Wrapf(err, "gen message: %d", i)
		}
	}
	return nil
}

func (c Code) ValidateBasic() error {
	if c.CodeID == 0 {
		return sdkerrors.Wrap(ErrEmpty, "code id")
	}
	if err := c.CodeInfo.ValidateBasic(); err != nil {
		return sdkerrors.Wrap(err, "code info")
	}
	if err := validateWasmCode(c.CodeBytes); err != nil {
		return sdkerrors.Wrap(err, "code bytes")
	}
	return nil
}

func (c Contract) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(c.ContractAddress); err != nil {
		return sdkerrors.Wrap(err, "contract address")
	}
	if err := c.ContractInfo.ValidateBasic(); err != nil {
		return sdkerrors.Wrap(err, "contract info")
	}

	if c.ContractInfo.Created != nil {
		return sdkerrors.Wrap(ErrInvalid, "created must be empty")
	}
	for i := range c.ContractState {
		if err := c.ContractState[i].ValidateBasic(); err != nil {
			return sdkerrors.Wrapf(err, "contract state %d", i)
		}
	}
	return nil
}

// AsMsg returns the underlying cosmos-sdk message instance. Null when can not be mapped to a known type.
func (m GenesisState_GenMsgs) AsMsg() sdk.Msg {
	if msg := m.GetStoreCode(); msg != nil {
		return msg
	}
	if msg := m.GetInstantiateContract(); msg != nil {
		return msg
	}
	if msg := m.GetExecuteContract(); msg != nil {
		return msg
	}
	return nil
}

func (m GenesisState_GenMsgs) ValidateBasic() error {
	msg := m.AsMsg()
	if msg == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "unknown message")
	}
	return msg.ValidateBasic()
}

// ValidateGenesis performs basic validation of supply genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error {
	return data.ValidateBasic()
}
