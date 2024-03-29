.PHONY: setup
setup:
	bash ./hack/start-docker.sh

.PHONY: install-bun
install-bun: ## パッケージのインストール
	bun i

.PHONY: adr
adr: install-bun ## コードジェネレータを呼び出してADRを生成
	bun run plop adr

.PHONY: migrate
migrate: setup
	docker-compose run --rm migration

test: migrate
	cd ./packages/orion && \
		bash ./hack/run-local.sh go test -v ./... && \
		cd -
