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

Steps To Build New Image (optional)
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

Pre-requisites
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

Future work/ Enhancements
=========================
    i Nginx Ingress with TLS terimination is avail in /k8s/yet_to_test folder - you can give a try & update the comments/rise issue 
    Helm package for the K8s deployment
    ii REST API based client implementation approach

FAQ
===
i) How to scale the application ? 
    By increasing the replication count in the deployment, the application can be easily scaled 
ii) How to upgrade/rollback  the app ? 
    By changing the image/config in the deployment (default strategy: Rolling update) we can upgrade or rollback to the previous version 
iii) Is the application is stateful/stateless 
     The applciation is stateless & so it is easy to scale as a microservice 
iv)is Application is TDD aware? 
    yes, The application pacakge is bundled with go unit test for each functions , The test cases can be run autamatically for new changes, we can extend additional end to end functional cases as part of CI/CD ? 
v) Is it CI/CD aware ?
    yes, The build procedure mentioned above can be added in the jenkins pipeline to create new image to be pushed into the repo,when ever a set of commit happens 
vii) How to extend the eventstore to this application ? 
    yes, there multiple design patterns available, one common approach is adding side-car to stream a logs(std out/erro) logs to external logging system ( e.g elasticsearch) 
viii) How the dependencies are managed? 
    By packing the applcation as container solves the dependencies & environment issues(cross-platform) 
Having said this, The application is cloud-native, scalable microservice application :)

    
