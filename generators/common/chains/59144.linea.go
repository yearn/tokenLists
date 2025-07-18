package chains

import (
	"math"

	"github.com/migratooor/tokenLists/generators/common/models"
)

var LINEA = TChain{
	ID:            59144,
	Name:          `Linea`,
	Type:          `EVM`,
	LogoURI:       `https://token-assets-one.vercel.app/chains/59144/logo-128.png`,
	IsTestNet:     false,
	RpcURI:        `https://rpc.linea.build`,
	MaxBlockRange: 100_000_000,
	MaxBatchSize:  math.MaxInt64,
	WeightRatio:   0.0001,
	MulticallContract: TContractData{
		Address: `0xca11bde05977b3631167028862be2a173976ca11`,
		Block:   0,
	},
	Coin: models.TokenListToken{
		Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
		Name:     `Ethereum`,
		Symbol:   `ETH`,
		LogoURI:  `https://token-assets-one.vercel.app/api/token/59144/0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE/logo-128.png`,
		ChainID:  59144,
		Decimals: 18,
	},
	IgnoredTokens: []string{},
}
