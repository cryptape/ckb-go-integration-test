package rpc

import (
	"github.com/cryptape/ckb-go-integration-test/sdk"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetTransactionAndWitnessProof(t *testing.T) {

	t.Run("should return proof when query exist hash", func(t *testing.T) {
		t.Skip()
		gomega.RegisterTestingT(t)
		// get exist hash
		blockMsg, err := sdk.Client.GetBlock(sdk.Ctx, types.HexToHash("0xadaa049a601126abb71b08b4d7d522bd26ce50fe68ac75d7ebedd65b41ad8c1d"))
		gomega.Expect(err, nil)

		//txList := []
		var txs []types.Hash
		for _, tx := range blockMsg.Transactions {
			//txList = append(txList,tx.Hash)
			txs = append(txs, tx.Hash)
		}
		//todo wait sdk support GetTransactionAndWitnessProof
	})

}
