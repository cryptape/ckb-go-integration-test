package rpc

import (
	"github.com/cryptape/ckb-go-integration-test/sdk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBlockByNumber(t *testing.T) {
	t.Run("getBlockByNumber", func(t *testing.T) {
		block, err := sdk.Client.GetBlockByNumber(sdk.Ctx, 1)
		if err != nil {
			return
		}
		assert.NotNil(t, block.Transactions)
	})
}
