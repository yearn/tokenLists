package chains

import (
	"math"

	"github.com/migratooor/tokenLists/generators/common/models"
)

var POLYGON = TChain{
	ID:            137,
	Name:          `Polygon`,
	Type:          `EVM`,
	LogoURI:       `https://token-assets-one.vercel.app/chains/137/logo-128.png`,
	IsTestNet:     false,
	RpcURI:        `https://polygon.llamarpc.com`,
	MaxBlockRange: 100_000_000,
	MaxBatchSize:  math.MaxInt64,
	WeightRatio:   0.8,
	MulticallContract: TContractData{
		Address: `0xca11bde05977b3631167028862be2a173976ca11`,
		Block:   0,
	},
	Coin: models.TokenListToken{
		Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
		Name:     `Matic`,
		Symbol:   `MATIC`,
		LogoURI:  `https://token-assets-one.vercel.app/api/token/137/0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE/logo-128.png`,
		ChainID:  137,
		Decimals: 18,
	},
	IgnoredTokens: []string{
		`0xec6432B90e7fD4d9f872cc5C781f05B617DB861E`,
	},
	ExtraTokens: []string{
		`0x3c499c542cef5e3811e1192ce70d8cc03d5c3359`, // USDC
	},
}
