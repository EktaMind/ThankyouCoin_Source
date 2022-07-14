package thankyoucoin

import (
	"math/big"

	"github.com/EktaMind/Thank_Sirus/hash"
	"github.com/EktaMind/Thank_Sirus/inter/idx"
	"github.com/EktaMind/Thank_ethereum/common"

	"github.com/kalibroida/ThankyouCoin_Node/inter"
	"github.com/kalibroida/ThankyouCoin_Node/thankyoucoin/genesis"
	"github.com/kalibroida/ThankyouCoin_Node/thankyoucoin/genesis/gpos"
)

type Genesis struct {
	Accounts    genesis.Accounts
	Storage     genesis.Storage
	Delegations genesis.Delegations
	Blocks      genesis.Blocks
	RawEvmItems genesis.RawEvmItems
	Validators  gpos.Validators

	FirstEpoch    idx.Epoch
	PrevEpochTime inter.Timestamp
	Time          inter.Timestamp
	ExtraData     []byte

	TotalSupply *big.Int

	DriverOwner common.Address

	Rules Rules

	Hash func() hash.Hash
}
