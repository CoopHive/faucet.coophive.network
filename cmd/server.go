package cmd

import (
	"crypto/ecdsa"
	"errors"
	"flag"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"

	"github.com/CoopHive/faucet.coophive.network/config"
	"github.com/CoopHive/faucet.coophive.network/enums"
	"github.com/CoopHive/faucet.coophive.network/internal/chain"
	"github.com/CoopHive/faucet.coophive.network/internal/server"
)

var versionFlag bool

func init() {
	flag.BoolVar(&versionFlag, "version", false, "print version")

	flag.Parse()

	if versionFlag {
		fmt.Println(config.Conf.GetString(enums.VERSION))
		os.Exit(0)
	}

	// if err := godotenv.Load(configFile); err != nil {
	// 	log.Errorf("failed to load configfile-%s %v", configFile, err)
	// }
}

func Execute() {
	conf := config.Conf
	serverConfig := &server.Config{
		conf.GetString(enums.NETWORK),
		conf.GetString(enums.FAUCET_SYMBOL),
		conf.GetInt(enums.PORT),
		conf.GetInt(enums.FAUCET_MINUTES),
		conf.GetInt(enums.FAUCET_AMOUNT),
		conf.GetInt(enums.FAUCET_TOKENAMOUNT),
		conf.GetInt(enums.PROXY_COUNT),
		conf.GetString(enums.HCAPTCHA_SITEKEY),
		conf.GetString(enums.HCAPTCHA_SECRET),
	}

	privateKey, err := createPrivateKey(conf.GetString(enums.WEB3_PRIVATE_KEY))
	if err != nil {
		panic(fmt.Errorf("failed to read private key: %w", err))
	}

	chainID := conf.GetInt64(enums.WEB3_CHAIN_ID)

	provider := conf.GetString(enums.WEB3_RPC_URL)

	txBuilder, err := chain.NewTxBuilder(provider, privateKey, big.NewInt(chainID), common.HexToAddress(conf.GetString(enums.WALLET_TOKENADDRESS)))
	if err != nil {
		logrus.Info("provider:", provider)
		panic(fmt.Errorf("cannot connect to web3 provider: %w", err))
	}

	go server.NewServer(txBuilder, serverConfig).Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func createPrivateKey(hexkey string) (*ecdsa.PrivateKey, error) {

	if strings.Trim(hexkey, " ") == "" {
		return nil, errors.New("missing private key or keystore")
	}
	if chain.Has0xPrefix(hexkey) {
		hexkey = hexkey[2:]
	}
	return crypto.HexToECDSA(hexkey)
}

/*func getPrivateKeyFromFlags() (*ecdsa.PrivateKey, error) {
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
*/
