package main

import (
	"fmt"

	"config"
)

func main() {
	for key, value := range config.Configuration.Data {
		fmt.Println(key, value)
	}
}
