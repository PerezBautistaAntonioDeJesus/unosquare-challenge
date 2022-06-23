start: check_swagger
	GO111MODULE=off swagger generate spec -o ./api/swagger/swagger.yaml --scan-models
check_swagger:
    which swagger || GO111MODULE=off go get -v github.com/go-swagger/go-swagger/cmd/swagger