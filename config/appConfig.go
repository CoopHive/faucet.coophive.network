package config

import (
	"fmt"
	"os"
	"path"

	"github.com/CoopHive/faucet.coophive.network/enums"
)

var NETWORKS = []string{"coophive", "calibration"}

var appConfig = configMap[string]{
	enums.DEBUG: {
		desc:       "debug mode",
		defaultVal: "false",
	},
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

	enums.PORT: {
		"port to run",
		"8080",
	},
	enums.FAUCET_AMOUNT: {
		desc:       "Number of Ethers to transfer per user request",
		defaultVal: "1",
	},
	enums.FAUCET_TOKENAMOUNT: {
		desc:       "Number of Tokens to transfer per user request",
		defaultVal: "1",
	},
	enums.FAUCET_MINUTES: {
		desc:       "Number of minutes to wait between funding rounds",
		defaultVal: "1440",
	},
	enums.FAUCET_NAME: {
		desc:       "Network name to display on the frontend",
		defaultVal: NETWORKS[1],
	},
	enums.FAUCET_SYMBOL: {
		desc:       "Token symbol to display on the frontend",
		defaultVal: "HIVE",
	},
	enums.WALLET_KEYJSON: {
		desc:       "Keystore file to fund user requests with",
		defaultVal: "",
	},
	enums.WALLET_KEYPASS: {
		desc:       "Passphrase text file to decrypt keystore",
		defaultVal: "password.txt",
	},
	enums.WALLET_PRIVKEY: {
		desc:       "Private key hex to fund user requests with",
		defaultVal: "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", // FIXME: use WEB3_PRIVATE_KEY
	},
	enums.WALLET_PROVIDER: {
		desc:       "Endpoint for Ethereum JSON-RPC connection",
		defaultVal: "",
	},
	enums.WALLET_TOKENADDRESS: {
		desc:       "Address of ERC-20 token contract",
		defaultVal: "",
	},
	enums.HCAPTCHA_SITEKEY: {
		desc:       "hCaptcha sitekey",
		defaultVal: "",
	},
	enums.HCAPTCHA_SECRET: {
		desc:       "hCaptcha secret",
		defaultVal: "",
	},
	enums.PROXY_COUNT: {
		"proxy count",
		"0",
	},
}

const defaultNetwork = "coophive"

func init() {
	userDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	appConfig[enums.APP_DIR].defaultVal = path.Join(userDir, "coophive-faucet")
}
