// Copyright 2021 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const DepositTxType = 0x7E

type DepositTx struct {
	// SourceHash uniquely identifies the source of the deposit
	SourceHash common.Hash
	// From is exposed through the types.Signer, not through TxData
	From common.Address
	// nil means contract creation
	To *common.Address `rlp:"nil"`
	// Mint is minted on L2, locked on L1, nil if no minting.
	Mint *big.Int `rlp:"nil"`
	// Value is transferred from L2 balance, executed after Mint (if any)
	Value *big.Int
	// Guaranteed Gas is paid for on L1. It is always available and non-refundable.
	GuaranteedGas uint64
	// Additional Gas is bought and paid for on L2. It may not be provided if
	// there is not enough remaining gas or the price is too low.
	AdditionalGas      uint64
	AdditionalGasPrice *big.Int
	Data               []byte
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *DepositTx) copy() TxData {
	cpy := &DepositTx{
		SourceHash:         tx.SourceHash,
		From:               tx.From,
		To:                 copyAddressPtr(tx.To),
		Mint:               nil,
		Value:              new(big.Int),
		GuaranteedGas:      tx.GuaranteedGas,
		AdditionalGas:      tx.AdditionalGas,
		AdditionalGasPrice: new(big.Int),
		Data:               common.CopyBytes(tx.Data),
	}
	if tx.Mint != nil {
		cpy.Mint = new(big.Int).Set(tx.Mint)
	}
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	if tx.AdditionalGasPrice != nil {
		cpy.AdditionalGasPrice.Set(tx.AdditionalGasPrice)
	}
	return cpy
}

// DepositsNonce identifies a deposit, since go-ethereum abstracts all transaction types to a core.Message.
// Deposits do not set a nonce, deposits are included by the system and cannot be repeated or included elsewhere.
// Note: This is one less than the maximum value because of a go-etheruem test that fails otherwise.
const DepositsNonce uint64 = 0xffff_ffff_ffff_fffd

// accessors for innerTx.
func (tx *DepositTx) txType() byte           { return DepositTxType }
func (tx *DepositTx) chainID() *big.Int      { panic("deposits are not signed and do not have a chain-ID") }
func (tx *DepositTx) protected() bool        { return true }
func (tx *DepositTx) accessList() AccessList { return nil }
func (tx *DepositTx) data() []byte           { return tx.Data }
func (tx *DepositTx) gas() uint64            { return tx.GuaranteedGas }
func (tx *DepositTx) gasFeeCap() *big.Int    { return new(big.Int) }
func (tx *DepositTx) gasTipCap() *big.Int    { return new(big.Int) }
func (tx *DepositTx) gasPrice() *big.Int     { return new(big.Int) }
func (tx *DepositTx) value() *big.Int        { return tx.Value }
func (tx *DepositTx) nonce() uint64          { return DepositsNonce }
func (tx *DepositTx) to() *common.Address    { return tx.To }

func (tx *DepositTx) rawSignatureValues() (v, r, s *big.Int) {
	panic("deposit tx does not have a signature")
}

func (tx *DepositTx) setSignatureValues(chainID, v, r, s *big.Int) {
	panic("deposit tx does not have a signature")
}
