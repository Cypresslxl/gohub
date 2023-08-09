package main

import (
	"flag"
	"fmt"
)

func main() {
	var env string

	// Define a string flag named "env"
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")

	// Parse command-line arguments
	flag.Parse()

	// Print the value of the "env" flag
	fmt.Println("Environment:", env)
}
