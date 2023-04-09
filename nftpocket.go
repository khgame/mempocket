package tpocket

import (
	"context"
	"time"

	"github.com/khgame/memstore"
)

// NFTPocket - a pocket for nft
type NFTPocket struct {
	Pocket[NFT]

	// systemFTPocket !!! is a system pocket, to manager contract of preset itself
	systemFTPocket *FTPocket
}

// Create - create a nft,
// - pid is the preset id of the nft
// - nftID is the nft id, should be unique in the whole system
//
// at the beginning, the preset id must be set, to make sure the nft is created by the right preset
// after the nft is created, a nft can be addressed by nftID
//
// to avoid unexpected nft creation, the Set method is not exposed
// todo: consider remove method (witch is not implemented)
func (p *NFTPocket) Create(ctx context.Context, user string, nftID NftID, pID PresetID) (NFT, error) {
	nft := SealNFT(nftID)
	nft.PID = pID
	if err := p.storage.Set(user, &nft); err != nil {
		return NFT{}, err
	}
	// update the index in ft pocket
	//
	// when save
	// - index: permanentKey => users
	// - data: [permanentKey + user] => data
	//
	// (ft_pocket [permanentKey : system] -->
	//   (storage: [user] -->
	//		(data: [pid] --> (ft.contract ["nft"] --> (ContractRuntime:[nftID] --> [timestamp])))
	//   )
	// )
	if err := p.systemFTPocket.DoContract(ctx, user, pID, "nft", func(runtime *ContractRuntime) (*ContractRuntime, error) {
		indexTableVal, ok := (*runtime)["index"]
		var indexTable map[NftID]int64
		if ok {
			indexTable = indexTableVal.(map[NftID]int64)
		} else {
			indexTable = make(map[NftID]int64)
		}
		indexTable[nftID] = time.Now().Unix()
		(*runtime)["index"] = indexTable
		return runtime, nil
	}); err != nil {
		// if set ft count failed, delete the nft
		if err = p.storage.Delete(user, nft.StoreName()); err != nil {
			return NFT{}, err
		}
		return NFT{}, err
	}
	return nft, nil
}

// Get - get ft from pocket
func (p *NFTPocket) Get(ctx context.Context, user string, nftID NftID) (nft NFT, err error) {
	nft = SealNFT(nftID)
	if err = p.storage.Get(user, &nft); err != nil {
		return nft, err
	}
	return nft, nil
}

// ListByPID - get ft from pocket
func (p *NFTPocket) ListByPID(ctx context.Context, user string, pid PresetID) (nfts []NFT, err error) {
	var ft FT
	if ft, err = p.systemFTPocket.Get(ctx, user, pid); err != nil {
		return nil, err
	}
	if ft.Contracts == nil {
		return nil, nil
	}
	nftContract, ok := ft.Contracts["nft"]
	if !ok {
		return nil, nil
	}
	indexTable, ok := nftContract["index"]
	if !ok {
		return nil, nil
	}
	for nftID := range indexTable.(map[NftID]int64) {
		var nft NFT
		if nft, err = p.Get(ctx, user, nftID); err != nil {
			return nil, err
		}
		nfts = append(nfts, nft)
	}
	return nfts, nil
}

// Update - update nft contract storage
func (p *NFTPocket) Update(ctx context.Context, user string, nftID NftID, fnUpdate func(nft *NFT) (*NFT, error)) error {
	return p.storage.Update(user, SealNFT(nftID).StoreName(), fnUpdate)
}

func MakeNFTPocket(ctx context.Context, appID string, metaStr string,
	storage memstore.Storage[NFT], systemFTPocket *FTPocket,
) NFTPocket {
	return NFTPocket{
		Pocket: Pocket[NFT]{
			AppID:   appID,
			Meta:    metaStr,
			storage: storage,
		},
		systemFTPocket: systemFTPocket,
	}
}
