# go-scaffold

A simple tool to initialize a new project making use of existing templates, you can also add modules later and make use of code generation to generate parts of your project.

### Command Examples
```
# use specific repo
gaffold init project1 --repo rog-golang-buddies/golang-template-repository

# use standard language project
gaffold init project1 --lang go

# add barebone gin api server
gaffold add rog-golang-buddies/gaffold-go-gin

# add gin api server and generate functions using openapi spec
gaffold add rog-golang-buddies/gaffold-go-gin --spec openapi.yaml

# add barebone gRPC server
gaffold add rog-golang-buddies/gaffold-go-gRPC

# add gRPC server and generate proto bindings for go
gaffold add rog-golang-buddies/gaffold-go-gRPC --spec myservice.proto

# add gRPC server and generate proto bindings for go
gaffold add rog-golang-buddies/gaffold-go-cli --spec cli.spec
```