package config

import (
	"log"
	"os"
	"path"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"github.com/CoopHive/hive/enums"
)

var Module = fx.Options(
	fx.Provide(
		newConfig,
	),
	fx.Invoke(tempInitForFx),
	fx.WithLogger(func(conf *viper.Viper) (l fxevent.Logger) {
		if conf.GetBool(enums.DEBUG) {
			return &fxevent.ConsoleLogger{W: os.Stderr}
		}
		return fxevent.NopLogger
	}))

type out struct {
	fx.Out

	Conf *viper.Viper
}

func newConfig() (o out) {
	pf := pflag.NewFlagSet("conf", pflag.ContinueOnError)

	config := viper.New()

	checkDup := func(key string, block string) {
		if config.IsSet(key) {
			log.Fatalf("duplicate key found in config[%s]: %s", block, key)
		}
	}

	// formatEnvVar := func(key string) string {
	// 	k := strings.Replace("-", "_", key, -1)
	// 	k = strings.ToLower(k)
	// 	return k
	// }

	cmdFlags := map[string]bool{
		enums.APP_DIR:         true,
		enums.NETWORK:         true,
		enums.HIVE_CONTROLLER: true,
		enums.DEALER:          true,
		// enums.APP_DATA_DIR:   true,
		// enums.APP_PLUGIN_DIR: true,
	}

	for key, meta := range buildConfig {
		checkDup(key, "build")
		config.Set(key, meta.defaultVal)
	}

	for key, meta := range appConfig {
		checkDup(key, "app")

		config.SetDefault(key, meta.defaultVal)

		// automatic conversion of environment var key to `UPPER_CASE` will happen.
		config.BindEnv(key)

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

	if err := config.BindPFlags(pf); err != nil {
		logrus.Debugf("failed to load flags:%v", err)
	}

	if config.GetBool(enums.DEBUG) {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetReportCaller(true)
	}

	appDir := config.GetString(enums.APP_DIR)

	logrus.Debugln("appDir: ", appDir)

	config.Set(enums.APP_PLUGIN_DIR, path.Join(appDir, "plugins"))
	config.Set(enums.APP_DATA_DIR, path.Join(appDir, "data"))

	/*Network related config*/

	network := config.GetString(enums.NETWORK)

	logrus.Debugln("network: ", network)

	if true {
		c, err := loadDApp(network)

		if err != nil {
			log.Fatal("failed to load the network related dApps")
		}

		for key, val := range c {
			key = strings.ToLower(key)
			logrus.Debugln("setting network config: ", key, val)
			config.Set(key, val)
		}
		controller := config.Get(enums.HIVE_CONTROLLER)
		logrus.Debugln("controller: ", controller)

	}

	o.Conf = config
	return
}
