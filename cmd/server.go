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

	"github.com/CoopHive/faucet.coophive.network/config"
	"github.com/CoopHive/faucet.coophive.network/enums"
	"github.com/CoopHive/faucet.coophive.network/internal/chain"
	"github.com/CoopHive/faucet.coophive.network/internal/server"
)

func init() {
	flag.Parse()
	if *versionFlag {
		fmt.Println(config.Conf.GetString(enums.VERSION))
		os.Exit(0)
	}

	// if err := godotenv.Load(configFile); err != nil {
	// 	log.Errorf("failed to load configfile-%s %v", configFile, err)
	// }
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

	conf := config.Conf
	config := &server.Config{
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
