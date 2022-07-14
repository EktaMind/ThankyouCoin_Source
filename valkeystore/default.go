package valkeystore

import (
	"github.com/EktaMind/Thank_ethereum/accounts/keystore"

	"github.com/kalibroida/ThankyouCoin_Node/valkeystore/encryption"
)

func NewDefaultFileRawKeystore(dir string) *FileKeystore {
	enc := encryption.New(keystore.StandardScryptN, keystore.StandardScryptP)
	return NewFileKeystore(dir, enc)
}

func NewDefaultMemKeystore() *SyncedKeystore {
	return NewSyncedKeystore(NewCachedKeystore(NewMemKeystore()))
}

func NewDefaultFileKeystore(dir string) *SyncedKeystore {
	return NewSyncedKeystore(NewCachedKeystore(NewDefaultFileRawKeystore(dir)))
}
