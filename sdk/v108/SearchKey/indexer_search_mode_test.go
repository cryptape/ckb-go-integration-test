package SearchKey

import (
	"github.com/cryptape/ckb-go-integration-test/sdk"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/nervosnetwork/ckb-sdk-go/v2/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

//todo https://github.com/nervosnetwork/ckb-sdk-go/pull/202/commits
func Test_IndexerSearchMode(t *testing.T) {
	t.Run("GetTransactions_PrefixMode", func(t *testing.T) {
		s := &indexer.SearchKey{
			Script: &types.Script{
				CodeHash: types.HexToHash("0x58c5f491aba6d61678b7cf7edf4910b1f5e00ec0cde2f42e0abb4fd9aff25a63"),
				HashType: types.HashTypeType,
				Args:     ethcommon.FromHex("0xe53f35ccf63bb37a3bb0ac3b7f89808077a78eae"[0:4]),
			},
			ScriptType:       types.ScriptTypeLock,
			ScriptSearchMode: types.ScriptSearchModePrefix,
		}
		resp, err := sdk.Client.GetTransactions(sdk.Ctx, s, indexer.SearchOrderAsc, 10, "")
		assert.NoError(t, err)
		assert.True(t, len(resp.Objects) >= 1)
		assert.NotEqual(t, 0, resp.Objects[0].BlockNumber)
		assert.NotEqual(t, "", resp.Objects[0].IoType)
	})

	t.Run("GetTransactions_ExactMode", func(t *testing.T) {
		s1 := &indexer.SearchKey{
			Script: &types.Script{
				CodeHash: types.HexToHash("0x58c5f491aba6d61678b7cf7edf4910b1f5e00ec0cde2f42e0abb4fd9aff25a63"),
				HashType: types.HashTypeType,
				Args:     ethcommon.FromHex("0xe53f35ccf63bb37a3bb0ac3b7f89808077a78eae"[0:4]),
			},
			ScriptType:       types.ScriptTypeLock,
			ScriptSearchMode: types.ScriptSearchModeExact,
		}
		resp1, err := sdk.Client.GetTransactions(sdk.Ctx, s1, indexer.SearchOrderAsc, 10, "")
		assert.NoError(t, err)
		assert.Equal(t, 0, len(resp1.Objects))

		s2 := &indexer.SearchKey{
			Script: &types.Script{
				CodeHash: types.HexToHash("0x58c5f491aba6d61678b7cf7edf4910b1f5e00ec0cde2f42e0abb4fd9aff25a63"),
				HashType: types.HashTypeType,
				Args:     ethcommon.FromHex("0xe53f35ccf63bb37a3bb0ac3b7f89808077a78eae"),
			},
			ScriptType:       types.ScriptTypeLock,
			ScriptSearchMode: types.ScriptSearchModeExact,
		}
		resp2, err := sdk.Client.GetTransactions(sdk.Ctx, s2, indexer.SearchOrderAsc, 10, "")
		assert.NoError(t, err)
		assert.True(t, len(resp2.Objects) >= 1)
		assert.NotEqual(t, 0, resp2.Objects[0].BlockNumber)
		assert.NotEqual(t, "", resp2.Objects[0].IoType)
	})
}
