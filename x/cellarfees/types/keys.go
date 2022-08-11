package types

const (
	// ModuleName is the module name constant used in many places
	ModuleName = "cellarfees"

	// StoreKey is the store key string for cellarfees
	StoreKey = ModuleName

	// RouterKey is the message route for cellarfees
	RouterKey = ModuleName

	// QuerierRoute is the querier route for cellarfees
	QuerierRoute = ModuleName
)

var (
	CellarFeePoolKey = []byte{0x00} // key for global cellar fee pool state
)
