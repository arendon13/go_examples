build.arm:
	GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap main.go

build.amd:
	GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap main.go

binary.zip:
	zip myGoFunction.zip bootstrap

build.zip.arm: build.arm binary.zip

build.zip.amd: build.amd binary.zip

s3.upload:
	aws s3 cp myGoFunction.zip s3://recipes-api-pplns-arti/myGoFunction