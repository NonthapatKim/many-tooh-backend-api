package main

import (
	"fmt"
	"os"

	"github.com/chanitt/go-hexagonal-template/internal/server"
)

func main() {

	if err := server.RunServer(); err != nil {
		fmt.Println("Run server failed!", err)
		os.Exit(1)
	}

}
