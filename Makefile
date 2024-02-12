generate:
	cd server && poetry run python -m grpc_tools.protoc -I../api/proto --python_out=./gen --grpc_python_out=../gen ../api/proto/anime_radio.proto
	protoc ./api/proto/anime_radio.proto --go_out=./client/gen --go-grpc_out=./client/gen
