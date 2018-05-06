TAG ?= master
VCS_REF = $(shell git rev-parse --verify HEAD)
BUILD_DATE = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

build-develop:
	docker build \
		-t avegao/iot-fronius-push-service\:develop \
		--build-arg="VCS_REF=$(VCS_REF)" \
		--build-arg="BUILD_DATE=$(BUILD_DATE)" \
		.

push-develop:
	docker push avegao/iot-fronius-push-service\:develop

publish-develop: build-develop push-develop

tag-develop-version:
	docker tag avegao/iot-fronius-push-service\:develop avegao/iot-fronius-push-service\:$(VERSION)
	docker push avegao/iot-fronius-push-service\:$(VERSION)

promote-develop-to-master:
	docker pull avegao/iot-fronius-push-service\:develop
	docker tag avegao/iot-fronius-push-service\:develop avegao/iot-fronius-push-service\:master
	docker tag avegao/iot-fronius-push-service\:develop avegao/iot-fronius-push-service\:latest
	docker push avegao/iot-fronius-push-service\:master
	docker push avegao/iot-fronius-push-service\:latest

build-master:
	docker build \
		-t avegao/iot-fronius-push-service\:master \
		-t avegao/iot-fronius-push-service\:latest \
		--build-arg="VCS_REF=$(VCS_REF)" \
		--build-arg="BUILD_DATE=$(BUILD_DATE)" \
		.

push-master:
	docker push avegao/iot-fronius-push-service\:master
	docker push avegao/iot-fronius-push-service\:latest

publish-master: build-master push-master

tag-version:
	docker pull avegao/iot-fronius-push-service\:master
	docker tag avegao/iot-fronius-push-service\:master avegao/iot-fronius-push-service\:$(VERSION)
	docker push avegao/iot-fronius-push-service\:$(VERSION)
