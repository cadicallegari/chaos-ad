version ?= latest
APP_NAME = mypkg
IMAGE = $(APP_NAME):$(version)

release: publish
	git tag -a $(version) -m "Generated release "$(version)
	git push origin $(version)

publish: image
	docker push $(IMAGE)

image:
	docker build -t $(IMAGE) .

shell: image
	docker-compose run --rm $(APP_NAME) sh

logs-%:
	docker-compose logs $*

stop:
	docker-compose stop

cleanup: stop
	docker-compose rm -f -v

check: image
	docker-compose run --rm $(APP_NAME) ./hack/check.sh $(parameters)

check-integration: image
	docker-compose run --rm $(APP_NAME) ./hack/check-integration.sh $(parameters)

coverage: image
	docker-compose run --rm --entrypoint sh $(APP_NAME) ./hack/check.sh --coverage

coverage-show: coverage
	xdg-open ./tests/coverage/index.html

init:
	pip install -r requirements.txt

# test:
# 	@python tests/*/*_test.py

install: init
	pip install -e .

# test_build_server:
# 	cd hack && docker build -t mock_server .

# test_run_server: test_build_server
# 	docker run --rm -it -p 4567:4567 mock_server
