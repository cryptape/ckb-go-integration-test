package rpc

import (
	"github.com/cryptape/ckb-go-integration-test/sdk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTransaction(t *testing.T) {

	//todo https://github.com/nervosnetwork/ckb-sdk-go/pull/203
	t.Run("GetTransaction", func(t *testing.T) {
		header, err := sdk.Client.GetTipHeader(sdk.Ctx)
		if err != nil {
			return
		}
		block, err1 := sdk.Client.GetBlockByNumber(sdk.Ctx, header.Number)
		if err1 != nil {
			return
		}
		hash := block.Transactions[0].Hash
		transaction, err2 := sdk.Client.GetTransaction(sdk.Ctx, hash)
		if err2 != nil {
			return
		}
		assert.Nil(t, transaction.Cycles)
	})
}
