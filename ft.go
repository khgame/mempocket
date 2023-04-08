package tpocket

import (
	"fmt"
)

type FT struct {
	// resource preset id (unique)
	PID PresetID `json:"pid"`
	// quantity
	Quantity int64 `json:"q"`
}

func (ft FT) StoreName() string {
	return fmt.Sprintf("ft:%d", ft.PID)
}

// SealFT - seal a ft with preset id
func SealFT[T ~int64 | ~int32 | ~int16 | ~int8 | ~int | ~uint32 | ~uint16 | ~uint8 | ~uint](pid T) FT {
	return FT{
		PID: PresetID(pid),
	}
}
