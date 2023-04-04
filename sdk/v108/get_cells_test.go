package v108

import (
	"github.com/cryptape/ckb-go-integration-test/sdk"
	"github.com/nervosnetwork/ckb-sdk-go/v2/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetCells(t *testing.T) {
	t.Run("should return sum(cell).Capacity  equal GetCellsCapacity ,when script_search_mode == extra", func(t *testing.T) {
		gomega.RegisterTestingT(t)
		getCellsResponse, err := sdk.Client.GetCells(sdk.Ctx, &indexer.SearchKey{
			Script: &types.Script{
				Args:     []byte(""),
				HashType: types.HashTypeData,
				CodeHash: types.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")},
			ScriptType:       types.ScriptTypeLock,
			ScriptSearchMode: types.ScriptSearchModeExact,
			WithData:         false,
		}, indexer.SearchOrderAsc, 1000, "")
		gomega.Expect(err, nil)
		// getCellsResponse  cap == getCells_capacity result
		var getCellsResponseCap uint64
		for _, object := range getCellsResponse.Objects {
			getCellsResponseCap = object.Output.Capacity + getCellsResponseCap
		}
		capacityResponse, err := sdk.Client.GetCellsCapacity(sdk.Ctx, &indexer.SearchKey{
			Script: &types.Script{
				Args:     []byte(""),
				HashType: types.HashTypeData,
				CodeHash: types.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")},
			ScriptType:       types.ScriptTypeLock,
			ScriptSearchMode: types.ScriptSearchModeExact,
		})
		gomega.Expect(getCellsResponseCap, capacityResponse.Capacity)
	})

	t.Run("should return sum(cell).Capacity  equal GetCellsCapacity  ,when script_search_mode == pre", func(t *testing.T) {
		gomega.RegisterTestingT(t)
		getCellsResponse, err := sdk.Client.GetCells(sdk.Ctx, &indexer.SearchKey{
			Script: &types.Script{
				Args:     []byte(""),
				HashType: types.HashTypeData,
				CodeHash: types.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")},
			ScriptType:       types.ScriptTypeLock,
			ScriptSearchMode: types.ScriptSearchModePrefix,
			WithData:         false,
		}, indexer.SearchOrderAsc, 1000, "")
		gomega.Expect(err, nil)
		capacityResponse, err := sdk.Client.GetCellsCapacity(sdk.Ctx, &indexer.SearchKey{
			Script: &types.Script{
				Args:     []byte(""),
				HashType: types.HashTypeData,
				CodeHash: types.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")},
			ScriptType:       types.ScriptTypeLock,
			ScriptSearchMode: types.ScriptSearchModePrefix,
			WithData:         false,
		})
		gomega.Expect(err, nil)
		var getCellsResponseCap uint64
		for _, object := range getCellsResponse.Objects {
			getCellsResponseCap = object.Output.Capacity + getCellsResponseCap
		}
		gomega.Expect(getCellsResponseCap, capacityResponse.Capacity)
	})
	t.Run("should return sum(cell).Capacity  equal GetCellsCapacity   when script_search_mode == nil", func(t *testing.T) {
		gomega.RegisterTestingT(t)
		getCellsResponse, err := sdk.Client.GetCells(sdk.Ctx, &indexer.SearchKey{
			Script: &types.Script{
				Args:     []byte(""),
				HashType: types.HashTypeData,
				CodeHash: types.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")},
			ScriptType: types.ScriptTypeLock,
			WithData:   false,
		}, indexer.SearchOrderAsc, 1000, "")
		capacityResponse, err := sdk.Client.GetCellsCapacity(sdk.Ctx, &indexer.SearchKey{
			Script: &types.Script{
				Args:     []byte(""),
				HashType: types.HashTypeData,
				CodeHash: types.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")},
			ScriptType: types.ScriptTypeLock,
			WithData:   false,
		})
		gomega.Expect(err, nil)
		var getCellsResponseCap uint64
		for _, object := range getCellsResponse.Objects {
			getCellsResponseCap = object.Output.Capacity + getCellsResponseCap
		}
		gomega.Expect(getCellsResponseCap, capacityResponse.Capacity)
	})
}
