package main

import (
	"fmt"

	"github.com/savaki/s3config"
)

type Config struct {
	Env string `hcl:"env"`
}

func main() {
	config := Config{}
	s3config.NewDecoder("bucket-name", "path-to-config", s3config.HCL).Decode(&config)
	fmt.Println("env is", config.Env)
}
