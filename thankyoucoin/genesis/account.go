package genesis

import (
	"math/big"

	"github.com/EktaMind/Thank_Sirus/hash"
	"github.com/EktaMind/Thank_Sirus/inter/idx"
	"github.com/EktaMind/Thank_Sirus/kvdb"
	"github.com/EktaMind/Thank_ethereum/common"
	"github.com/EktaMind/Thank_ethereum/core/types"

	"github.com/kalibroida/ThankyouCoin_Node/inter"
)

type (
	// Accounts specifies the changes to EVM accounts after applying RawEvmItems.
	Accounts interface {
		ForEach(fn func(common.Address, Account))
	}
	Storage interface {
		ForEach(fn func(common.Address, common.Hash, common.Hash))
	}
	Delegations interface {
		ForEach(fn func(common.Address, idx.ValidatorID, Delegation))
	}
	Blocks interface {
		ForEach(fn func(idx.Block, Block))
	}
	RawEvmItems kvdb.Iteratee

	Delegation struct {
		Stake              *big.Int
		Rewards            *big.Int
		LockedStake        *big.Int
		LockupFromEpoch    idx.Epoch
		LockupEndTime      inter.Timestamp
		LockupDuration     inter.Timestamp
		EarlyUnlockPenalty *big.Int
	}
	// Account is an account in the state of the genesis block.
	Account struct {
		Code         []byte
		Balance      *big.Int
		Nonce        uint64
		SelfDestruct bool
	}

	Block struct {
		Time        inter.Timestamp
		Atropos     hash.Event
		Txs         types.Transactions
		InternalTxs types.Transactions
		Root        hash.Hash
		Receipts    []*types.ReceiptForStorage
	}
)
