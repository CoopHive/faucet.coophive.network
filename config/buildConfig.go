package config

import (
	"github.com/spf13/viper"

	"github.com/coophive/faucet.coophive.network/enums"
)

var version string
var commitSha string

var buildConfig = configMap[string]{
	// app specific
	enums.APP_NAME: {
		"app name",
		"CoopHive Faucet",
	},
	enums.ENV: {
		"environment",
		enums.DEV,
	},
	enums.VERSION: {
		desc:       "version",
		defaultVal: version,
	},
	enums.COMMIT_SHA: {
		desc:       "commit sha",
		defaultVal: commitSha,
	},

	enums.RELEASE_URL: {
		desc:       "release url",
		defaultVal: "https://github.com/CoopHive/faucet.coophive.network/releases",
	},
}

var Conf = viper.New()
