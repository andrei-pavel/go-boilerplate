package main

import (
	"fmt"

	"github.com/andrei-pavel/go-boilerplate/config"
)

func main() {
	for key, value := range config.Configuration.Data {
		fmt.Println(key, value)
	}
}
