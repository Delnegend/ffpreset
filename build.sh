#/bin/bash
rm -rf ./dist
mkdir ./dist

echo "Building ffpreset-linux-amd64 ..."
GOOS=linux GOARCH=amd64 go build -o ./dist/ffpreset-linux-amd64

echo "Building ffpreset-linux-arm64 ..."
GOOS=linux GOARCH=arm64 go build -o ./dist/ffpreset-linux-arm64

echo "Building ffpreset-apple-intel ..."
GOOS=darwin GOARCH=amd64 go build -o ./dist/ffpreset-apple-intel

echo "Building ffpreset-apple-silicon ..."
GOOS=darwin GOARCH=arm64 go build -o ./dist/ffpreset-apple-silicon

echo "Building ffpreset-windows-amd64 ..."
GOOS=windows GOARCH=amd64 go build -o ./dist/ffpreset-windows-amd64.exe

echo "Building ffpreset-windows-arm64 ..."
GOOS=windows GOARCH=arm64 go build -o ./dist/ffpreset-windows-arm64.exe