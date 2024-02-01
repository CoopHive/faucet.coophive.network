package main

import (
	"github.com/joho/godotenv"

	"github.com/coophive/faucet.coophive.network/cmd"
)

//go:generate pnpm install
func main() {
	cmd.Execute()
}

func init() {
	_ = godotenv.Load(cmd.GetFromEnv("DOTENV_FILE", ".env"))

}
