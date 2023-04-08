package tpocket

import (
	"github.com/khgame/memstore"
)

// FTPocket : app_id:pocket_name
// user -->|pid| { pid, quantity } ...
type FTPocket struct {

	// AppID - the app id of this pocket, should be unique
	// generally, it's an application id be assigned by the platform
	// e.g. "com.khgame.001"
	AppID string `json:"app_id"`

	// pocket name - the usage of this pocket, should be unique in an app
	// e.g. "resource", "items", "barn", "exp", "coin"
	PocketName string `json:"name"`

	// storage - the ft pocket embed storage and provide high level api
	// to operate it. the implementation of memstore.Storage should be injected
	// by the caller.
	storage memstore.Storage[FT]
}

func MakeTPocket(appID string, name string, storage memstore.Storage[FT]) FTPocket {
	return FTPocket{
		AppID:      appID,
		PocketName: name,
		storage:    storage,
	}
}

// Get - get ft from pocket
func (p *FTPocket) Get(user string, pid PresetID) (ft FT, err error) {
	ft = SealFT(pid)
	if err = p.storage.Get(user, &ft); err != nil {
		return ft, err
	}
	return ft, nil
}

// Set - set ft to pocket
func (p *FTPocket) Set(user string, ft FT) error {
	return p.storage.Set(user, &ft)
}

// Incr - incr ft quantity
func (p *FTPocket) Incr(user string, amount FT) error {
	return p.storage.Update(user, amount.StoreName(), func(ft *FT) (*FT, error) {
		if ft == nil {
			return &amount, nil
		}
		ft.Quantity += amount.Quantity
		return ft, nil
	})
}
