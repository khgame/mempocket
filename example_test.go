package tpocket_test

import (
	"context"
	"fmt"

	"gihub.com/khgame/tpocket"
	"github.com/khgame/memstore"
)

var weaponPocket tpocket.FTPocket

// SetWeapon - set weapon to user
func SetWeapon(ctx context.Context, user string, weaponPresetId int64, count int64) error {
	ft := tpocket.SealFT(weaponPresetId)
	ft.Quantity = count // set quantity to 100
	if err := weaponPocket.Set(ctx, user, ft); err != nil {
		return err
	}
	return nil
}

func initPocket() tpocket.FTPocket {
	// init pocket
	app := "game1"
	pocketName := "user_weapon"
	persistKey := fmt.Sprintf("%s:%s", app, pocketName)

	// create a storage
	storage := memstore.NewInMemoryStorage[tpocket.FT](persistKey)
	weaponPocket = tpocket.MakeFTPocket(context.Background(), app, pocketName, storage)
	return weaponPocket
}

// Example_FT tests the seal method of FT type with testify
func Example_FT() {
	// create ft pocket
	initPocket()

	pid := tpocket.PresetID(1234)
	if err := SetWeapon(context.Background(), "user001", pid, 100); err != nil {
		panic(err)
	}
	if err := weaponPocket.Incr(context.Background(), "user001", pid, 333); err != nil {
		panic(err)
	}
	v, err := weaponPocket.Get(context.Background(), "user001", pid)
	if err != nil {
		panic(err)
	}

	fmt.Printf("quantity: %d", v.Quantity)
	// Output:
	// quantity: 433
}
