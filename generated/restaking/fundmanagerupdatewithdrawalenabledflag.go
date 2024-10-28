// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// FundManagerUpdateWithdrawalEnabledFlag is the `fund_manager_update_withdrawal_enabled_flag` instruction.
type FundManagerUpdateWithdrawalEnabledFlag struct {
	Enabled *bool

	// [0] = [SIGNER] fund_manager
	//
	// [1] = [] receipt_token_mint
	//
	// [2] = [WRITE] fund_account
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewFundManagerUpdateWithdrawalEnabledFlagInstructionBuilder creates a new `FundManagerUpdateWithdrawalEnabledFlag` instruction builder.
func NewFundManagerUpdateWithdrawalEnabledFlagInstructionBuilder() *FundManagerUpdateWithdrawalEnabledFlag {
	nd := &FundManagerUpdateWithdrawalEnabledFlag{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	nd.AccountMetaSlice[0] = ag_solanago.Meta(Addresses["5UpLTLA7Wjqp7qdfjuTtPcUw3aVtbqFA5Mgm34mxPNg2"]).SIGNER()
	nd.AccountMetaSlice[1] = ag_solanago.Meta(Addresses["FRAGSEthVFL7fdqM8hxfxkfCZzUvmg21cqPJVvC1qdbo"])
	return nd
}

// SetEnabled sets the "enabled" parameter.
func (inst *FundManagerUpdateWithdrawalEnabledFlag) SetEnabled(enabled bool) *FundManagerUpdateWithdrawalEnabledFlag {
	inst.Enabled = &enabled
	return inst
}

// SetFundManagerAccount sets the "fund_manager" account.
func (inst *FundManagerUpdateWithdrawalEnabledFlag) SetFundManagerAccount(fundManager ag_solanago.PublicKey) *FundManagerUpdateWithdrawalEnabledFlag {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(fundManager).SIGNER()
	return inst
}

// GetFundManagerAccount gets the "fund_manager" account.
func (inst *FundManagerUpdateWithdrawalEnabledFlag) GetFundManagerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *FundManagerUpdateWithdrawalEnabledFlag) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *FundManagerUpdateWithdrawalEnabledFlag {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(receiptTokenMint)
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *FundManagerUpdateWithdrawalEnabledFlag) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetFundAccountAccount sets the "fund_account" account.
func (inst *FundManagerUpdateWithdrawalEnabledFlag) SetFundAccountAccount(fundAccount ag_solanago.PublicKey) *FundManagerUpdateWithdrawalEnabledFlag {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(fundAccount).WRITE()
	return inst
}

func (inst *FundManagerUpdateWithdrawalEnabledFlag) findFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: fund
	seeds = append(seeds, []byte{byte(0x66), byte(0x75), byte(0x6e), byte(0x64)})
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindFundAccountAddressWithBumpSeed calculates FundAccount account address with given seeds and a known bump seed.
func (inst *FundManagerUpdateWithdrawalEnabledFlag) FindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *FundManagerUpdateWithdrawalEnabledFlag) MustFindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundAccountAddress finds FundAccount account address with given seeds.
func (inst *FundManagerUpdateWithdrawalEnabledFlag) FindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *FundManagerUpdateWithdrawalEnabledFlag) MustFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundAccountAccount gets the "fund_account" account.
func (inst *FundManagerUpdateWithdrawalEnabledFlag) GetFundAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst FundManagerUpdateWithdrawalEnabledFlag) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_FundManagerUpdateWithdrawalEnabledFlag,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst FundManagerUpdateWithdrawalEnabledFlag) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *FundManagerUpdateWithdrawalEnabledFlag) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Enabled == nil {
			return errors.New("Enabled parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.FundManager is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ReceiptTokenMint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.FundAccount is not set")
		}
	}
	return nil
}

func (inst *FundManagerUpdateWithdrawalEnabledFlag) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("FundManagerUpdateWithdrawalEnabledFlag")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Enabled", *inst.Enabled))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      fund_manager", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("receipt_token_mint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("             fund_", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj FundManagerUpdateWithdrawalEnabledFlag) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Enabled` param:
	err = encoder.Encode(obj.Enabled)
	if err != nil {
		return err
	}
	return nil
}
func (obj *FundManagerUpdateWithdrawalEnabledFlag) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Enabled`:
	err = decoder.Decode(&obj.Enabled)
	if err != nil {
		return err
	}
	return nil
}

// NewFundManagerUpdateWithdrawalEnabledFlagInstruction declares a new FundManagerUpdateWithdrawalEnabledFlag instruction with the provided parameters and accounts.
func NewFundManagerUpdateWithdrawalEnabledFlagInstruction(
	// Parameters:
	enabled bool,
	// Accounts:
	fundManager ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	fundAccount ag_solanago.PublicKey) *FundManagerUpdateWithdrawalEnabledFlag {
	return NewFundManagerUpdateWithdrawalEnabledFlagInstructionBuilder().
		SetEnabled(enabled).
		SetFundManagerAccount(fundManager).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetFundAccountAccount(fundAccount)
}