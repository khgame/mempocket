package tpocket

import (
	"context"

	"github.com/khgame/memstore"
)

// FTPocket : app_id:pocket_name
// user -->|pid| { pid, quantity } ...
type FTPocket Pocket[FT]

// Get - get ft from pocket
func (p *FTPocket) Get(ctx context.Context, user string, pid PresetID) (ft FT, err error) {
	ft = SealFT(pid)
	if err = p.storage.Get(user, &ft); err != nil {
		return ft, err
	}
	return ft, nil
}

// MGet - get a list of ft from pocket
func (p *FTPocket) MGet(ctx context.Context, user string, pids []PresetID) (fts []FT, err error) {
	visited := make(map[PresetID]bool)
	for _, pid := range pids {
		if visited[pid] {
			continue
		}
		visited[pid] = true
		ft := SealFT(pid)
		if err = p.storage.Get(user, &ft); err != nil {
			return nil, err
		}
		fts = append(fts, ft)
	}
	return fts, nil
}

// List - get ft from pocket
func (p *FTPocket) List(ctx context.Context, user string, filter func(ftName string) bool) (fts []FT, err error) {
	ftNames, err := p.storage.List(user)
	if err != nil {
		return nil, err
	}
	// because memstore is running in memory, so it's fast enough to get all ft
	// thus we don't need to provide a batch-get method in the memstore
	// todo:
	// - if we need to support a large number of ft, we should provide a batch-get method in the memstore
	// - if the storage is not in memory, we should provide a batch-get method in the memstore
	queryLst := make([]PresetID, 0, len(ftNames))

	for _, ftName := range ftNames {
		if filter != nil && !filter(ftName) {
			continue
		}
		pid, eDecode := DecodeFTStoreName[PresetID](ftName)
		if eDecode != nil {
			return nil, eDecode
		}
		queryLst = append(queryLst, pid)
	}

	return p.MGet(ctx, user, queryLst)
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
			ft = &query
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

func MakeFTPocket(ctx context.Context, appID string, metaStr string, storage memstore.Storage[FT]) FTPocket {
	return FTPocket{
		AppID:   appID,
		Meta:    metaStr,
		storage: storage,
	}
}
