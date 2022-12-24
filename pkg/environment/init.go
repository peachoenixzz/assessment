package environment

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/peachoenixz/assessment/pkg/log"
)

type CustomEnv struct {
	PORT     string
	DATABASE string
}

func (ce CustomEnv) checkCustomEnv() bool {
	return ce.PORT != "" && ce.DATABASE != ""
}

//func (ce CustomEnv) customVariableEnv() {
//	err := os.Setenv("PORT", ce.PORT)
//	if err != nil {
//		return
//	}
//	err = os.Setenv("DATABASE_URL", ce.DATABASE)
//	if err != nil {
//		return
//	}
//}

func InitEnv() {
	ce := CustomEnv{
		PORT:     os.Getenv("PORT"),
		DATABASE: os.Getenv("DATABASE_URL"),
	}
	if !ce.checkCustomEnv() {
		ReadEnv("environment")
	}
	//ReadEnv("environment")
}

func ReadEnv(filename string) string {
	if err := godotenv.Load(fmt.Sprintf("%v.env", filename)); err != nil {
		log.ErrorLog(err, "ENV")
		return "failed"
	}
	log.InfoLog("READ ENV FILE SUCCESS", "ENV")
	return "success"
}
