package main

import (
	"fmt"
	"os"

	env "github.com/peachoenixz/assessment/pkg/environment"
)

func main() {
	env.InitEnv()
	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", os.Getenv("PORT"))
}
