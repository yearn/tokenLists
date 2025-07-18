package chains

import (
	"math"

	"github.com/migratooor/tokenLists/generators/common/models"
)

var BINANCE_SMART_CHAIN = TChain{
	ID:            56,
	Name:          `BSC`,
	Type:          `EVM`,
	LogoURI:       `https://token-assets-one.vercel.app/chains/56/logo-128.png`,
	IsTestNet:     false,
	RpcURI:        `https://1rpc.io/bnb`,
	MaxBlockRange: 100_000_000,
	MaxBatchSize:  math.MaxInt64,
	WeightRatio:   0.75,
	MulticallContract: TContractData{
		Address: `0xca11bde05977b3631167028862be2a173976ca11`,
		Block:   0,
	},
	Coin: models.TokenListToken{
		Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
		Name:     `Binance Smart Chain`,
		Symbol:   `BNB`,
		LogoURI:  `https://token-assets-one.vercel.app/api/token/56/0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE/logo-128.png`,
		ChainID:  56,
		Decimals: 18,
	},
	IgnoredTokens: []string{
		`0xc00e94Cb662C3520282E6f5717214004A7f26888`,
	},
}
