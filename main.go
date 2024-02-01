package main

import (
	"github.com/chainflag/eth-faucet/cmd"
)

//go:generate pnpm install
func main() {
	cmd.Execute()
}
