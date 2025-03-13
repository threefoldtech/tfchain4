package types

const (
	// ModuleName defines the module name
	ModuleName = "tfchain4"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_tfchain4"
)

var (
	ParamsKey = []byte("p_tfchain4")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
