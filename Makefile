generate:
	cd server && poetry run python -m grpc_tools.protoc -I../api/proto --python_out=../gen/python --grpc_python_out=../gen/python ../api/proto/anime_radio.proto
	protoc ./api/proto/anime_radio.proto --go_out=./gen/go --go-grpc_out=./gen/go
