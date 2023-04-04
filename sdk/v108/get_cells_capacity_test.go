package v108

import (
	"fmt"
	"github.com/cryptape/ckb-go-integration-test/sdk"
	"github.com/nervosnetwork/ckb-sdk-go/v2/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/onsi/gomega"
	"testing"
)

func TestGetCellsCapacity(t *testing.T) {

	t.Run("should return extraResult ,when script_search_mode == extra", func(t *testing.T) {
		gomega.RegisterTestingT(t)
		capacityResponse, err := sdk.Client.GetCellsCapacity(sdk.Ctx, &indexer.SearchKey{
			Script: &types.Script{
				Args:     []byte(""),
				HashType: types.HashTypeData,
				CodeHash: types.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")},
			ScriptType:       types.ScriptTypeLock,
			ScriptSearchMode: types.ScriptSearchModeExact,
			WithData:         false,
		})
		fmt.Println("capacityResponse :", capacityResponse.Capacity)

		gomega.Expect(err, nil)
		gomega.Expect(capacityResponse.BlockNumber > 0, true)
		gomega.Expect(len(capacityResponse.BlockHash.Hex()) == len("0x69b62910dff2a608e112695f19ef3c399bae7b8e9b443b34661316b358f0f75b"), true)
		gomega.Expect(capacityResponse.Capacity > 0, true)
	})

	t.Run("should return capacityResponse.Capacity > script_search_mode:extra  ,when script_search_mode == pre", func(t *testing.T) {
		gomega.RegisterTestingT(t)
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
		capacityResponseByExact, err := sdk.Client.GetCellsCapacity(sdk.Ctx, &indexer.SearchKey{
			Script: &types.Script{
				Args:     []byte(""),
				HashType: types.HashTypeData,
				CodeHash: types.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")},
			ScriptType:       types.ScriptTypeLock,
			ScriptSearchMode: types.ScriptSearchModeExact,
			WithData:         false,
		})
		gomega.Expect(capacityResponse.Capacity > capacityResponseByExact.Capacity, true)
	})
	t.Run("should return capacityResponse eq script_search_mode:pre result  when script_search_mode == nil", func(t *testing.T) {
		gomega.RegisterTestingT(t)
		capacityResponse, err := sdk.Client.GetCellsCapacity(sdk.Ctx, &indexer.SearchKey{
			Script: &types.Script{
				Args:     []byte(""),
				HashType: types.HashTypeData,
				CodeHash: types.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")},
			ScriptType: types.ScriptTypeLock,
			WithData:   false,
		})
		gomega.Expect(err, nil)

		capacityResponseByPrefix, err := sdk.Client.GetCellsCapacity(sdk.Ctx, &indexer.SearchKey{
			Script: &types.Script{
				Args:     []byte(""),
				HashType: types.HashTypeData,
				CodeHash: types.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")},
			ScriptType:       types.ScriptTypeLock,
			ScriptSearchMode: types.ScriptSearchModePrefix,
			WithData:         false,
		})
		gomega.Expect(err, nil)
		gomega.Expect(capacityResponse, capacityResponseByPrefix)

	})

}
