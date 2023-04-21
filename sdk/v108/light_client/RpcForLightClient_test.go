package light_client

import (
	"github.com/cryptape/ckb-go-integration-test/sdk"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nervosnetwork/ckb-sdk-go/v2/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/v2/lightclient"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

var scriptForTest = &types.Script{
	CodeHash: types.HexToHash("0x9bd7e06f3ecf4be0f2fcd2188b23f1b9fcc88e5d4b65a8637b17723bbda3cce8"),
	HashType: types.HashTypeType,
	Args:     ethcommon.FromHex("0x4049ed9cec8a0d39c7a1e899f0dacb8a8c28ad14"),
}

//todo https://github.com/nervosnetwork/ckb-sdk-go/pull/198

func TestLightCLientRpc(t *testing.T) {

	t.Run("TestSetScripts", func(t *testing.T) {
		scriptDetail := lightclient.ScriptDetail{
			// ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq2qf8keemy2p5uu0g0gn8cd4ju23s5269qk8rg4r
			Script:      scriptForTest,
			ScriptType:  types.ScriptTypeLock,
			BlockNumber: 7033100,
		}
		sdk.MockClient.LoadMockingTestFromFile(t, "set_scripts", []*lightclient.ScriptDetail{&scriptDetail})
		err := sdk.C.SetScripts(sdk.Ctx, []*lightclient.ScriptDetail{&scriptDetail})
		assert.NoError(t, err)
	})

	t.Run("TestGetScrpits", func(t *testing.T) {
		sdk.MockClient.LoadMockingTestFromFile(t, "get_scripts")
		scriptDetails, err := sdk.C.GetScripts(sdk.Ctx)
		assert.NoError(t, err)
		assert.NotEmpty(t, scriptDetails)
		assert.NotEmpty(t, scriptDetails[0].Script)
		assert.NotEmpty(t, scriptDetails[0].ScriptType)
	})

	t.Run("TestGetTipHeader", func(t *testing.T) {
		sdk.MockClient.LoadMockingTestFromFile(t, "get_tip_header")
		header, err := sdk.C.GetTipHeader(sdk.Ctx)
		assert.NoError(t, err)
		assert.NotEmpty(t, header)
	})

	t.Run("TestGetGenesisBlock", func(t *testing.T) {
		sdk.MockClient.LoadMockingTestFromFile(t, "get_genesis_block")
		block, err := sdk.C.GetGenesisBlock(sdk.Ctx)
		assert.NoError(t, err)
		assert.NotEmpty(t, block)
		assert.NotEmpty(t, block.Transactions)
		assert.NotEmpty(t, block.Header)
	})

	t.Run("TestGetHeader", func(t *testing.T) {
		sdk.MockClient.LoadMockingTestFromFile(t, "get_header", types.HexToHash("0x10639e0895502b5688a6be8cf69460d76541bfa4821629d86d62ba0aae3f9606"))
		header, err := sdk.C.GetHeader(sdk.Ctx,
			types.HexToHash("0x10639e0895502b5688a6be8cf69460d76541bfa4821629d86d62ba0aae3f9606"))
		assert.NoError(t, err)
		assert.NotEmpty(t, header)
	})

	t.Run("TestGetTransaction", func(t *testing.T) {
		sdk.MockClient.LoadMockingTestFromFile(t, "get_transaction", types.HexToHash("0x8f8c79eb6671709633fe6a46de93c0fedc9c1b8a6527a18d3983879542635c9f"))
		txWitHeader, err := sdk.C.GetTransaction(sdk.Ctx,
			types.HexToHash("0x8f8c79eb6671709633fe6a46de93c0fedc9c1b8a6527a18d3983879542635c9f"))
		assert.NoError(t, err)
		assert.NotEmpty(t, txWitHeader.Transaction)
		assert.NotEmpty(t, txWitHeader.TxStatus)
	})

	t.Run(" TestFetchHeader", func(t *testing.T) {
		sdk.MockClient.LoadMockingTestFromFile(t, "fetch_header", types.HexToHash("0xcb5eae958e3ea24b0486a393133aa33d51224ffaab3c4819350095b3446e4f70"))
		fetchedHeader, err := sdk.C.FetchHeader(sdk.Ctx,
			types.HexToHash("0xcb5eae958e3ea24b0486a393133aa33d51224ffaab3c4819350095b3446e4f70"))
		assert.NoError(t, err)
		assert.NotEmpty(t, fetchedHeader.Status)
		assert.NotEmpty(t, *fetchedHeader.Data)
	})

	t.Run("TestFetchTransaction", func(t *testing.T) {
		sdk.MockClient.LoadMockingTestFromFile(t, "fetch_transaction", types.HexToHash("0x716e211698d3d9499aae7903867c744b67b539beeceddad330e73d1b6b617aef"))
		fetchedTransaction, err := sdk.C.FetchTransaction(sdk.Ctx,
			types.HexToHash("0x716e211698d3d9499aae7903867c744b67b539beeceddad330e73d1b6b617aef"))
		assert.NoError(t, err)
		assert.NotEmpty(t, fetchedTransaction.Status)
	})

	t.Run("TestGetCells", func(t *testing.T) {
		s := &indexer.SearchKey{
			Script:     scriptForTest,
			ScriptType: types.ScriptTypeLock,
		}
		sdk.MockClient.LoadMockingTestFromFile(t, "get_cells", s, indexer.SearchOrderAsc, hexutil.Uint64(10)) // this is a special cast, make it same with the actual call
		resp, err := sdk.C.GetCells(sdk.Ctx, s, indexer.SearchOrderAsc, 10, "")
		assert.NoError(t, err)
		assert.NotEmpty(t, resp)
		assert.NotEmpty(t, resp.Objects[0].BlockNumber)
		assert.NotEmpty(t, resp.Objects[0].OutPoint)
		assert.NotEmpty(t, resp.Objects[0].Output)
	})

	t.Run("TestGetTransactions", func(t *testing.T) {
		s := &indexer.SearchKey{
			Script:     scriptForTest,
			ScriptType: types.ScriptTypeLock,
		}
		sdk.MockClient.LoadMockingTestFromFile(t, "get_transactions", s, indexer.SearchOrderAsc, hexutil.Uint64(10)) // this is a special cast, make it same with the actual call
		resp, err := sdk.C.GetTransactions(sdk.Ctx, s, indexer.SearchOrderAsc, 10, "")
		assert.NoError(t, err)
		assert.NotEmpty(t, resp)
		assert.NotEmpty(t, resp.Objects[0].BlockNumber)
		assert.NotEmpty(t, resp.Objects[0].IoType)
		assert.NotEmpty(t, resp.Objects[0].Transaction)
	})

	t.Run("TestGetTransactionsGrouped", func(t *testing.T) {
		s := &indexer.SearchKey{
			Script:     scriptForTest,
			ScriptType: types.ScriptTypeLock,
		}
		payload := &struct {
			indexer.SearchKey
			GroupByTransaction bool `json:"group_by_transaction"`
		}{
			SearchKey:          *s,
			GroupByTransaction: true,
		}
		sdk.MockClient.LoadMockingTestFromFilePatched(t, "get_transactions_grouped", "get_transactions", payload, indexer.SearchOrderAsc, hexutil.Uint64(10)) // this is a special cast, make it same with the actual call
		resp, err := sdk.C.GetTransactionsGrouped(sdk.Ctx, s, indexer.SearchOrderAsc, 10, "")
		assert.NoError(t, err)
		assert.NotEqual(t, 0, len(resp.Objects))
		assert.NotEqual(t, 0, resp.Objects[0].BlockNumber)
		assert.NotEqual(t, 0, resp.Objects[0].Transaction)
		assert.NotEmpty(t, resp.Objects[0].Cells[0])
		assert.NotEmpty(t, resp.Objects[0].Cells[0].IoType)
	})

	t.Run("TestGetCellsCapacity", func(t *testing.T) {
		s := &indexer.SearchKey{
			Script:     scriptForTest,
			ScriptType: types.ScriptTypeLock,
		}
		sdk.MockClient.LoadMockingTestFromFile(t, "get_cells_capacity", s)
		resp, err := sdk.C.GetCellsCapacity(sdk.Ctx, s)
		assert.NoError(t, err)
		assert.NotEmpty(t, resp.BlockNumber)
		assert.NotEmpty(t, resp.BlockHash)
		assert.NotEmpty(t, resp.Capacity)
	})

	t.Run("TestGetPeers", func(t *testing.T) {
		sdk.MockClient.LoadMockingTestFromFile(t, "get_peers")
		peers, err := sdk.C.GetPeers(sdk.Ctx)
		if err != nil {
			t.Fatal(err)
		}
		assert.True(t, len(peers) > 0)
		assert.True(t, len(peers[0].Addresses) > 0)
		assert.True(t, len(peers[0].Protocols) > 0)
	})

	t.Run("TestClient_LocalNodeInfo", func(t *testing.T) {
		sdk.MockClient.LoadMockingTestFromFile(t, "local_node_info")
		nodeInfo, err := sdk.C.LocalNodeInfo(sdk.Ctx)
		if err != nil {
			t.Fatal(err)
		}
		assert.True(t, len(nodeInfo.Addresses) > 0)
		assert.True(t, len(nodeInfo.Protocols) > 0)
		assert.True(t, len(nodeInfo.Protocols[0].SupportVersions) > 0)
	})

}
