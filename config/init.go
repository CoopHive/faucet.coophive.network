package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"

	"github.com/CoopHive/faucet.coophive.network/enums"
)

func init() {

	configFile := os.Getenv("CONFIG_FILE")

	if configFile == "" {
		configFile = ".env"
	}

	logrus.Infof("Loading config from %s", configFile)

	if err := godotenv.Load(configFile); err != nil {
		logrus.Errorf(".env loading error %v", err)
	}

	// fmt.Printf("CoopHive: %s\n", hive.VERSION)
	pf := pflag.NewFlagSet("conf", pflag.ContinueOnError)

	checkDup := func(key string, block string) {
		if Conf.IsSet(key) {
			log.Fatalf("duplicate key found in Conf[%s]: %s", block, key)
		}
	}

	// formatEnvVar := func(key string) string {
	// 	k := strings.Replace("-", "_", key, -1)
	// 	k = strings.ToLower(k)
	// 	return k
	// }

	cmdFlags := map[string]bool{
		enums.APP_DIR:            false,
		enums.FAUCET_PRIVATE_KEY: false,
		enums.FAUCET_PORT:        false,
	}

	for key, meta := range buildConfig {
		checkDup(key, "build")
		Conf.Set(key, meta.defaultVal)
	}

	for key, meta := range appConfig {
		checkDup(key, "app")

		Conf.SetDefault(key, meta.defaultVal)

		// automatic conversion of environment var key to `UPPER_CASE` will happen.
		Conf.BindEnv(key)

		if !cmdFlags[key] {
			// key := strings.Replace("_", "-", key, -1)
			// read command-line arguments
			pf.String(key, meta.defaultVal, meta.desc)
			pflag.String(key, meta.defaultVal, meta.desc) // to show in usage
		}
	}

	if err := pf.Parse(os.Args[1:]); err != nil {
		logrus.Debugf("failed to parse args due to %v", err)
	}

	if err := Conf.BindPFlags(pf); err != nil {
		logrus.Debugf("failed to load flags:%v", err)
	}

	if Conf.GetBool(enums.DEBUG) {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetReportCaller(true)
	}

	appDir := Conf.GetString(enums.APP_DIR)

	logrus.Debugln("appDir: ", appDir)

	/*Network related Conf*/

	network := Conf.GetString(enums.NETWORK)

	logrus.Debugln("network: ", network)

	pKey := Conf.GetString(enums.FAUCET_PRIVATE_KEY)
	if pKey == "" {
		logrus.Info("setting web3 private key with faucet priv key")
		Conf.Set(enums.FAUCET_PRIVATE_KEY, Conf.Get(enums.WEB3_PRIVATE_KEY))
	}

	port := Conf.GetString(enums.FAUCET_PORT)

	if port != "" {
		Conf.Set(enums.PORT, port)
		logrus.Info("setting faucet port with PORT", port)

	}

}
