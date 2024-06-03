package ethclient

import (
    "context"
    "fmt"
    "os"
    "path/filepath"

    "github.com/ethereum/go-ethereum/ethclient"
    gethRPC "github.com/ethereum/go-ethereum/rpc"
    "github.com/pkg/errors"
    "github.com/prysmaticlabs/prysm/beacon-chain/db"
    contracts "github.com/prysmaticlabs/prysm/contracts/deposit-contract"
    "github.com/prysmaticlabs/prysm/shared/event"
    "github.com/prysmaticlabs/prysm/shared/params"
)


type RPcClient interface {
	BatchCall(b []gethRPC.BatchElem) error
}

type Service struct {
	
}