package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "carpchaintestnet"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_carpchaintestnet"
)

var (
	ParamsKey = collections.NewPrefix("p_carpchaintestnet")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
