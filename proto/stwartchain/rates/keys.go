package rates

const (
	// ModuleName defines the module name
	ModuleName = "rates"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_rates"
)

var (
	ParamsKey = []byte("p_rates")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	TokensKey      = "Rates/value/"
	TokensCountKey = "Rates/count/"
)

// RatesKey returns the store key to retrieve a Rates from the index fields
func RatesKey(
	denom string,
) []byte {
	var key []byte

	denomBytes := []byte(denom)
	key = append(key, denomBytes...)
	key = append(key, []byte("/")...)

	return key
}
