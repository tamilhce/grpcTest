Simple GrpcMicroservice & Client
===============================

Steps To deploy the server 
---------------------------

Steps To run the CLI Client
----------------------------
    cd github.com/tamilhce/grpcTest/cliClient/pkg
    ./cliClient --address <serviceIp> --port <servicePort> --ops sum 10 20

Steps To Build New Image (optionally)
========================================
**proto**
    cd github.com/tamilhce/
    protoc --proto_path=proto  --go_out=plugins=grpc:proto service.proto
**server docker image**
    cd github.com/tamilhce/grpcTest/server
    docker build -t <docker-id>/grpctest 
    docker docker push <docker-id>/grpctest

**client binary**
    cd github.com/tamilhce/grpcTest/cliClient
    go build -o pkg/cliClient
