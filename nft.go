package tpocket

import (
	"fmt"

	"github.com/khgame/memstore"
)

type (
	NFT struct {
		ID              NftID                      `json:"id"`
		PID             PresetID                   `json:"pid"`
		ContractStorage map[string]ContractRuntime `json:"contracts,omitempty"`
		Status          int64                      `json:"status,omitempty"`
	}
)

// StoreName - store name of NFT is nft:{nftID}
// presetID is not included in store name, because the nftID is unique
// in the whole system, then the nft can be retrieved by nftID
func (nft NFT) StoreName() string {
	return fmt.Sprintf("nft:%d", nft.ID)
}

// SealNFT - seal a nft with preset id
func SealNFT(id NftID) NFT {
	return NFT{
		ID:              id,
		ContractStorage: map[string]ContractRuntime{},
	}
}

var _ memstore.StorableType = NFT{}
