rm -rf assets/*
go-assets-builder resources -p assets -o assets/assets.go
echo Package assets success
mkdir -p build
go build -o build/
echo Build success