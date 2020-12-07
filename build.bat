@echo off
del /Q assets

@REM @REM Don't pack resources
@REM @echo package assets > assets/assets.go
@REM @echo.>> assets/assets.go
@REM @echo import "github.com/jessevdk/go-assets" >> assets/assets.go
@REM @echo.>> assets/assets.go
@REM @echo var Assets = assets.NewFileSystem(map[string][]string{}, map[string]*assets.File{}, "") >> assets/assets.go

@REM Pack resources
go-assets-builder.exe -p assets -o assets/assets.go resources
@echo Package assets success

if not exist build (
    mkdir build
)
go build -ldflags="-w -s" -o build/
@echo Build success