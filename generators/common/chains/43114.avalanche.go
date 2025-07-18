package chains

import (
	"math"

	"github.com/migratooor/tokenLists/generators/common/models"
)

var AVALANCHE = TChain{
	ID:            43114,
	Name:          `Avax`,
	Type:          `EVM`,
	LogoURI:       `https://token-assets-one.vercel.app/chains/43114/logo-128.png`,
	IsTestNet:     false,
	RpcURI:        `https://1rpc.io/avax/c`,
	MaxBlockRange: 100_000_000,
	MaxBatchSize:  math.MaxInt64,
	WeightRatio:   0.5,
	MulticallContract: TContractData{
		Address: `0xca11bde05977b3631167028862be2a173976ca11`,
		Block:   0,
	},
	Coin: models.TokenListToken{
		Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
		Name:     `Ethereum`,
		Symbol:   `ETH`,
		LogoURI:  `https://token-assets-one.vercel.app/api/token/43114/0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE/logo-128.png`,
		ChainID:  43114,
		Decimals: 18,
	},
	IgnoredTokens: []string{},
}
