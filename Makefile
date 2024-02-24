#include .env
#export

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

install:
	goreleaser build --single-target --clean -o bin/faucet --snapshot

prerelease:
	echo "Version is $(VERSION)"
	goreleaser check
	goreleaser build --single-target --clean


snapshot:
	goreleaser build --clean --snapshot

sync:
	make snapshot

	scp dist/cli_linux_amd64_v1/bin hive:/usr/local/bin/faucet
	scp dist/cli_linux_amd64_v1/bin hive1:/usr/local/bin/faucet

release-snapshot:
	goreleaser release --clean --snapshot

install-faucet:
	goreleaser build --single-target --clean -o ./bin/faucet --snapshot
	#go install
	cp bin/faucet ~/go/bin/faucet

.PHONY: release install-unix install-win build release release-linux make-bin install


build-frontend:
	cd web; npm run build;


dc:
	docker-compose pull && docker-compose up


sync-env:
	scp .env.* hive:./faucet/
