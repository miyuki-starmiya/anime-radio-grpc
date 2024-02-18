DOCKER_REGISTRY=asia-northeast1-docker.pkg.dev
GCP_PROJECT_ID=grand-century-318202
REPO_NAME=anime-radio-grpc
FULL_IMAGE_PATH=$(DOCKER_REGISTRY)/$(GCP_PROJECT_ID)/$(REPO_NAME)

generate:
	cd server && poetry run python -m grpc_tools.protoc -I../api/proto --python_out=./gen --grpc_python_out=../gen ../api/proto/anime_radio.proto
	protoc ./api/proto/anime_radio.proto --go_out=./client/gen --go-grpc_out=./client/gen

build_and_push_server:
	cd server && docker build -t anime_radio_server .
	docker tag anime_radio_server:latest $(FULL_IMAGE_PATH)/python-server:latest
	docker push $(FULL_IMAGE_PATH)/python-server:latest

build_and_push_client:
	cd client && docker build -t anime_radio_client .
	docker tag anime_radio_client:latest $(FULL_IMAGE_PATH)/go-client:latest
	docker push $(FULL_IMAGE_PATH)/go-client:latest
