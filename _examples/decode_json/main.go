package main

import (
	"fmt"

	"github.com/savaki/s3config"
)

type Config struct {
	Env string `json:"env"`
}

func main() {
	config := Config{}
	s3config.NewDecoder("bucket-name", "path-to-config").Decode(&config)
	fmt.Println("env is", config.Env)
}
