package tpocket

import (
	"context"
	"github.com/khgame/memstore"
)

// FTPocket : app_id:pocket_name
// user -->|pid| { pid, quantity } ...
type FTPocket Pocket[FT]

func MakeFTPocket(ctx context.Context, appID string, name string, storage memstore.Storage[FT]) FTPocket {
	return FTPocket{
		AppID:      appID,
		PocketName: name,
		storage:    storage,
	}
}

// Get - get ft from pocket
func (p *FTPocket) Get(ctx context.Context, user string, pid PresetID) (ft FT, err error) {
	ft = SealFT(pid)
	if err = p.storage.Get(user, &ft); err != nil {
		return ft, err
	}
	return ft, nil
}

// Incr - incr ft quantity
func (p *FTPocket) Incr(ctx context.Context, user string, pid PresetID, quantity int64) error {
	amount := SealFT(pid)
	amount.Quantity = quantity
	return p.storage.Update(user, amount.StoreName(), func(ft *FT) (*FT, error) {
		if ft == nil {
			return &amount, nil
		}
		ft.Quantity += amount.Quantity
		return ft, nil
	})
}

// DoContract - do contract
func (p *FTPocket) DoContract(ctx context.Context, user string, pid PresetID, contractID string,
	execute func(runtime *ContractRuntime) (*ContractRuntime, error),
) error {
	query := SealFT(pid)
	return p.storage.Update(user, query.StoreName(), func(ft *FT) (*FT, error) {
		if ft == nil {
			return nil, nil
		}
		if ft.Contracts == nil {
			ft.Contracts = make(map[string]ContractRuntime)
		}
		runtime, ok := ft.Contracts[contractID]
		if !ok {
			runtime = make(ContractRuntime)
		}
		newRuntime, err := execute(&runtime)
		if err != nil {
			return nil, err
		}
		ft.Contracts[contractID] = *newRuntime
		return ft, nil
	})
}

// Set - set ft to pocket
func (p *FTPocket) Set(ctx context.Context, user string, v FT) error {
	return p.storage.Set(user, &v)
}
