package config

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"

	"github.com/CoopHive/faucet.coophive.network/enums"
)

func init() {
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
		enums.APP_DIR: true,
		enums.NETWORK: true,
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

		if cmdFlags[key] {
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

}
