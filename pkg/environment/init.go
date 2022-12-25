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

func InitEnv() error {
	ce := CustomEnv{
		PORT:     os.Getenv("PORT"),
		DATABASE: os.Getenv("DATABASE_URL"),
	}
	if !ce.checkCustomEnv() {
		err := ReadEnv("environment")
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadEnv(filename string) error {
	if err := godotenv.Load(fmt.Sprintf("%v.env", filename)); err != nil {
		log.ErrorLog(err, "ENV")
		return err
	}
	log.InfoLog("READ ENV FILE SUCCESS", "ENV")
	return nil
}
