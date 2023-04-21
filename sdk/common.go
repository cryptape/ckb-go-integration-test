package sdk

import (
	"context"
	"github.com/nervosnetwork/ckb-sdk-go/v2/lightclient"
	"github.com/nervosnetwork/ckb-sdk-go/v2/mocking"
	"github.com/nervosnetwork/ckb-sdk-go/v2/rpc"
)

var (
	Client, _ = rpc.DialContext(context.Background(), "https://testnet.ckbapp.dev")
	Ctx       = context.Background()

	C, _       = lightclient.DialMockContext(context.Background(), "http://localhost:9000")
	MockClient = interface{}(C.GetRawClient()).(*mocking.MockClient)
)
