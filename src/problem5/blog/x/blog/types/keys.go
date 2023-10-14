package types

const (
	// ModuleName defines the module name
	ModuleName = "blog"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_blog"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	// Prefix used to uniquely identify a post in system
	PostKey = "Post/value/"

	// used to keep track of the ID of the latest post added to the store
	PostCountKey = "Post/count/"
)
