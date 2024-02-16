package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/CoopHive/faucet.coophive.network/cmd"
)

func main() {
	cmd.Execute()
}

func init() {
	configFile := os.Getenv("CONFIG_FILE")

	if configFile == "" {
		configFile = ".env"
	}

	if err := godotenv.Load(configFile); err != nil {
		logrus.Errorf(".env loading error %v", err)
	}

}
