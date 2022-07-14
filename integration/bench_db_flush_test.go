package integration

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/EktaMind/Thank_Sirus/abft"
	"github.com/EktaMind/Thank_Sirus/hash"
	"github.com/EktaMind/Thank_Sirus/inter/idx"
	"github.com/EktaMind/Thank_Sirus/kvdb"
	"github.com/EktaMind/Thank_Sirus/kvdb/leveldb"
	"github.com/EktaMind/Thank_Sirus/utils/cachescale"
	"github.com/EktaMind/Thank_ethereum/common"
	"github.com/syndtr/goleveldb/leveldb/opt"

	"github.com/kalibroida/ThankyouCoin_Node/gossip"
	"github.com/kalibroida/ThankyouCoin_Node/integration/makegenesis"
	"github.com/kalibroida/ThankyouCoin_Node/inter"
	"github.com/kalibroida/ThankyouCoin_Node/thankyoucoin/genesisstore"
	"github.com/kalibroida/ThankyouCoin_Node/utils"
	"github.com/kalibroida/ThankyouCoin_Node/vecmt"
)

func BenchmarkFlushDBs(b *testing.B) {
	rawProducer, dir := dbProducer("flush_bench")
	defer os.RemoveAll(dir)
	genStore := makegenesis.FakeGenesisStore(1, utils.ToThx(1), utils.ToThx(1))
	_, _, store, s2, s3, _ := MakeEngine(rawProducer, InputGenesis{
		Hash: genStore.Hash(),
		Read: func(store *genesisstore.Store) error {
			buf := bytes.NewBuffer(nil)
			err := genStore.Export(buf)
			if err != nil {
				return err
			}
			return store.Import(buf)
		},
		Close: func() error {
			return nil
		},
	}, Configs{
		Photon:      gossip.DefaultConfig(cachescale.Identity),
		PhotonStore: gossip.DefaultStoreConfig(cachescale.Identity),
		Sirius:      abft.DefaultConfig(),
		SiriusStore: abft.DefaultStoreConfig(cachescale.Identity),
		VectorClock: vecmt.DefaultConfig(cachescale.Identity),
	})
	defer store.Close()
	defer s2.Close()
	defer s3.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := idx.Block(0)
		randUint32s := func() []uint32 {
			arr := make([]uint32, 128)
			for i := 0; i < len(arr); i++ {
				arr[i] = uint32(i) ^ (uint32(n) << 16) ^ 0xd0ad884e
			}
			return []uint32{uint32(n), uint32(n) + 1, uint32(n) + 2}
		}
		for !store.IsCommitNeeded(false) {
			store.SetBlock(n, &inter.Block{
				Time:        inter.Timestamp(n << 32),
				Atropos:     hash.Event{},
				Events:      hash.Events{},
				Txs:         []common.Hash{},
				InternalTxs: []common.Hash{},
				SkippedTxs:  randUint32s(),
				GasUsed:     uint64(n) << 24,
				Root:        hash.Hash{},
			})
			n++
		}
		err := store.Commit()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func cache64mb(string) int {
	return 64 * opt.MiB
}

func dbProducer(name string) (kvdb.IterableDBProducer, string) {
	dir, err := ioutil.TempDir("", name)
	if err != nil {
		panic(err)
	}
	return leveldb.NewProducer(dir, cache64mb), dir
}
