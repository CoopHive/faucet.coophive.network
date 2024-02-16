package config

import (
	"fmt"
	"os"
	"path"
	"strconv"

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

	enums.FAUCET_AMOUNT: {
		desc:       "Number of Ethers to transfer per user request",
		defaultVal: FAUCET_ETHER_AMOUNT,
	},
	enums.FAUCET_TOKENAMOUNT: {
		desc:       "Number of Tokens to transfer per user request",
		defaultVal: FAUCET_TOKEN_AMOUNT,
	},
	enums.FAUCET_MINUTES: {
		desc:       "Number of minutes to wait between funding rounds",
		defaultVal: strconv.Itoa(FaucetInterval()),
	},
	enums.FAUCET_NAME: {
		desc:       "Network name to display on the frontend",
		defaultVal: GetFromEnv("FAUCET_NAME", "CALIBRATION"),
	},
	enums.FAUCET_SYMBOL: {
		desc:       "Token symbol to display on the frontend",
		defaultVal: GetFromEnv("FAUCET_SYMBOL", "HIVE"),
	},
	enums.WALLET_KEYJSON: {
		desc:       "Keystore file to fund user requests with",
		defaultVal: GetFromEnv("KEYSTORE", ""),
	},
	enums.WALLET_KEYPASS: {
		desc:       "Passphrase text file to decrypt keystore",
		defaultVal: GetFromEnv("KEYSTORE_PASS", "password.txt"),
	},
	enums.WALLET_PRIVKEY: {
		desc:       "Private key hex to fund user requests with",
		defaultVal: os.Getenv("PRIVATE_KEY"),
	},
	enums.WALLET_PROVIDER: {
		desc:       "Endpoint for Ethereum JSON-RPC connection",
		defaultVal: os.Getenv("WEB3_PROVIDER"),
	},
	enums.WALLET_TOKENADDRESS: {
		desc:       "Address of ERC-20 token contract",
		defaultVal: os.Getenv("TOKEN_ADDRESS"),
	},
	enums.HCAPTCHA_SITEKEY: {
		desc:       "hCaptcha sitekey",
		defaultVal: os.Getenv("HCAPTCHA_SITEKEY"),
	},
	enums.HCAPTCHA_SECRET: {
		desc:       "hCaptcha secret",
		defaultVal: os.Getenv("HCAPTCHA_SECRET"),
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
