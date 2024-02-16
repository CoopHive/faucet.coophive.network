package config

import (
	"fmt"
	"os"
	"path"

	"github.com/CoopHive/hive/enums"
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
		"$HOME/coophive",
	},
	// derived
	// enums.APP_DATA_DIR: {
	// 	"App Data Location: typically used for storing github repos and results",
	// 	"$APP_DIR/data",
	// },
	//
	// enums.APP_PLUGIN_DIR: {
	// 	"Plugin Path: typically used for storing plugins",
	// 	"$APP_DIR/plugins",
	// },

	enums.HIVE_CONTROLLER: { // TODO: network dependent
		"Web3 Controller Address",
		"0xCCAaFD2AdD790788436f10e2C84585C46388b9aF",
	},
	enums.HIVE_SOLVER: {
		"Hive Solver  Address",
		"0xd4646ef9f7336b06841db3019b617ceadf435316",
	},
	enums.HIVE_MEDIATION: {
		"Hive Mediation Addresss : can be set of addresses separated by ','",
		"0x2d83ced7562e406151bd49c749654429907543b4",
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
