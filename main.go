package main

import (
	"fmt"
	"os"

	"github.com/MendelGusmao/thomson-cli/modules"
	_ "github.com/MendelGusmao/thomson-cli/modules/gateway"
)

func main() {
	module, err := modules.Root.Find(os.Args[1:])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(module.URI())
}
