package main

import (
	"os"

	"github.com/ali-furkan/wo/cmd/wo"
)

func main() {
	code := wo.Run()
	os.Exit(code)
}
