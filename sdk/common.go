package sdk

import (
	"context"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
)

var (
	Client, _ = rpc.DialContext(context.Background(), "https://testnet.ckbapp.dev")
	Ctx       = context.Background()
)
