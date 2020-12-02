@echo off
del /Q assets
go-assets-builder.exe config -p assets -o assets/assets.go
@echo Package assets success
go build
@echo Build success