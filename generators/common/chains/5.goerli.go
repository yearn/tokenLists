package chains

import (
	"math"

	"github.com/migratooor/tokenLists/generators/common/models"
)

var GOERLI = TChain{
	ID:            5,
	Name:          `Goerli`,
	Type:          `EVM`,
	RpcURI:        `https://gateway.tenderly.co/public/goerli`,
	LogoURI:       `https://token-assets-one.vercel.app/chains/1/logo-128.png`,
	IsTestNet:     true,
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
		LogoURI:  `https://token-assets-one.vercel.app/api/token/5/0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE/logo-128.png`,
		ChainID:  1,
		Decimals: 18,
	},
	IgnoredTokens: []string{},
}
