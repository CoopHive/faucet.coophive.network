# CoopHive Faucet

[![Build](https://img.shields.io/github/actions/workflow/status/CoopHive/faucet.coophive.network/build.yml?branch=main)](https://github.com/CoopHive/faucet.coophive.network/actions/workflows/build.yml)
[![Release](https://img.shields.io/github/v/release/CoopHive/faucet.coophive.network)](https://github.com/CoopHive/faucet.coophive.network/releases)
[![Report](https://goreportcard.com/badge/github.com/CoopHive/faucet.coophive.network)](https://goreportcard.com/report/github.com/CoopHive/faucet.coophive.network)
[![Go](https://img.shields.io/github/go-mod/go-version/CoopHive/faucet.coophive.network)](https://go.dev/)
[![License](https://img.shields.io/github/license/CoopHive/faucet.coophive.network)](./LICENSE)

The faucet is a web application with the goal of distributing small amounts of Ether in private and test networks.

## Features

* Allow to configure the funding account via private key or keystore
* Asynchronous processing Txs to achieve parallel execution of user requests
* Rate limiting by ETH address and IP address as a precaution against spam
* Prevent X-Forwarded-For spoofing by specifying the count of reverse proxies

## Get started

### Installation

#### For Linux, Unix, MacOS


```shell
OSARCH=$(uname -m | awk '{if ($0 ~ /arm64|aarch64/) print "arm64"; else if ($0 ~ /x86_64|amd64/) print "amd64"; else print "unsupported_arch"}') && export OSARCH
echo $OSARCH
OSNAME=$(uname -s | awk '{if ($1 == "Darwin") print "darwin"; else if ($1 == "Linux") print "linux"; else print "unsupported_os"}') && export OSNAME;
echo $OSNAME


tag="v0.5.4"
url="https://github.com/CoopHive/faucet.coophive.network/releases/download/$tag/faucet-$OSNAME-$OSARCH"

curl -sSL -o faucet $url
chmod +x faucet
./faucet version

```

#### Via GUI

1. Go to https://github.com/CoopHive/faucet.coophive.network/releases/
2. Navigate to latest stable semver release i.e release of format vX.Y.Z

#### Via Go 1.21+

`go install github.com/CoopHive/faucet.coophive.networke@latest`


## Usage

**Use private key to fund users**

```bash
./faucet -httpport 8080 -wallet.provider http://localhost:8545 -wallet.privkey privkey
```

**Use keystore to fund users**

```bash
./faucet -httpport 8080 -wallet.provider http://localhost:8545 -wallet.keyjson keystore -wallet.keypass password.txt
```

### Configuration

You can configure the funder by using environment variables instead of command-line flags as follows:
```bash
export WEB3_PROVIDER=rpc endpoint
export PRIVATE_KEY=hex private key
```

or

```bash
export WEB3_PROVIDER=rpc endpoint
export KEYSTORE=keystore path
echo "your keystore password" > `pwd`/password.txt
```

Then run the faucet application without the wallet command-line flags:
```bash
./faucet -httpport 8080
```

**Optional Flags**

The following are the available command-line flags(excluding above wallet flags):

| Flag               | Description                                      | Default Value |
|--------------------|--------------------------------------------------|---------------|
| --port             | Listener port to serve HTTP connection           | 8080          |
| --proxy_count      | Count of reverse proxies in front of the server  | 0             |
| --eth_drip         | Number of Ethers to transfer per user request    | 0             |
| --faucet_minutes   | Number of minutes to wait between funding rounds | 60            |
| --faucet_name      | Network name to display on the frontend          | testnet       |
| --faucet_symbol    | Token symbol to display on the frontend          | ETH           |
| --hcaptcha_sitekey | hCaptcha sitekey                                 |               |
| --hcaptcha_secret  | hCaptcha secret                                  |               |

### Docker 

```bash
docker run -p 8080:8080 -e WEB3_PROVIDER="" -e PRIVATE_KEY="$PRIVATE_KEY" ghcr.io/coophive/faucet:latest
```

<!--
or

```bash
docker run -d -p 8080:8080 -e WEB3_PROVIDER=rpc endpoint -e KEYSTORE=keystore path -v `pwd`/keystore:/app/keystore -v `pwd`/password.txt:/app/password.txt CoopHive/faucet.coophive.network:1.1.0
```

-->

## License

Distributed under the MIT License. See LICENSE for more information.
