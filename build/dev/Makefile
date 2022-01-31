APP_NAME=capstone
BRANCH=dev
VERSION=v1.0.0
DEVICE=mac
IMAGE=$(DEVICE)/$(APP_NAME):$(VERSION)-$(BRANCH)
CONTAINER=$(APP_NAME)_$(BRANCH)
DB_NET=pg_dev
                                                              
build: src/.ssh .env
	docker build -t $(IMAGE) .
# If the nework create fails, does not stop the commands sequence
	docker network create --driver bridge $(CONTAINER) || true
	docker volume create $(CONTAINER)
	docker run \
		--rm \
		--name $(CONTAINER)-build \
		--mount type=volume,source=$(CONTAINER),destination=/go \
		--mount type=bind,source=$(shell pwd)/src/.ssh,destination=/home/$(APP_NAME)/.ssh \
		--mount type=bind,source=$(shell pwd)/build.sh,destination=/build.sh,readonly \
		--hostname api-$(BRANCH).$(APP_NAME).$(DEVICE) \
		$(IMAGE) /build.sh
# --interactive --tty \
# $(IMAGE) bash

clean:
	docker stop $(CONTAINER) || true
	docker rm $(CONTAINER) || true
	docker network rm $(CONTAINER)
	docker volume rm $(CONTAINER)

run: src/.ssh .env
	docker run \
		--detach \
		--name $(CONTAINER) \
		--hostname api-$(BRANCH).$(APP_NAME).$(DEVICE) \
		--env-file .env \
		--network $(DB_NET) \
		--mount type=volume,source=$(CONTAINER),destination=/go \
		--mount type=bind,source=$(shell pwd)/src/.ssh,destination=/home/$(APP_NAME)/.ssh \
		--restart unless-stopped \
		--interactive --tty \
		$(IMAGE) bash

stop:
	docker stop $(CONTAINER)
	docker rm $(CONTAINER)

restart:
	make stop
	make run

exec:
	docker exec \
	--workdir /go/src/$(APP_NAME) \
	--interactive --tty \
	$(CONTAINER) bash

.PHONY: build clean run stop restart exec