package chains

import (
	"math"

	"github.com/migratooor/tokenLists/generators/common/models"
)

var ETHEREUM = TChain{
	ID:            1,
	Name:          `Ethereum`,
	Type:          `EVM`,
	LogoURI:       `https://token-assets-one.vercel.app/chains/1/logo-128.png`,
	IsTestNet:     false,
	RpcURI:        `https://eth.public-rpc.com`,
	MaxBlockRange: 100_000_000,
	MaxBatchSize:  math.MaxInt64,
	WeightRatio:   0.99,
	MulticallContract: TContractData{
		Address: `0xca11bde05977b3631167028862be2a173976ca11`,
		Block:   0,
	},
	Coin: models.TokenListToken{
		Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
		Name:     `Ethereum`,
		Symbol:   `ETH`,
		LogoURI:  `https://token-assets-one.vercel.app/api/token/1/0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE/logo-128.png`,
		ChainID:  1,
		Decimals: 18,
	},
	IgnoredTokens: []string{
		`0xdF5e0e81Dff6FAF3A7e52BA697820c5e32D806A8`,
	},
	ExtraTokens: []string{
		`0x9a96ec9B57Fb64FbC60B423d1f4da7691Bd35079`, // Ajna
		`0x6df0e641fc9847c0c6fde39be6253045440c14d3`, // DINERO
	},
}
