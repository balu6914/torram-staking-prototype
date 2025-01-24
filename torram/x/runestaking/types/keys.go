package types

const (
	// ModuleName defines the module name
	ModuleName = "runestaking"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_runestaking"
)

var (
	ParamsKey = []byte("p_runestaking")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
