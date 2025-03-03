build.arm:
	GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o ./bin/bootstrap main.go

build.amd:
	GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o ./bin/bootstrap main.go

binary.zip:
	zip ./bin/myFunction.zip ./bin/bootstrap

build.zip.arm: build.arm binary.zip

build.zip.amd: build.amd binary.zip