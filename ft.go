package tpocket

import (
	"fmt"

	"github.com/khgame/memstore"
)

type FT struct {
	// resource preset id (unique)
	PID PresetID `json:"pid"`
	// quantity
	Quantity int64 `json:"q"`
	// contracts
	Contracts map[string]ContractRuntime `json:"c,omitempty"`
	// memo
	Memo string `json:"memo,omitempty"`
}

func (ft FT) StoreName() string {
	return fmt.Sprintf("ft:%d", ft.PID)
}

// DecodeFTStoreName - decode id from store name
func DecodeFTStoreName[T PIDLike](storeName string) (T, error) {
	var id int64
	// scan the store name to get the id
	if _, err := fmt.Sscanf(storeName, "ft:%d", &id); err != nil {
		return 0, err
	}
	return T(id), nil
}

// SealFT - seal a ft with preset id
func SealFT[T PIDLike](pid T) FT {
	return FT{
		PID: PresetID(pid),
	}
}

var _ memstore.StorableType = FT{}
