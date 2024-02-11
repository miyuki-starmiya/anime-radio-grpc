# anime-radio-grpc
Send latest anime radio infomation to Slack via gRPC

## Directory Structure
```sh
my-grpc-project/
├── api/
│   └── proto/
│       └── service.proto      # Protobuf definitions
├── gen/
│   ├── go/                    # Go generated files
│   │   └── service.pb.go      # Generated Go gRPC client code (example file name)
│   └── python/                # Python generated files
│       ├── service_pb2.py     # Generated Python gRPC code
│       └── service_pb2_grpc.py # Generated Python gRPC server code
├── go-client/
│   ├── go.mod                 # Go module file
│   ├── go.sum                 # Go module checksums
│   └── main.go                # Go client implementation
└── python-server/
    ├── requirements.txt       # Python dependencies
    └── server.py              # Python server implementation
```
