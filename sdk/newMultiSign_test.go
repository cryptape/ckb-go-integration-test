package sdk

import (
	"testing"

	"github.com/nervosnetwork/ckb-sdk-go/v2/address"
	"github.com/nervosnetwork/ckb-sdk-go/v2/collector"
	"github.com/nervosnetwork/ckb-sdk-go/v2/collector/builder"
	"github.com/nervosnetwork/ckb-sdk-go/v2/systemscript"
	"github.com/nervosnetwork/ckb-sdk-go/v2/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/v2/transaction/signer"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"github.com/stretchr/testify/assert"
)

func TestMultisigTxPoolAccept(t *testing.T) {
	// 跳过实际发送交易，仅测试交易验证
	// t.Skip("Skip actual transaction sending, only for manual testing")

	t.Run("TestLegacyMultisigTxPoolAccept", func(t *testing.T) {
		// 创建Legacy多签配置
		multisigConfig := systemscript.NewMultisigConfig(0, 2)
		multisigConfig.AddKeyHash([]byte{0x73, 0x36, 0xb0, 0xba, 0x90, 0x06, 0x84, 0xcb, 0x3c, 0xb0, 0x0f, 0x0d, 0x46, 0xd4, 0xf6, 0x4c, 0x09, 0x94, 0xa5, 0x62})
		multisigConfig.AddKeyHash([]byte{0x57, 0x24, 0xc1, 0xe3, 0x92, 0x5a, 0x52, 0x06, 0x94, 0x4d, 0x75, 0x3a, 0x6f, 0x3e, 0xda, 0xed, 0xf9, 0x77, 0xd7, 0x7f})

		// 验证Legacy多签交易
		result, err := testMultisigTxPoolAccept(t, multisigConfig, systemscript.MultisigLegacy)

		// 断言验证
		assert.NoError(t, err, "Legacy multisig transaction should be accepted")
		if err == nil {
			assert.NotNil(t, result, "Result should not be nil")
			assert.Greater(t, result.Cycles, uint64(0), "Cycles should be greater than 0")
			assert.Greater(t, result.Fee, uint64(0), "Fee should be greater than 0")
			t.Logf("Legacy multisig transaction accepted with cycles: %d, fee: %d", result.Cycles, result.Fee)
		}
	})

	t.Run("TestV2MultisigTxPoolAccept", func(t *testing.T) {
		// 创建V2多签配置
		multisigConfig := systemscript.NewMultisigConfig(0, 2)
		multisigConfig.AddKeyHash([]byte{0x73, 0x36, 0xb0, 0xba, 0x90, 0x06, 0x84, 0xcb, 0x3c, 0xb0, 0x0f, 0x0d, 0x46, 0xd4, 0xf6, 0x4c, 0x09, 0x94, 0xa5, 0x62})
		multisigConfig.AddKeyHash([]byte{0x57, 0x24, 0xc1, 0xe3, 0x92, 0x5a, 0x52, 0x06, 0x94, 0x4d, 0x75, 0x3a, 0x6f, 0x3e, 0xda, 0xed, 0xf9, 0x77, 0xd7, 0x7f})

		// 验证V2多签交易
		result, err := testMultisigTxPoolAccept(t, multisigConfig, systemscript.MultisigV2)

		// 断言验证
		assert.NoError(t, err, "V2 multisig transaction should be accepted")
		if err == nil {
			assert.NotNil(t, result, "Result should not be nil")
			assert.Greater(t, result.Cycles, uint64(0), "Cycles should be greater than 0")
			assert.Greater(t, result.Fee, uint64(0), "Fee should be greater than 0")
			t.Logf("V2 multisig transaction accepted with cycles: %d, fee: %d", result.Cycles, result.Fee)
		}
	})

	t.Run("CompareMultisigVersions", func(t *testing.T) {
		// 创建多签配置
		multisigConfig := systemscript.NewMultisigConfig(0, 2)
		multisigConfig.AddKeyHash([]byte{0x73, 0x36, 0xb0, 0xba, 0x90, 0x06, 0x84, 0xcb, 0x3c, 0xb0, 0x0f, 0x0d, 0x46, 0xd4, 0xf6, 0x4c, 0x09, 0x94, 0xa5, 0x62})
		multisigConfig.AddKeyHash([]byte{0x57, 0x24, 0xc1, 0xe3, 0x92, 0x5a, 0x52, 0x06, 0x94, 0x4d, 0x75, 0x3a, 0x6f, 0x3e, 0xda, 0xed, 0xf9, 0x77, 0xd7, 0x7f})

		// 验证Legacy多签交易
		legacyResult, legacyErr := testMultisigTxPoolAccept(t, multisigConfig, systemscript.MultisigLegacy)

		// 验证V2多签交易
		v2Result, v2Err := testMultisigTxPoolAccept(t, multisigConfig, systemscript.MultisigV2)

		// 断言验证
		assert.NoError(t, legacyErr, "Legacy multisig transaction should be accepted")
		assert.NoError(t, v2Err, "V2 multisig transaction should be accepted")

		if legacyErr == nil && v2Err == nil {
			// 比较两种多签的性能差异
			t.Logf("Legacy multisig: cycles=%d, fee=%d", legacyResult.Cycles, legacyResult.Fee)
			t.Logf("V2 multisig: cycles=%d, fee=%d", v2Result.Cycles, v2Result.Fee)

			// 计算差异百分比
			cyclesDiff := float64(v2Result.Cycles) / float64(legacyResult.Cycles) * 100
			feeDiff := float64(v2Result.Fee) / float64(legacyResult.Fee) * 100

			t.Logf("V2 vs Legacy: cycles=%.2f%%, fee=%.2f%%", cyclesDiff, feeDiff)
		}
	})
}

// testMultisigTxPoolAccept 创建并验证多签交易
func testMultisigTxPoolAccept(t *testing.T, multisigConfig *systemscript.MultisigConfig, multisigVersion systemscript.MultisigVersion) (*types.EntryCompleted, error) {
	network := types.NetworkTest

	// 创建多签地址
	args := multisigConfig.Hash160()
	var script *types.Script
	var err error

	if multisigVersion == systemscript.MultisigLegacy {
		script = &types.Script{
			CodeHash: systemscript.GetCodeHash(network, systemscript.Secp256k1Blake160MultisigAllLegacy),
			HashType: types.HashTypeType,
			Args:     args,
		}
	} else {
		script = &types.Script{
			CodeHash: systemscript.GetCodeHash(network, systemscript.Secp256k1Blake160MultisigAllV2),
			HashType: types.HashTypeData1,
			Args:     args,
		}
	}

	sender, _ := address.Address{
		Script:  script,
		Network: network,
	}.Encode()

	receiver := "ckt1qzda0cr08m85hc8jlnfp3zer7xulejywt49kt2rr0vthywaa50xwsq2qf8keemy2p5uu0g0gn8cd4ju23s5269qk8rg4r"

	// 创建输入迭代器
	iterator, err := collector.NewLiveCellIteratorFromAddress(Client, sender)
	if err != nil {
		return nil, err
	}

	// 构建交易
	builder := builder.NewCkbTransactionBuilder(network, iterator)
	builder.FeeRate = 1000
	if err := builder.AddOutputByAddress(receiver, 50100000000); err != nil {
		return nil, err
	}
	builder.AddChangeOutputByAddress(sender)
	txWithGroups, err := builder.Build(multisigConfig)
	if err != nil {
		return nil, err
	}

	// 签名交易
	txSigner := signer.GetTransactionSignerInstance(network)
	// 第一个签名
	ctx1, _ := transaction.NewContextWithPayload("0x4fd809631a6aa6e3bb378dd65eae5d71df895a82c91a615a1e8264741515c79c", multisigConfig)
	if _, err = txSigner.SignTransaction(txWithGroups, ctx1); err != nil {
		return nil, err
	}
	// 第二个签名
	ctx2, _ := transaction.NewContextWithPayload("0x7438f7b35c355e3d2fb9305167a31a72d22ddeafb80a21cc99ff6329d92e8087", multisigConfig)
	if _, err = txSigner.SignTransaction(txWithGroups, ctx2); err != nil {
		return nil, err
	}

	// 使用TestTxPoolAccept验证交易
	result, err := Client.TestTxPoolAccept(Ctx, txWithGroups.TxView)
	return result, err
}
