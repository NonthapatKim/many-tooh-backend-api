package main

import (
	"fmt"
	"os"

	"github.com/NonthapatKim/many_tooth_backend_api/internal/server"
)

func main() {

	if err := server.RunServer(); err != nil {
		fmt.Println("Run server failed!", err)
		os.Exit(1)
	}

}
