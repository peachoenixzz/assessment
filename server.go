package main

import (
	"fmt"
	"github.com/peachoenixz/assessment/internal/api"
	env "github.com/peachoenixz/assessment/pkg/environment"
	"github.com/peachoenixz/assessment/pkg/log"
)

func main() {
	err := env.InitEnv()
	if err != nil {
		log.ErrorLog(fmt.Sprintf("Error Read Environment : %v", err), "SERVICE ENV")
		return
	}
	api.EchoStart()
}
