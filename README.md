# go-gin-crud-gorm
## This project is a sample project, that illustrate how we can use GIN/GORM together with postgresql database to do some crud operations. 
## Install AIR : go install github.com/cosmtrek/air@latest
### $(go env GOPATH)/bin/air
### OR alias air='$(go env GOPATH)/bin/air'


## LINK : https://codevoweb.com/setup-golang-gorm-restful-api-project-with-postgres/ 

## for Migration run -> go run migrate/migrate.go 

# PROTO Related thing: [ https://codevoweb.com/golang-grpc-server-and-client-signup-user-verify-email ]
## Install PROTO COMpiler : 
### apt install -y protobuf-compiler & sudo apt install golang-goprotobuf-dev [ https://grpc.io/docs/protoc-installation/]
``````
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get github.com/golang/protobuf
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

export PATH="$PATH:$(go env GOPATH)/bin"

   sudo snap install protobuf or sudo apt  install protobuf-compiler

``````
