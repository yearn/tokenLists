package chains

import (
	"math"

	"github.com/migratooor/tokenLists/generators/common/models"
)

var POLYGON_ZKEVM = TChain{
	ID:            1101,
	Name:          `Polygon ZKEvm`,
	Type:          `EVM`,
	LogoURI:       `https://token-assets-one.vercel.app/chains/1101/logo-128.png`,
	IsTestNet:     false,
	RpcURI:        `https://endpoints.omniatech.io/v1/polygon-zkevm/mainnet/public`,
	MaxBlockRange: 100_000_000,
	MaxBatchSize:  math.MaxInt64,
	WeightRatio:   0.05,
	MulticallContract: TContractData{
		Address: `0xca11bde05977b3631167028862be2a173976ca11`,
		Block:   0,
	},
	Coin: models.TokenListToken{
		Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
		Name:     `Matic`,
		Symbol:   `MATIC`,
		LogoURI:  `https://token-assets-one.vercel.app/api/token/1101/0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE/logo-128.png`,
		ChainID:  1101,
		Decimals: 18,
	},
	IgnoredTokens: []string{},
}
