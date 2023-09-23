.PHONY: setup
setup:
	bash ./hack/start-docker.sh

migrate:
	docker-compose run --rm migration
