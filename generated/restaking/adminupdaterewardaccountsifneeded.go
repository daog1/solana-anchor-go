// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// AdminUpdateRewardAccountsIfNeeded is the `admin_update_reward_accounts_if_needed` instruction.
type AdminUpdateRewardAccountsIfNeeded struct {
	DesiredAccountSize *uint32 `bin:"optional"`
	Initialize         *bool

	// [0] = [WRITE, SIGNER] payer
	//
	// [1] = [SIGNER] admin
	//
	// [2] = [] system_program
	//
	// [3] = [] receipt_token_mint
	//
	// [4] = [WRITE] reward_account
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAdminUpdateRewardAccountsIfNeededInstructionBuilder creates a new `AdminUpdateRewardAccountsIfNeeded` instruction builder.
func NewAdminUpdateRewardAccountsIfNeededInstructionBuilder() *AdminUpdateRewardAccountsIfNeeded {
	nd := &AdminUpdateRewardAccountsIfNeeded{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	nd.AccountMetaSlice[1] = ag_solanago.Meta(Addresses["fragkamrANLvuZYQPcmPsCATQAabkqNGH6gxqqPG3aP"]).SIGNER()
	nd.AccountMetaSlice[2] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	nd.AccountMetaSlice[3] = ag_solanago.Meta(Addresses["FRAGSEthVFL7fdqM8hxfxkfCZzUvmg21cqPJVvC1qdbo"])
	return nd
}

// SetDesiredAccountSize sets the "desired_account_size" parameter.
func (inst *AdminUpdateRewardAccountsIfNeeded) SetDesiredAccountSize(desired_account_size uint32) *AdminUpdateRewardAccountsIfNeeded {
	inst.DesiredAccountSize = &desired_account_size
	return inst
}

// SetInitialize sets the "initialize" parameter.
func (inst *AdminUpdateRewardAccountsIfNeeded) SetInitialize(initialize bool) *AdminUpdateRewardAccountsIfNeeded {
	inst.Initialize = &initialize
	return inst
}

// SetPayerAccount sets the "payer" account.
func (inst *AdminUpdateRewardAccountsIfNeeded) SetPayerAccount(payer ag_solanago.PublicKey) *AdminUpdateRewardAccountsIfNeeded {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *AdminUpdateRewardAccountsIfNeeded) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAdminAccount sets the "admin" account.
func (inst *AdminUpdateRewardAccountsIfNeeded) SetAdminAccount(admin ag_solanago.PublicKey) *AdminUpdateRewardAccountsIfNeeded {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *AdminUpdateRewardAccountsIfNeeded) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *AdminUpdateRewardAccountsIfNeeded) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *AdminUpdateRewardAccountsIfNeeded {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *AdminUpdateRewardAccountsIfNeeded) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *AdminUpdateRewardAccountsIfNeeded) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *AdminUpdateRewardAccountsIfNeeded {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(receiptTokenMint)
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *AdminUpdateRewardAccountsIfNeeded) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetRewardAccountAccount sets the "reward_account" account.
func (inst *AdminUpdateRewardAccountsIfNeeded) SetRewardAccountAccount(rewardAccount ag_solanago.PublicKey) *AdminUpdateRewardAccountsIfNeeded {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(rewardAccount).WRITE()
	return inst
}

func (inst *AdminUpdateRewardAccountsIfNeeded) findFindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: reward
	seeds = append(seeds, []byte{byte(0x72), byte(0x65), byte(0x77), byte(0x61), byte(0x72), byte(0x64)})
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

// FindRewardAccountAddressWithBumpSeed calculates RewardAccount account address with given seeds and a known bump seed.
func (inst *AdminUpdateRewardAccountsIfNeeded) FindRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindRewardAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *AdminUpdateRewardAccountsIfNeeded) MustFindRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindRewardAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindRewardAccountAddress finds RewardAccount account address with given seeds.
func (inst *AdminUpdateRewardAccountsIfNeeded) FindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindRewardAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *AdminUpdateRewardAccountsIfNeeded) MustFindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindRewardAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetRewardAccountAccount gets the "reward_account" account.
func (inst *AdminUpdateRewardAccountsIfNeeded) GetRewardAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst AdminUpdateRewardAccountsIfNeeded) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AdminUpdateRewardAccountsIfNeeded,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AdminUpdateRewardAccountsIfNeeded) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AdminUpdateRewardAccountsIfNeeded) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Initialize == nil {
			return errors.New("Initialize parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.ReceiptTokenMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.RewardAccount is not set")
		}
	}
	return nil
}

func (inst *AdminUpdateRewardAccountsIfNeeded) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AdminUpdateRewardAccountsIfNeeded")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("  DesiredAccountSize (OPT)", inst.DesiredAccountSize))
						paramsBranch.Child(ag_format.Param("          Initialize", *inst.Initialize))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("             payer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("             admin", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("    system_program", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("receipt_token_mint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("           reward_", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj AdminUpdateRewardAccountsIfNeeded) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `DesiredAccountSize` param (optional):
	{
		if obj.DesiredAccountSize == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.DesiredAccountSize)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Initialize` param:
	err = encoder.Encode(obj.Initialize)
	if err != nil {
		return err
	}
	return nil
}
func (obj *AdminUpdateRewardAccountsIfNeeded) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `DesiredAccountSize` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.DesiredAccountSize)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Initialize`:
	err = decoder.Decode(&obj.Initialize)
	if err != nil {
		return err
	}
	return nil
}

// NewAdminUpdateRewardAccountsIfNeededInstruction declares a new AdminUpdateRewardAccountsIfNeeded instruction with the provided parameters and accounts.
func NewAdminUpdateRewardAccountsIfNeededInstruction(
	// Parameters:
	desired_account_size uint32,
	initialize bool,
	// Accounts:
	payer ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	rewardAccount ag_solanago.PublicKey) *AdminUpdateRewardAccountsIfNeeded {
	return NewAdminUpdateRewardAccountsIfNeededInstructionBuilder().
		SetDesiredAccountSize(desired_account_size).
		SetInitialize(initialize).
		SetPayerAccount(payer).
		SetAdminAccount(admin).
		SetSystemProgramAccount(systemProgram).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetRewardAccountAccount(rewardAccount)
}