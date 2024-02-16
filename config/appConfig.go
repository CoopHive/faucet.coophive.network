package config

import (
	"fmt"
	"os"
	"path"

	"github.com/coophive/faucet.coophive.network/enums"
)

const DEFAULT_DEALER = "std-autoaccept"

var appConfig = configMap[string]{
	enums.DEBUG: {
		desc:       "debug mode",
		defaultVal: "false",
	},
	enums.DEALER: {
		"Dealer for accepting/denying solver-matched deals",
		DEFAULT_DEALER,
	},
	// 	enums.DEALER_PATH: {
	// 		"Dealer path for resource provider",
	// 		"./autoaccept.so", // TODO: set default to empty
	// 	},

	enums.APP_DIR: {
		"App Location Directory",
		"$HOME/coophive-faucet",
	},
	enums.NETWORK: {
		fmt.Sprintf("supported networks:%v", NETWORKS),
		defaultNetwork,
	},

	enums.WEB3_RPC_URL: {
		"rpc url",
		"ws://testnet.co-ophive.network:8546",
	},
	enums.WEB3_CHAIN_ID: {
		"chain id of the network",
		"1337",
	},
	enums.WEB3_PRIVATE_KEY: {
		"private key",
		"0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
	},
}

const defaultNetwork = "builtin" // coophive

func init() {
	userDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	appConfig[enums.APP_DIR].defaultVal = path.Join(userDir, "coophive")
}