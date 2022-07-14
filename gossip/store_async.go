package gossip

import (
	"github.com/EktaMind/Thank_Sirus/kvdb"
	"github.com/EktaMind/Thank_Sirus/kvdb/table"
)

type asyncStore struct {
	mainDB kvdb.Store
	table  struct {
		// Network tables
		Peers kvdb.Store `table:"Z"`
	}
}

func newAsyncStore(db kvdb.Store) *asyncStore {
	s := &asyncStore{
		mainDB: db,
	}

	table.MigrateTables(&s.table, s.mainDB)

	return s
}

// Close leaves underlying database.
func (s *asyncStore) Close() {
	table.MigrateTables(&s.table, nil)

	_ = s.mainDB.Close()
}
