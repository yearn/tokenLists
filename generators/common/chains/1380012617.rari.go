package chains

import (
	"math"

	"github.com/migratooor/tokenLists/generators/common/models"
)

var RARI = TChain{
	ID:            1380012617,
	Name:          `Rari`,
	Type:          `EVM`,
	LogoURI:       `https://token-assets-one.vercel.app/chains/1380012617/logo-128.png`,
	IsTestNet:     false,
	RpcURI:        `https://mainnet.rpc.rarichain.org/http`,
	MaxBlockRange: 100_000_000,
	MaxBatchSize:  math.MaxInt64,
	WeightRatio:   0.0,
	MulticallContract: TContractData{
		Address: `0xb6D5B39F96d379569d47cC84024f3Cd78c5Ef651`,
		Block:   0,
	},
	Coin: models.TokenListToken{
		Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
		Name:     `Ethereum`,
		Symbol:   `ETH`,
		LogoURI:  `https://token-assets-one.vercel.app/api/token/1380012617/0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE/logo-128.png`,
		ChainID:  1380012617,
		Decimals: 18,
	},
	IgnoredTokens: []string{},
}
