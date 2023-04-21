package rpc

import (
	"github.com/cryptape/ckb-go-integration-test/sdk"
	"github.com/onsi/gomega"
	"testing"
)

func TestEstimateCycles(t *testing.T) {
	t.Run("estimateCyclesTest", func(t *testing.T) {
		gomega.RegisterTestingT(t)
		header, err := sdk.Client.GetTipHeader(sdk.Ctx)
		if err != nil {
			return
		}
		block, err1 := sdk.Client.GetBlockByNumber(sdk.Ctx, header.Number)
		if err1 != nil {
			return
		}
		tx := block.Transactions[0]
		cycles, err := sdk.Client.EstimateCycles(sdk.Ctx, tx)
		if err != nil {
			return
		}
		gomega.Expect(cycles.Cycles).To(gomega.Equal(uint64(0)))
	})
}
