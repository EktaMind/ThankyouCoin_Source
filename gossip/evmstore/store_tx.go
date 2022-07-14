package evmstore

import (
	"github.com/EktaMind/Thank_ethereum/common"
	"github.com/EktaMind/Thank_ethereum/core/types"
)

// SetTx stores non-event transaction.
func (s *Store) SetTx(txid common.Hash, tx *types.Transaction) {
	s.rlp.Set(s.table.Txs, txid.Bytes(), tx)
}

// GetTx returns stored non-event transaction.
func (s *Store) GetTx(txid common.Hash) *types.Transaction {
	tx, _ := s.rlp.Get(s.table.Txs, txid.Bytes(), &types.Transaction{}).(*types.Transaction)

	return tx
}
