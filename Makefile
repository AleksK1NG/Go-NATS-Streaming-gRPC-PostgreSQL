.PHONY:

# ==============================================================================
# Docker

develop:
	echo "Starting develop docker compose"
	docker-compose -f docker-compose.yaml up --build

upload:
	docker build -t alexanderbryksin/stan_microservice:latest -f ./Dockerfile .
	docker push alexanderbryksin/stan_microservice:latest
	#APP_VERSION=latest docker-compose up

pull:
	docker pull alexanderbryksin/stan_microservice:latest


# ==============================================================================
# Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -t -d -v ./...
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache

# ==============================================================================
# Linters

run-linter:
	echo "Starting linters"
	golangci-lint run ./...


# ==============================================================================
# Docker support

FILES := $(shell docker ps -aq)

down-local:
	docker stop $(FILES)
	docker rm $(FILES)

clean:
	docker system prune -f

logs-local:
	docker logs -f $(FILES)


# ==============================================================================
# Make local SSL Certificate

cert:
	echo "Generating SSL certificates"
	cd ./ssl && sh instructions.sh