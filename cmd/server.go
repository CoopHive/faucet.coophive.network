package cmd

import (
	"crypto/ecdsa"
	"errors"
	"flag"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/coophive/faucet.coophive.network/internal/chain"
	"github.com/coophive/faucet.coophive.network/internal/server"
)

func GetFromEnv(key string, defaultVal string) (val string) {
	val = os.Getenv(key)

	if val == "" {
		val = defaultVal
	}
	return
}

var PORT = func() int {
	p, err := strconv.Atoi(GetFromEnv("PORT", "8080"))
	if err != nil {
		return 8080
	}
	return p
}()
var PROXY_COUNT = func() int {
	p, err := strconv.Atoi(GetFromEnv("PROXY_COUNT", "0"))
	if err != nil {
		return 0
	}
	return p
}()

var FAUCET_ETHER_AMOUNT = func() int {
	p, err := strconv.Atoi(GetFromEnv("PROXY_COUNT", "1"))
	if err != nil {
		return 0
	}
	return p
}()
var FAUCET_TOKEN_AMOUNT = func() int {
	p, err := strconv.Atoi(GetFromEnv("PROXY_COUNT", "1"))
	if err != nil {
		return 1
	}
	return p
}()

var FAUCET_INTERVAL = func() int {
	p, err := strconv.Atoi(GetFromEnv("FAUCET_INTERVAL", "1440"))
	if err != nil {
		return 1440
	}
	return p
}

var (
	appVersion = "v0.1.0"
	chainIDMap = map[string]int{"goerli": 5, "sepolia": 11155111, "calib": 314159, "fvm": 314}

	httpPortFlag = flag.Int("httpport", PORT, "Listener port to serve HTTP connection")
	proxyCntFlag = flag.Int("proxycount", PROXY_COUNT, "Count of reverse proxies in front of the server")
	versionFlag  = flag.Bool("version", true, "Print version number")

	payoutFlag       = flag.Int("faucet.amount", FAUCET_ETHER_AMOUNT, "Number of Ethers to transfer per user request")
	payoutTokensFlag = flag.Int("faucet.tokenamount", FAUCET_TOKEN_AMOUNT, "Number of Tokens to transfer per user request")
	intervalFlag     = flag.Int("faucet.minutes", FAUCET_INTERVAL(), "Number of minutes to wait between funding rounds")
	netnameFlag      = flag.String("faucet.name", GetFromEnv("FAUCET_NAME", "CoopHive v0 Testnet"), "Network name to display on the frontend")
	symbolFlag       = flag.String("faucet.symbol", GetFromEnv("FAUCET_SYMBOL", "HIVE"), "Token symbol to display on the frontend")

	keyJSONFlag  = flag.String("wallet.keyjson", GetFromEnv("KEYSTORE", ""), "Keystore file to fund user requests with")
	keyPassFlag  = flag.String("wallet.keypass", GetFromEnv("KEYSTORE_PASS", "password.txt"), "Passphrase text file to decrypt keystore")
	privKeyFlag  = flag.String("wallet.privkey", os.Getenv("PRIVATE_KEY"), "Private key hex to fund user requests with")
	providerFlag = flag.String("wallet.provider", os.Getenv("WEB3_PROVIDER"), "Endpoint for Ethereum JSON-RPC connection")
	tokenAddress = flag.String("wallet.tokenaddress", os.Getenv("TOKEN_ADDRESS"), "Address of ERC-20 token contract")

	hcaptchaSiteKeyFlag = flag.String("hcaptcha.sitekey", os.Getenv("HCAPTCHA_SITEKEY"), "hCaptcha sitekey")
	hcaptchaSecretFlag  = flag.String("hcaptcha.secret", os.Getenv("HCAPTCHA_SECRET"), "hCaptcha secret")
)

func init() {
	flag.Parse()
	if *versionFlag {
		fmt.Println(appVersion)
		os.Exit(0)
	}
}

func Execute() {
	privateKey, err := getPrivateKeyFromFlags()
	if err != nil {
		panic(fmt.Errorf("failed to read private key: %w", err))
	}
	var chainID *big.Int
	if value, ok := chainIDMap[strings.ToLower(*netnameFlag)]; ok {
		chainID = big.NewInt(int64(value))
	}

	txBuilder, err := chain.NewTxBuilder(*providerFlag, privateKey, chainID, common.HexToAddress(*tokenAddress))
	if err != nil {
		panic(fmt.Errorf("cannot connect to web3 provider: %w", err))
	}
	config := server.NewConfig(*netnameFlag, *symbolFlag, *httpPortFlag, *intervalFlag, *payoutFlag, *payoutTokensFlag, *proxyCntFlag, *hcaptchaSiteKeyFlag, *hcaptchaSecretFlag)
	go server.NewServer(txBuilder, config).Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func getPrivateKeyFromFlags() (*ecdsa.PrivateKey, error) {
	if *privKeyFlag != "" {
		hexkey := *privKeyFlag
		if chain.Has0xPrefix(hexkey) {
			hexkey = hexkey[2:]
		}
		return crypto.HexToECDSA(hexkey)
	} else if *keyJSONFlag == "" {
		return nil, errors.New("missing private key or keystore")
	}

	keyfile, err := chain.ResolveKeyfilePath(*keyJSONFlag)
	if err != nil {
		return nil, err
	}
	password, err := os.ReadFile(*keyPassFlag)
	if err != nil {
		return nil, err
	}

	return chain.DecryptKeyfile(keyfile, strings.TrimRight(string(password), "\r\n"))
}
