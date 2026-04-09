@echo on

cd %~dp0

cd ../

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags="-s -w" -o ./data/linux-amd64/example

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm64
go build -ldflags="-s -w" -o ./data/linux-arm64/example

SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -ldflags="-s -w" -o ./data/darwin-amd64/example

SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=arm64
go build -ldflags="-s -w" -o ./data/darwin-arm64/example

SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=amd64
go build -ldflags="-s -w" -o ./data/windows-amd64/example.exe

SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=amd64
