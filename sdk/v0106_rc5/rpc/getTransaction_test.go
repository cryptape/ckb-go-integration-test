package rpc

import (
	"github.com/cryptape/ckb-go-integration-test/sdk"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetTransaction(t *testing.T) {

	t.Run("GetTransaction", func(t *testing.T) {
		gomega.RegisterTestingT(t)
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
		gomega.Expect(transaction.Cycles).To(gomega.Not(nil))
	})
}
