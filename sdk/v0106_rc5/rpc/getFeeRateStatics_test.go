package rpc

import (
	"github.com/cryptape/ckb-go-integration-test/sdk"
	"github.com/onsi/gomega"
	"log"
	"testing"
)

func TestGetFeeRateStatics(t *testing.T) {
	for _, data := range testDatas {
		t.Run("getFeeRateStatics", func(t *testing.T) {
			gomega.RegisterTestingT(t)
			statics, err := sdk.Client.GetFeeRateStatics(sdk.Ctx, data.target)
			if err != nil {
				return
			}
			log.Printf("median:%d\t mean:%d", statics.Median, statics.Mean)
			//gomega.Expect(statics.Mean).To(gomega.Not(nil))
			//gomega.Expect(statics.Median).To(gomega.Not(nil))
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
