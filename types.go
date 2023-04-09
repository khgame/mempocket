package tpocket

import "github.com/khgame/memstore"

type (
	// PresetID - resource preset id, should be unique
	// for a same user, the same resource preset id can only have one ft
	PresetID = int64

	// NftID - nft id, should be unique
	// for and a same user and same preset id, there can be multiple nft ids
	NftID = int64

	PIDLike interface {
		~int64 | ~int32 | ~int16 | ~int8 | ~int | ~uint32 | ~uint16 | ~uint8 | ~uint
	}

	ContractRuntime map[string]any
)

// Pocket : app_id:pocket_name
type Pocket[T memstore.StorableType] struct {
	// AppID - the app id of this pocket, should be unique
	// generally, it's an application id be assigned by the platform
	// e.g. "com.khgame.001"
	AppID string `json:"app_id"`

	// Meta - the meta info of this pocket, it's customized by the caller
	// in practice, it can be a name of the pocket, or record the persist
	// key of the related storage
	// e.g. "resource", "items", "barn", "exp", "coin"
	Meta string `json:"name"`

	// storage - the ft pocket embed storage and provide high level api
	// to operate it. the implementation of memstore.Storage should be injected
	// by the caller.
	storage memstore.Storage[T]
}
