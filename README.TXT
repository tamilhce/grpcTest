Simple GrpcMicroservice & Client
===============================
This project contains end to end functionality to test simple Go GRpc based microservice depoloyed in K8s. 
It has 3 components 
    i - cli to access the compute service (add, mulitpy)
    ii - server docker Manifest
    iii - kubernetes deployement files 


Steps To deploy the server in the K8s cluster
----------------------------------------------
    ***refer to pre-requesties section before running the k8s deployement 
    - To deploy the server, service object 
    kubectl apply -f grpcTest/K8s/grpcServerDeploy.yaml
    - To deploy the ingress 
    kubectl apply -f grpcTest/K8s/voyager-ingress.yaml
    
Steps To run the CLI Client
----------------------------
    go get  github.com/tamilhce/grpcTest/cliClient
    cd <GOPATH>/src/github.com/tamilhce/grpcTest/cliClient
    go build -o <GOPATH>/bin/cliClient .
    <GOPATH>/bin/cliClient --address <serviceIp> --port <servicePort> --ops <operation> int1 int2
    i.e /home/aviuser/go/bin/cliClient --address 10.79.58.24 --port 80 --ops sum 10 20

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

To Run the Unit Test
====================
    cd github.com/tamilhce/grpcTest/server
    go test -v 

Pre-requesties
===============
i)  voyager  Ingress controller shoud be installed in prior 
    https://appscode.com/products/voyager/v12.0.0-rc.1/setup/install/
    Reason for choosing  voyager
    grpc support for Nginx Ingress is available in Nginx controller version >= 0.30
    In my current setup Nginx controller version is 0.25, facing some permission issue to upgrade to >=0.30
    voyager - HA proxy based, simple to use & stable 

Limitations
=============
    ingress grpc annotations are Ingress controller specific.
    In our case the Ingress controller & Service Type Load balancer should be made avilable in the cluster

Future work/ Enhacements
=========================
    i Nginx Ingress with TLS terimination is avail in /k8s/yet_to_test folder - you can give a try & update the comments/rise issue 
    Helm package for the K8s deployment
    ii REST API based client implementation approach

Best-Practices 
==============
    i) The application is loosely coupled & easly scalable by increasing the replica count 
    ii) easy to upgrade (default : rolling) thanks to K8s 
    iii) used to git for code tracking 
    iv) the application has unit tests to validate the sanity(TDD)
    v) can be integrated with the CI/CD readily to create new docker images  & you can upgrade the appication
       seemslessly without any down time.
    v) you can implement side-car for redirecting the std out logs to evenstore like elastic store /fluentd - Again thanks to K8s 
    vi)This application is stateless  & meets the cloud-native standards. 
