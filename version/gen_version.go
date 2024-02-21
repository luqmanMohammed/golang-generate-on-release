//go:build ignore

package main

import (
	"fmt"
	"os"
)

const code string = `
// Auto Generated by gen_version.go. DO NOT EDIT.

package version

func WrapperVersion() string {
	return "%s"
}
`

func main() {
	file, err := os.Create("version.go")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, code, os.Getenv("VERSION"))
}
