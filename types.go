package tpocket

type (
	// PresetID - resource preset id, should be unique
	// for a same user, the same resource preset id can only have one ft
	PresetID = int64

	// NftID - nft id, should be unique
	// for and a same user and same preset id, there can be multiple nft ids
	NftID = int64
)
