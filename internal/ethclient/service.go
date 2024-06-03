package ethclient

import (
    "context"
    "math/big"
    "github.com/ethereum/go-ethereum/rpc"
)

type Client struct {
    rpcClient *rpc.Client
}

// Dial connects to the given IPC path.
func Dial(ipcPath string) (*Client, error) {
    rpcClient, err := rpc.Dial(ipcPath)
    if err != nil {
        return nil, err
    }
    return &Client{rpcClient}, nil
}

// BlockNumber retrieves the current block number.
func (c *Client) BlockNumber() (uint64, error) {
    var result *big.Int
    err := c.rpcClient.CallContext(context.Background(), &result, "eth_blockNumber")
    if err != nil {
        return 0, err
    }
    return result.Uint64(), nil
}
