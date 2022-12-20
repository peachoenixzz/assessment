package main

import (
	"github.com/peachoenixz/assessment/internal/api"
	env "github.com/peachoenixz/assessment/pkg/environment"
)

func main() {
	env.InitEnv()
	api.EchoStart()
}
