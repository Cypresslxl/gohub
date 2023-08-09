package main

import (
	"fmt"
	"strings"
)

func main() {
	// var env string

	// Define a string flag named "env"
	// flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")

	// Parse command-line arguments
	// flag.Parse()

	// Print the value of the "env" flag
	// fmt.Println("Environment:", env)
	str := "broken pipe"
	if ok := strings.Contains(str, "ken"); !ok {
		fmt.Println("not ok")
	} else {
		fmt.Println("ok")
	}
}
