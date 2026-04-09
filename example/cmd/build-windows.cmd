@echo on

cd %~dp0

cd ../

SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=amd64
go build -ldflags="-s -w" -o ./data/windows-amd64/example.exe
