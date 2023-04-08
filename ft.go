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

// SealFT - seal a ft with preset id
func SealFT[T PIDLike](pid T) FT {
	return FT{
		PID: PresetID(pid),
	}
}

var _ memstore.StorableType = FT{}
