package main

import (
	"fmt"

	"github.com/kekaswork/grpc-auth/internal/config"
)

func main() {
	// TODO: init app configuration
	cfg := config.MustLoad()

	fmt.Println(cfg)
	// TODO: init logger
	// TODO: init application itself
	// TODO: start grpc server
}
