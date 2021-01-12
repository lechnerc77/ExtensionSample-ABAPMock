RELEASE=0.0.2
APP=abap-mock-in-go
DOCKER_ACCOUNT=<Your Dokcer ID>
CONTAINER_IMAGE=${DOCKER_ACCOUNT}/${APP}:${RELEASE}

.PHONY: build-image push-image

build-image:
	docker build -t $(CONTAINER_IMAGE) --no-cache --rm .

push-image: build-image
	docker push $(CONTAINER_IMAGE)

docker-run:
	docker run -it --rm -p 8000:8000 $(CONTAINER_IMAGE)