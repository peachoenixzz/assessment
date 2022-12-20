package environment

import (
	"fmt"
	"github.com/peachoenixz/assessment/pkg/log"

	"github.com/joho/godotenv"
)

func ReadEnv(filename string) string {
	if err := godotenv.Load(fmt.Sprintf("%v.env", filename)); err != nil {
		log.ErrorLog(err, "ENV")
		return "failed"
	}
	log.InfoLog("READ ENV FILE SUCCESS", "ENV")
	return "success"
}
