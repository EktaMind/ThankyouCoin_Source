package vecmt2dagidx

import (
	"github.com/EktaMind/Thank_Sirus/abft"
	"github.com/EktaMind/Thank_Sirus/abft/dagidx"
	"github.com/EktaMind/Thank_Sirus/hash"
	"github.com/EktaMind/Thank_Sirus/inter/idx"
	"github.com/EktaMind/Thank_Sirus/vecfc"

	"github.com/kalibroida/ThankyouCoin_Node/vecmt"
)

type Adapter struct {
	*vecmt.Index
}

var _ abft.DagIndex = (*Adapter)(nil)

type AdapterSeq struct {
	*vecmt.HighestBefore
}

type BranchSeq struct {
	vecfc.BranchSeq
}

// Seq is a maximum observed e.Seq in the branch
func (b *BranchSeq) Seq() idx.Event {
	return b.BranchSeq.Seq
}

// MinSeq is a minimum observed e.Seq in the branch
func (b *BranchSeq) MinSeq() idx.Event {
	return b.BranchSeq.MinSeq
}

// Size of the vector clock
func (b AdapterSeq) Size() int {
	return b.VSeq.Size()
}

// Get i's position in the byte-encoded vector clock
func (b AdapterSeq) Get(i idx.Validator) dagidx.Seq {
	seq := b.HighestBefore.VSeq.Get(i)
	return &BranchSeq{seq}
}

func (v *Adapter) GetMergedHighestBefore(id hash.Event) dagidx.HighestBeforeSeq {
	return AdapterSeq{v.Index.GetMergedHighestBefore(id)}
}

func Wrap(v *vecmt.Index) *Adapter {
	return &Adapter{v}
}
