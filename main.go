package main

import (
	"fmt"
	"github.com/kyu08/gcp-url-generator/config"
)

func main() {
	for _, bookmark := range config.Main() {
		fmt.Printf("%v", bookmark)
	}
}
