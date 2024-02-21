package main

import (
	"fmt"

	"github.com/luqmanMohammed/golang-generate-on-release/version"
)

func main() {
	fmt.Println(version.WrapperVersion())
}
