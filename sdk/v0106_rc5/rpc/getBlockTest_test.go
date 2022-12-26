package rpc

import (
	"github.com/cryptape/ckb-go-integration-test/sdk"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBlock(t *testing.T) {
	t.Run("getBlock", func(t *testing.T) {
		gomega.RegisterTestingT(t)
		blockHash, err := sdk.Client.GetBlockHash(sdk.Ctx, 1)
		if err != nil {
			return
		}
		block, err := sdk.Client.GetBlock(sdk.Ctx, types.HexToHash(blockHash.String()))
		if err != nil {
			return
		}
		assert.NotNil(t, block.Transactions)
	})
}
