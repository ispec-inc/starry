.PHONY: setup
setup:
	bash ./hack/start-docker.sh

migrate: setup
	docker-compose run --rm migration

test: migrate
	cd ./packages/orion && \
		bash ./hack/run-local.sh go test -v ./... && \
		cd -
