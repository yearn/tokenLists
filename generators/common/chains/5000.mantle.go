package chains

import (
	"math"

	"github.com/migratooor/tokenLists/generators/common/models"
)

var MANTLE = TChain{
	ID:            5000,
	Name:          `Mantle`,
	Type:          `EVM`,
	LogoURI:       `https://token-assets-one.vercel.app/chains/5000/logo-128.png`,
	IsTestNet:     false,
	RpcURI:        `https://rpc.mantle.xyz`,
	MaxBlockRange: 100_000_000,
	MaxBatchSize:  math.MaxInt64,
	WeightRatio:   0.05,
	MulticallContract: TContractData{
		Address: `0xca11bde05977b3631167028862be2a173976ca11`,
		Block:   0,
	},
	Coin: models.TokenListToken{
		Address:  `0xdeaddeaddeaddeaddeaddeaddeaddeaddead0000`,
		Name:     `Mantle`,
		Symbol:   `MNT`,
		LogoURI:  `https://token-assets-one.vercel.app/api/token/5000/0xdeaddeaddeaddeaddeaddeaddeaddeaddead0000/logo-128.png`,
		Decimals: 18,
		ChainID:  5000,
	},
	IgnoredTokens: []string{},
}
