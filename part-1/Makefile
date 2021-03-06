version     ?= latest
chaos-adimg  = cadicallegari/chaos-ad:$(version)
devimg       = chaos-addev
packagename  = cadicallegari/chaos-ad
workdir      = /go/src/$(packagename)
runargs      = --rm -v `pwd`:$(workdir) --workdir $(workdir) $(devimg)
runcmd       = docker run -ti $(runargs)
runcompose   = docker-compose run $(runargs)
gitversion   = $(version)

guard-%:
	@ if [ "${${*}}" = "" ]; then \
		echo "Variable '$*' not set"; \
		exit 1; \
	fi

release: guard-version publish
	git tag -a $(version) -m "Generated release "$(version)
	git push origin $(version)

publish: image
	docker push $(chaos-adimg)

image: build
	docker build -t $(chaos-adimg) .

imagedev:
	docker build -t $(devimg) -f ./hack/Dockerfile.dev .

vendor: imagedev
	$(runcmd) ./hack/vendor.sh
	sudo chown -R $(USER):$(id -g -n) ./vendor
	sudo chown -R $(USER):$(id -g -n) ./Godeps

build: imagedev
	$(runcmd) go build -v -ldflags "-X main.Version=$(gitversion)" -o ./cmd/chaos-ad/chaos-ad ./cmd/chaos-ad/main.go

check: imagedev
	$(runcompose) ./hack/check.sh $(pkg) $(test) $(args)

check-integration: imagedev
	$(runcompose) ./hack/check-integration.sh $(pkg) $(test) $(args)

run: image
	docker-compose run --service-ports --entrypoint "/app/chaos-ad" --rm chaos-ad

shell: imagedev
	$(runcmd)

coverage: imagedev
	$(runcmd) ./hack/coverage.sh

coverage-show: coverage
	xdg-open coverage.html

cleanup:
	docker-compose down
