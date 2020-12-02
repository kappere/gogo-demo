package main

import (
	_ "gogo-demo/boot"

	gogo "wataru.com/gogo"
)

func main() {
	gogo.HttpServer().Run()
}
