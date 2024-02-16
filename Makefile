include .env
export

binName = faucet-$(shell uname -s)-$(shell uname -m)

build-ci:
	go build -v -ldflags="\
		-X 'github.com/CoopHive/faucet.coophive.network/config.version=$$(git describe --tags --abbrev=0)' \
		-X 'github.com/CoopHive/faucet.coophive.network/config.commitSha=$$(git rev-parse HEAD)' \
	" -o ./bin/faucet-ci .
	./bin/faucet-ci --version

export VERSION=$(git describe --tags --abbrev=0)
export COMMIT_SHA=$(git rev-parse HEAD)

build:
	goreleaser build --single-target --clean -o bin/faucet --snapshot

prerelease:
	echo "Version is $(VERSION)"
	goreleaser check
	goreleaser build --single-target --clean

release:
	goreleaser release --clean

install-faucet:
	goreleaser build --single-target --clean -o ./bin/${binName} --snapshot

.PHONY: release install-unix install-win build release release-linux make-bin

