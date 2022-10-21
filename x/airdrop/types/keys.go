package types

const (
	// ModuleName defines the module name
	ModuleName = "airdrop"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_airdrop"

	// Keep track of the index of claimed
	ClaimedKey      = "Claimed/value/"
	ClaimedCountKey = "Claimed/count/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
