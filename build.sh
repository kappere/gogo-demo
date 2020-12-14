mkdir -p assets
rm -rf assets/*

# # Don't pack resources
# echo 'package assets' > assets/assets.go
# echo '' >> assets/assets.go
# echo 'import "github.com/jessevdk/go-assets"' >> assets/assets.go
# echo '' >> assets/assets.go
# echo 'var Assets = assets.NewFileSystem(map[string][]string{}, map[string]*assets.File{}, "")' >> assets/assets.go

# Pack resources
go-assets-builder -p assets -o assets/assets.go resources
echo Package assets success

mkdir -p build
go build -ldflags="-w -s" -o build/
echo Build success