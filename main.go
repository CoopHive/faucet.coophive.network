package main

import (
	"github.com/coophive/faucet.coophive.network/cmd"
)

//go:generate pnpm install
func main() {
	cmd.Execute()
}
