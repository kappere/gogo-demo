@echo off
del /Q assets
go-assets-builder.exe resources -p assets -o assets/assets.go
@echo Package assets success
if not exist build (
    mkdir build
)
go build -o build/
@echo Build success