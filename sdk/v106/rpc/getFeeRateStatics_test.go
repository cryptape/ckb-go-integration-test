package rpc

import (
	"github.com/cryptape/ckb-go-integration-test/sdk"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestGetFeeRateStatics(t *testing.T) {
	for _, data := range testDatas {
		t.Run("getFeeRateStatics", func(t *testing.T) {
			statics, err := sdk.Client.GetFeeRateStatics(sdk.Ctx, data.target)
			if err != nil {
				return
			}
			log.Printf("median:%d\t mean:%d", statics.Median, statics.Mean)
			assert.NotNil(t, statics.Mean)
			assert.NotNil(t, statics.Median)
		})
	}
}

var testDatas = []struct {
	target uint64
}{
	{0},
	{1},
	{101},
	{102},
}
