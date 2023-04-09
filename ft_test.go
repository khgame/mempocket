package tpocket_test

import (
	"testing"

	"github.com/khgame/tpocket"
	"github.com/stretchr/testify/assert"
)

// Test_FT_Seal tests the seal method of FT type with testify
func Test_FT_Seal(t *testing.T) {
	ft := tpocket.SealFT(1)
	assert.Equal(t, tpocket.FT{
		PID: 1,
	}, ft)

	ft = tpocket.SealFT(int64(1))
	assert.Equal(t, tpocket.FT{
		PID: 1,
	}, ft)

	ft = tpocket.SealFT(int16(1))
	assert.Equal(t, tpocket.FT{
		PID: 1,
	}, ft)

	ft = tpocket.SealFT(int8(1))
	assert.Equal(t, tpocket.FT{
		PID: 1,
	}, ft)

	ft = tpocket.SealFT(int(1))
	assert.Equal(t, tpocket.FT{
		PID: 1,
	}, ft)

	ft = tpocket.SealFT(uint32(1))
	assert.Equal(t, tpocket.FT{
		PID: 1,
	}, ft)

	ft = tpocket.SealFT(uint16(1))
	assert.Equal(t, tpocket.FT{
		PID: 1,
	}, ft)

	ft = tpocket.SealFT(uint8(1))
	assert.Equal(t, tpocket.FT{
		PID: 1,
	}, ft)

	ft = tpocket.SealFT(uint(1))
	assert.Equal(t, tpocket.FT{
		PID: 1,
	}, ft)
}
